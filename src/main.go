package main

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/config/client-config"
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/server"
)

func main() {
	config.LoadServiceConfig()
	clientconfig.LoadClientsConfig()

	database.LoadClientsDatabaseConnectionKeys()
	server.RunServer()
}
