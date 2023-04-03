package route

import (
	"measure/config"
	v1 "measure/webserver/handler/v1"

	"github.com/go-chi/chi/v5"
)

func SetupRoute(app *config.App) {
	router := app.Router

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/user", v1.GetListUser(app))
		// r.Get("/tenant", v1.GetListTenant(app))
		r.Get("/tenant", v1.PostCreateTenant(app))
	})
}
