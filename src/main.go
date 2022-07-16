package main

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	config.LoadEnvironmentVariables()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.RunServer()
}
