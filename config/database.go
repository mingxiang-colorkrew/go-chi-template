package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func (app *App) setupDb() {
	db, error := sql.Open("postgres", app.env.databaseUrl)

	if error != nil {
		log.Fatal("Unable to connect to database")
	}

	app.db = db
}
