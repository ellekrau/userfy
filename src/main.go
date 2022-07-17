package main

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.LoadEnvironmentVariables()

	database.StartDatabase()

	server.RunServer()
}
