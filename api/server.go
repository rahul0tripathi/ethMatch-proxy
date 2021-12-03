package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/orgs/ethMatch/p2pmatch/api/routers"
	"github.com/orgs/ethMatch/p2pmatch/common"
	"go.uber.org/zap"
	"net/http"
)

func RunHTTPServer() {
	baseRouter := chi.NewRouter()
	baseRouter.Mount("/ticket", routers.TicketRouter)
	baseRouter.Mount("/lobby", routers.LobbyRouter)
	common.Logger.Info("starting up API server")
	err := http.ListenAndServe(":3333", baseRouter)
	if err != nil {
		common.Logger.Error("failed to startup API server", zap.Error(err))
	}
}
