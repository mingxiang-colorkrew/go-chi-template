package config

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (app *App) MigrateDbForTest() {
	if app.env != AppEnvTest {
		log.Fatal("Should not migrate DB in non-test env")
	}
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

func (app *App) MockServerForTest() *chi.Mux {
	if app.env != AppEnvTest {
		log.Fatal("Should not mock server in non-test env")
	}

	return app.router
}
