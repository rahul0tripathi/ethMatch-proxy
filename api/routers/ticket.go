package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/orgs/ethMatch/p2pmatch/api/controllers"
)

var (
	TicketRouter = chi.NewRouter()
)

func init() {
	TicketRouter.Route("/{"+controllers.PlayerAddr+"}", func(r chi.Router) {
		r.Get("/", controllers.GetUserTickets)
		r.Post("/", controllers.AddNewTicket)
	})
}
