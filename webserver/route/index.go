package route

import (
	"measure/config"
	v1 "measure/webserver/handler/v1"

	"github.com/go-chi/chi/v5"
)

func SetupRoute(app *config.App) {
	router := app.Router

	router.Group(func(r chi.Router) {
		r.Get("/", v1.GetListUser(app))
	})
}
