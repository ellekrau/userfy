package server

import (
	"fmt"
	"github.com/ellekrau/userfy/config"
	authmiddleware "github.com/ellekrau/userfy/server/middlewares/auth"
	createuserhttphandler "github.com/ellekrau/userfy/use-case/create-user/http-handler"
	generatejwthttphandler "github.com/ellekrau/userfy/use-case/get-jwt-token/http-handler"
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
		log.Fatalln("starting server error: ", err.Error())
	}
}
