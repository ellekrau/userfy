package database

import (
	"database/sql"
	"github.com/ellekrau/mercafacil/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

var database *sql.DB

var errDatabaseConnection = "database connection error: "
var errInvalidDatabaseConnection = "invalid database name"

func GetDatabase() *sql.DB {
	return database
}

func StartDatabase() {
	switch strings.ToLower(config.Database.DB) {
	case "postgres":
		StartPostgresDatabase()
		return
	case "mysql":
		StartMySQLDatabase()
		return
	default:
		log.Fatalln(errDatabaseConnection, errInvalidDatabaseConnection, " ", config.Database.DB)
	}
}
