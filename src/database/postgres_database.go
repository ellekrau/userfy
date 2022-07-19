package database

import (
	"database/sql"
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	_ "github.com/lib/pq"
	"log"
)

func openPostgresDatabase() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	var err error
	if database, err = sql.Open("postgres", connectionString); err != nil {
		log.Fatalln(errDatabaseConnection, err.Error())
	}
}
