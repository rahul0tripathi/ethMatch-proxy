package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/orgs/ethMatch/p2pmatch/api/controllers"
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
