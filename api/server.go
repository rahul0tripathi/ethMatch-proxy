package api

import (
	"github.com/ethMatch/proxy/api/routers"
	"github.com/ethMatch/proxy/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"sync"
	"time"
)

var (
	Ws = WsHandler{
		players:       make(map[string]*websocket.Conn),
		io:            sync.Mutex{},
		flushInterval: time.Second * 2,
	}
)

func RunHTTPServer() {

	go Ws.FlushConnections()
	upgrader := websocket.Upgrader{
		HandshakeTimeout:  3000,
		EnableCompression: true,
	}
	baseRouter := chi.NewRouter()

	baseRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "ws://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))
	baseRouter.Mount("/ticket", routers.TicketRouter)
	baseRouter.Mount("/lobby", routers.LobbyRouter)
	baseRouter.Mount("/session", routers.SessionRouter)
	baseRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WsUpgradeHandler(upgrader, &Ws, w, r)
	})
	common.Logger.Info("starting up API server")
	err := http.ListenAndServe(":3004", baseRouter)
	if err != nil {
		common.Logger.Error("failed to startup API server", zap.Error(err))
	}
}
