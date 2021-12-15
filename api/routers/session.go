package routers

import (
	"github.com/ethMatch/proxy/api/controllers"
	"github.com/go-chi/chi/v5"
)

var (
	SessionRouter = chi.NewRouter()
)

func init() {
	SessionRouter.Route("/", func(r chi.Router) {
		r.Post("/add_player", controllers.AddPlayerToSession)
		r.Post("/submit_result", controllers.SubmitSessionResult)
		r.Get("/{"+controllers.Session+"}", controllers.GetGameSession)
	})
}
