package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var database *sql.DB

var errDatabaseConnection = "database connection error: "

func GetDatabase() *sql.DB {
	return database
}

func StartDatabase() {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	if database, err = sql.Open(config.Database.DB, connectionString); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}

	if err = database.Ping(); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}
}
