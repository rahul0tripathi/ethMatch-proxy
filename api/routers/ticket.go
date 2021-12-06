package routers

import (
	"github.com/ethMatch/proxy/api/controllers"
	"github.com/go-chi/chi/v5"
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
