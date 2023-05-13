package provider

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDbProvider(dbUrl string) *sql.DB {
	db, error := sql.Open("postgres", dbUrl)

	if error != nil {
		log.Fatal("Unable to connect to database")
	}

	return db
}
