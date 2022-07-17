package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	_ "github.com/lib/pq"
	"log"
)

var errInvalidDBPortValue = "invalid db port value"

func StartPostgresDatabase() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	var err error
	if database, err = sql.Open(config.Database.DB, connectionString); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}

	if err = database.Ping(); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}
}
