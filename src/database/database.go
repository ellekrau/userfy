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
	switch strings.ToLower(config.OldDatabase.DB) {
	case "postgres":
		openPostgresDatabase()
		break
	case "mysql":
		openMySQLDatabase()
		break
	default:
		log.Fatalln(errDatabaseConnection, errInvalidDatabaseConnection, config.OldDatabase.DB)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}
}

var connections map[string]*sql.DB

func StartDatabases() {
	connections = make(map[string]*sql.DB)

	for _, client := range config.GetClientsConfig().Clients {
		var connection *sql.DB
		switch strings.ToLower(client.Database.Database) {
		case "postgres":
			connection = openPostgresDatabaseWithReturn(client.Database)
			break
		case "mysql":
			connection = openMySQLDatabaseWithReturn(client.Database)
			break
		default:
			log.Fatalln("") // TODO
		}

		if err := connection.Ping(); err != nil {
			log.Fatalln(errDatabaseConnection, err.Error()) // TODO improve error message
		}

		connections[client.Key] = connection
	}
}

func GetDatabaseByKey(key string) *sql.DB {
	return connections[key]
}
