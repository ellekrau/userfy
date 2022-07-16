package main

import (
	"github.com/ellekrau/mercafacil/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.RunServer()
}
