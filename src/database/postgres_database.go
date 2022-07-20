package database

import (
	"database/sql"
	"fmt"
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	_ "github.com/lib/pq"
)

func openPostgresDatabaseWithReturn(dbConfig clientconfig.Database) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	var connection *sql.DB
	var err error

	if connection, err = sql.Open("postgres", connectionString); err != nil {
		return nil, fmt.Errorf("open MySQL database connection error: %v", err)
	}

	return connection, nil
}
