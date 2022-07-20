package database

import (
	"database/sql"
	"fmt"
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

func StartDatabaseByClientKey(clientKey string) error {
	client, err := config.GetClient(clientKey)
	if err != nil {
		return fmt.Errorf("start database by client error: %v", err)
	}

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

	if err = connection.Ping(); err != nil {
		return fmt.Errorf("database ping error key['%s'] err: %v", clientKey, err)
	}

	connections[client.Key] = connection
	return nil
}

func GetDatabaseByClientKey(clientKey string) (*sql.DB, error) {
	if connections[clientKey] != nil {
		return connections[clientKey], nil
	}

	if err := StartDatabaseByClientKey(clientKey); err != nil {
		return nil, fmt.Errorf("get databasy by client key error key['%s'] error: %v", clientKey, err)
	}

	return connections[clientKey], nil
}
