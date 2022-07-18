package main

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/server"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	godotenv.Load("./../.env")
	config.LoadEnvironmentVariables()
	database.StartDatabase()
	server.RunServer()
}
