package config

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type App struct {
	env     AppEnv
	envVars *EnvProvider
	db      *sql.DB
	Router  *chi.Mux
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

func (app *App) setupRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	app.Router = r
}

func (app *App) Start() {
	servePath := ":" + app.EnvVars().ServerPort()
	log.Print("server listening on " + servePath)
	http.ListenAndServe(servePath, app.Router)
}

func NewApp(appEnv AppEnv) *App {
	app := App{}
	app.setupEnv(appEnv)
	app.setupRouter()
	return &app
}
