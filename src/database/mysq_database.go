package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func openMySQLDatabase() {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	if database, err = sql.Open("mysql", connectionString); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}
}
