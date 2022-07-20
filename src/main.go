package main

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/server"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	loadEnvironmentVariables()
	config.LoadServiceConfig()
	config.LoadClientsConfig()

	database.LoadDatabaseClientKeys()
	server.RunServer()
}

func loadEnvironmentVariables() {
	if err := godotenv.Load("./../.env"); err != nil {
		log.Fatalln("error in load .env")
	}
}
