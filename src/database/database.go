package database

import (
	"database/sql"
	"fmt"
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"strings"
)

var clientsDBConnections clientsDBConnectionsThreadSafeMap

func LoadClientDBConnections() {
	// Used to avoid race condition in map read/write
	clientsDBConnections = clientsDBConnectionsThreadSafeMap{
		dbConnections: make(map[string]*sql.DB),
	}
}

func GetDatabaseByClientKey(clientKey string) (*sql.DB, error) {
	db, err := clientsDBConnections.getByClientKey(clientKey)
	if err != nil {
		return nil, fmt.Errorf("get database by client key error: %v", err)
	}

	return db, nil
}

func startDatabaseByClientKey(clientKey string) (dbConnection *sql.DB, err error) {
	var client clientconfig.Client
	if client, err = clientconfig.GetClient(clientKey); err != nil {
		return nil, fmt.Errorf("start database by client error: client config not found: %v", err)
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
		return nil, fmt.Errorf("start database by client error: client key['%s'] not configured", clientKey)
	}

	if err != nil {
		return nil, err
	}

	if err = connection.Ping(); err != nil {
		return nil, fmt.Errorf("start database by client key error: database ping error key['%s'] err: %v", clientKey, err)
	}

	return connection, nil
}
