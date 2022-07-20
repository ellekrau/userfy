package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config/client-config"
	"strings"
)

var (
	clientsDBConnections map[string]*sql.DB
)

func LoadClientsDatabaseConnectionKeys() {
	clientsDBConnections = make(map[string]*sql.DB)
	for _, client := range clientconfig.GetClientsConfig().Clients {
		clientsDBConnections[client.Key] = nil
	}
}

func StartDatabaseByClientKey(clientKey string) (err error) {
	var client clientconfig.Client
	if client, err = clientconfig.GetClient(clientKey); err != nil {
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

	clientsDBConnections[client.Key] = connection
	return nil
}

func GetDatabaseByClientKey(clientKey string) (*sql.DB, error) {
	if clientsDBConnections[clientKey] != nil {
		return clientsDBConnections[clientKey], nil
	}

	if err := StartDatabaseByClientKey(clientKey); err != nil {
		return nil, fmt.Errorf("get database by client key error key['%s'] error: %v", clientKey, err)
	}

	return clientsDBConnections[clientKey], nil
}
