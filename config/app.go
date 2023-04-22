package config

import (
	"database/sql"
	"path"
	"runtime"
)

type App struct {
	env     *EnvProvider
	db      *sql.DB
	rootDir string
}

func (app *App) EnvVars() *EnvProvider {
	return app.env
}

func (app *App) DB() *sql.DB {
	return app.db
}

func (app *App) setEnv() {
	app.env = NewEnvProvider(app.rootDir)
}

func (app *App) setRootDir() {
	_, b, _, _ := runtime.Caller(0)
	app.rootDir = path.Join(path.Dir(b), "..")
}

func NewApp() *App {
	app := App{}

	app.setRootDir()
	app.setEnv()
	app.setupDb()

	return &app
}
