package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func openMySQLDatabase() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.OldDatabase.User, config.OldDatabase.Password, config.OldDatabase.Host, config.OldDatabase.Port, config.OldDatabase.Name)

	var err error
	if database, err = sql.Open("mysql", connectionString); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}
}

func openMySQLDatabaseWithReturn(dbConfig config.Database) *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	var connection *sql.DB
	var err error

	if connection, err = sql.Open("mysql", connectionString); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}

	return connection
}
