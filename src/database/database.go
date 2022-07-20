package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	"strings"
)

var (
	connections map[string]*sql.DB
)

func LoadDatabaseClientKeys() {
	connections = make(map[string]*sql.DB)
	for _, client := range config.GetClientsConfig().Clients {
		connections[client.Key] = nil
	}
}

func StartDatabaseByClientKey(clientKey string) (err error) {
	var client config.Client
	if client, err = config.GetClient(clientKey); err != nil {
		return fmt.Errorf("start database by client error: %v", err)
	}

	var connection *sql.DB
	switch strings.ToLower(client.Database.Database) {
	case "postgres":
		connection, err = openPostgresDatabaseWithReturn(client.Database)
		break
	case "mysql":
		connection, err = openMySQLDatabase(client.Database)
		break
	default:
		return fmt.Errorf("start database by client error: client key['%s'] not configured", clientKey)
	}

	if err != nil {
		return err
	}

	if err = connection.Ping(); err != nil {
		return fmt.Errorf("start databse by client key error: database ping error key['%s'] err: %v", clientKey, err)
	}

	connections[client.Key] = connection
	return nil
}

func GetDatabaseByClientKey(clientKey string) (*sql.DB, error) {
	if connections[clientKey] != nil {
		return connections[clientKey], nil
	}

	if err := StartDatabaseByClientKey(clientKey); err != nil {
		return nil, fmt.Errorf("get database by client key error key['%s'] error: %v", clientKey, err)
	}

	return connections[clientKey], nil
}
