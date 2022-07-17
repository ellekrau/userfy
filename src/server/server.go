package server

import (
	"fmt"
	"github.com/ellekrau/mercafacil/config"
	authmiddleware "github.com/ellekrau/mercafacil/server/middlewares/authentication"
	createuserhttphandler "github.com/ellekrau/mercafacil/use-case/create-user/http-handler"
	generatejwthttphandler "github.com/ellekrau/mercafacil/use-case/generate-jwt/http-handler"
	"github.com/gin-gonic/gin"
	"log"
)

func RunServer() {
	// Logger and recovery middlewares attached by default
	router := gin.Default()

	// Simple JWT generator
	// Created only for testing
	router.GET("/token", generatejwthttphandler.GenerateJWT)

	router.POST("/user", authmiddleware.Auth(), createuserhttphandler.CreateUserHttpHandler)

	if err := router.Run(fmt.Sprintf(":%s", config.Service.Port)); err != nil {
		log.Fatalln("error starting server: ", err.Error())
	}
}
