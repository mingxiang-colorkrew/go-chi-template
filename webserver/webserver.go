package webserver

import (
	"fmt"
	"go_chi_template/config"
	"go_chi_template/oapi"
	"go_chi_template/webserver/handler"
	"go_chi_template/webserver/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/jwtauth"
)

type Webserver struct {
	router     *chi.Mux
	serverAddr string
	app        *config.App
}

func (ws *Webserver) Router() *chi.Mux {
	return ws.router
}

func NewWebserver(app *config.App) *Webserver {
	handler := handler.NewHandler(app)
	serverAddr := ":" + app.EnvVars().ServerPort()

	r := chi.NewRouter()
	r.Use(middleware.NewLoggerMiddleware(app.Logger()))
	// r.Use(jwtauth.Verifier(app.JWTAuth()))
	// r.Use(middleware.NewAuthMiddleware(app))

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
