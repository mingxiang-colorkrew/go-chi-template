package config

import (
	"measure/webserver/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartApp() {
	r := chi.NewRouter()

	r = setupMiddleWare(r)

	r = router.Setup(r)

	http.ListenAndServe(":3000", r)
}

func setupMiddleWare(r *chi.Mux) *chi.Mux {
	r.Use(middleware.Logger)

	return r
}
