package database

import (
	"database/sql"
	"fmt"
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	_ "github.com/go-sql-driver/mysql"
)

func openMySQLDatabase(dbConfig clientconfig.Database) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	var connection *sql.DB
	var err error

	if connection, err = sql.Open("mysql", connectionString); err != nil {
		return nil, fmt.Errorf("open MySQL database connection error: %v", err)
	}

	return connection, nil
}
