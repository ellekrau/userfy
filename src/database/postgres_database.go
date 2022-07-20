package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	_ "github.com/lib/pq"
)

func openPostgresDatabaseWithReturn(dbConfig config.Database) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	var connection *sql.DB
	var err error

	if connection, err = sql.Open("postgres", connectionString); err != nil {
		return nil, fmt.Errorf("open MySQL database connection error: %v", err)
	}

	return connection, nil
}
