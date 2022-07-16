package server

import (
	createuserhttphandler "github.com/ellekrau/mercafacil/use-case/create-user/http-handler"
	"github.com/gin-gonic/gin"
	"log"
)

func RunServer() {
	// Logger and recovery middlewares attached by default
	router := gin.Default()

	router.POST("/user", createuserhttphandler.CreateUserHttpHandler)

	// TODO Pass port like ":3300"
	// TODO Get port from .env
	if err := router.Run(); err != nil {
		log.Fatalln("run router error: ", err.Error())
	}
}
