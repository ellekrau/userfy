package main

import (
	"github.com/ellekrau/userfy/config"
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"github.com/ellekrau/userfy/database"
	"github.com/ellekrau/userfy/server"
)

func main() {
	config.LoadServiceConfig()
	clientconfig.LoadClientsConfig()

	database.LoadClientDBConnections()

	server.RunServer()
}
