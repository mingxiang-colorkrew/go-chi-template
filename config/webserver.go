package config

import (
	"fmt"
	"log"
	"measure/oapi"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

type Webserver struct {
	router     *chi.Mux
	serverAddr string
	app        *App
}

func NewWebserver(app *App, handler oapi.StrictServerInterface) *Webserver {
	serverAddr := ":" + app.EnvVars().ServerPort()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	baseUrl := ""
	strictHandler := oapi.NewStrictHandler(handler, []oapi.StrictMiddlewareFunc{})
	oapi.HandlerFromMuxWithBaseURL(strictHandler, r, baseUrl)

	return &Webserver{
		router:     r,
		serverAddr: serverAddr,
	}
}

func (ws *Webserver) Start() {

	log.Print("WebServer listening on " + ws.serverAddr)

	s := &http.Server{
		Handler: ws.router,
		Addr:    ws.serverAddr,
	}

	log.Fatal(s.ListenAndServe())
}

func (ws *Webserver) PrintRoutes() {
	chi.Walk(
		ws.router,
		func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
			return nil
		},
	)
}
