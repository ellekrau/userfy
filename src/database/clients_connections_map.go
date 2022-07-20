package database

import (
	"database/sql"
	"sync"
)

type clientsDBConnectionsThreadSafeMap struct {
	dbConnections map[string]*sql.DB
	sync.Map
}

func (c *clientsDBConnectionsThreadSafeMap) getByClientKey(clientKey string) (*sql.DB, error) {
	// Used to avoid race condition in map read/write
	db, ok := c.Load(clientKey)
	if !ok {
		dbConnection, err := startDatabaseByClientKey(clientKey)
		if err != nil {
			return nil, err
		}
		// Thread safely adds clientKey and DB connection to the map
		c.Store(clientKey, dbConnection)
		return dbConnection, nil
	}

	return db.(*sql.DB), nil
}
