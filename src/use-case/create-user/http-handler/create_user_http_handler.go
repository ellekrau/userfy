package createuserhttphandler

import (
	"github.com/ellekrau/mercafacil/use-case/create-user/http-handler/contracts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHttpHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, contracts.NewResponse("Hello world!"))
}
