package config

import (
	"database/sql"
	"log"
	// "measure/webserver/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	env     AppEnv
	envVars *EnvProvider
	router  *chi.Mux
	db      *sql.DB
}

func (app *App) Env() AppEnv {
	return app.env
}

func (app *App) EnvVars() *EnvProvider {
	return app.envVars
}

func (app *App) Router() *chi.Mux {
	return app.router
}

func (app *App) DB() *sql.DB {
	return app.db
}

func (app *App) setupRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r = router.Setup(r)

	app.router = r
}

func (app *App) setupEnv(appEnv AppEnv) {
	app.env = appEnv
	app.envVars = NewEnvProvider(app.env)
}

func (app *App) Start() {
  servePath := ":" + app.EnvVars().ServerPort()
  log.Print("server listening on " + servePath)
	http.ListenAndServe(servePath, app.router)
}

func NewApp(appEnv AppEnv) *App {
	app := App{}
	app.setupEnv(appEnv)
	app.setupRouter()
	return &app
}
