package main

import (
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/server"
	"github.com/ellekrau/mercafacil/utils/security"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.LoadEnvironmentVariables()

	e, _ := security.Encrypt("abacate")
	fmt.Println(e)

	d, _ := security.Decrypt(e)
	fmt.Println(d)

	server.RunServer()
}
