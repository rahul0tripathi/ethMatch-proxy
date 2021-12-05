package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/orgs/ethMatch/p2pmatch/api/routers"
	"github.com/orgs/ethMatch/p2pmatch/common"
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
	baseRouter.Mount("/ticket", routers.TicketRouter)
	baseRouter.Mount("/lobby", routers.LobbyRouter)
	baseRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WsUpgradeHandler(upgrader, &Ws, w, r)
	})
	common.Logger.Info("starting up API server")
	err := http.ListenAndServe(":3333", baseRouter)
	if err != nil {
		common.Logger.Error("failed to startup API server", zap.Error(err))
	}
}
