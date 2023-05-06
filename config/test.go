package config

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (app *App) MigrateDbForTest() {
	driver, err := postgres.WithInstance(app.db, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+app.rootDir+"/db/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}
