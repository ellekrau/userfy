package main

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/database/database-config"
	"github.com/ellekrau/mercafacil/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//godotenv.Load("./../.env")
	config.LoadEnvironmentVariables()

	database.StartDatabase()

	server.RunServer()
}
