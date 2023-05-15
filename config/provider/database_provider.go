package provider

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
)

func NewDbProvider(env *EnvProvider) *sql.DB {
	db, error := sql.Open("postgres", env.databaseUrl)

	if error != nil {
		log.Fatal("Unable to connect to database")
	}

	return db
}

func NewTestDbProvider(env *EnvProvider) *sql.DB {
	txdb.Register("txdb", "postgresql", env.databaseUrl)
	db, error := sql.Open("txdb", "TestTransactionDB")

	if error != nil {
		log.Fatal("Unable to connect to database")
	}

	return db
}
