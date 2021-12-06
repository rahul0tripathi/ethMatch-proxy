package routers

import (
	"github.com/ethMatch/proxy/api/controllers"
	"github.com/go-chi/chi/v5"
)

var (
	LobbyRouter = chi.NewRouter()
)

func init() {
	LobbyRouter.Route("/{"+controllers.PlayerAddr+"}", func(r chi.Router) {
		r.Get("/proposals", controllers.GetUserLobbyProposals)
		r.Post("/sign", controllers.SignProposal)
	})
}
