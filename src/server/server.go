package server

import (
	"fmt"
	createuserhttphandler "github.com/ellekrau/mercafacil/use-case/create-user/http-handler"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func RunServer() {
	// Logger and recovery middlewares attached by default
	router := gin.Default()

	router.POST("/user", createuserhttphandler.CreateUserHttpHandler)

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
