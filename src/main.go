package main

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/server"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.LoadConfig()
	database.LoadDatabaseClientKeys()

	godotenv.Load("./../.env")

	config.LoadEnvironmentVariables()
	server.RunServer()
}
