package config

import (
	"measure/webserver/router"
	"net/http"
  "log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  "github.com/joho/godotenv"
)

func StartApp() {
  setupEnv()

	r := chi.NewRouter()
	r = setupMiddleWare(r)
	r = router.Setup(r)

	http.ListenAndServe(":3000", r)
}

func setupMiddleWare(r *chi.Mux) *chi.Mux {
	r.Use(middleware.Logger)

	return r
}

func setupEnv() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}
