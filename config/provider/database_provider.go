package provider

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDbProvider(env *EnvProvider) *sql.DB {
	db, error := sql.Open("postgres", env.databaseUrl)

	if error != nil {
		log.Fatal("Unable to connect to database")
	}

	return db
}
