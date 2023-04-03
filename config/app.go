package config

import (
	"database/sql"
	"fmt"
	"log"
	"measure/oapi"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	env     AppEnv
	envVars *EnvProvider
	db      *sql.DB
	router  *chi.Mux
}

func (app *App) Env() AppEnv {
	return app.env
}

func (app *App) EnvVars() *EnvProvider {
	return app.envVars
}

func (app *App) DB() *sql.DB {
	return app.db
}

func (app *App) setupEnv(appEnv AppEnv) {
	app.env = appEnv
	app.envVars = NewEnvProvider(app.env)
}

func (app *App) SetupRouter(handler oapi.StrictServerInterface) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

  baseUrl := "/api"
	strictHandler := oapi.NewStrictHandler(handler, []oapi.StrictMiddlewareFunc{})
	oapi.HandlerFromMuxWithBaseURL(strictHandler, r, baseUrl)

	app.router = r
}

func (app *App) Start() {

	serverAddr := ":" + app.EnvVars().ServerPort()
	log.Print("server listening on " + serverAddr)

	s := &http.Server{
		Handler: app.router,
		Addr:    serverAddr,
	}

	log.Fatal(s.ListenAndServe())
}

func (app *App) PrintRoutes() {
	chi.Walk(app.router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})
}

func NewApp(appEnv AppEnv) *App {
	app := App{}

	app.setupEnv(appEnv)
	app.setupDb()

	return &app
}
