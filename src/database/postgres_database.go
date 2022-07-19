package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	_ "github.com/lib/pq"
	"log"
)

func openPostgresDatabase() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.OldDatabase.User, config.OldDatabase.Password, config.OldDatabase.Host, config.OldDatabase.Port, config.OldDatabase.Name)

	var err error
	if database, err = sql.Open("postgres", connectionString); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}
}

func openPostgresDatabaseWithReturn(dbConfig config.Database) *sql.DB {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	var connection *sql.DB
	var err error

	if connection, err = sql.Open("postgres", connectionString); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}

	return connection
}
