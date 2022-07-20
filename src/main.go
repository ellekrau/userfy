package main

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/server"
)

func main() {
	config.LoadServiceConfig()
	config.LoadClientsConfig()

	database.LoadDatabaseClientKeys()
	server.RunServer()
}
