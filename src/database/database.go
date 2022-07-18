package database

import (
	"database/sql"
	"github.com/ellekrau/mercafacil/config"
	"log"
	"strings"
)

var (
	database                     *sql.DB
	errDatabaseConnection        = "database connection error: "
	errInvalidDatabaseConnection = "invalid database name "
)

func GetDatabase() *sql.DB {
	return database
}

func StartDatabase() {
	switch strings.ToLower(config.Database.DB) {
	case "postgres":
		openPostgresDatabase()
		break
	case "mysql":
		openMySQLDatabase()
		break
	default:
		log.Fatalln(errDatabaseConnection, errInvalidDatabaseConnection, config.Database.DB)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}
}
