package createuserhttphandler

import (
	"fmt"
	"github.com/ellekrau/mercafacil/use-case/create-user/http-handler/contracts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHttpHandler(c *gin.Context) {
	request, err := contracts.CreateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, contracts.NewResponse(err.Error()))
		return
	}
	fmt.Sprint(request)

	c.JSON(http.StatusNotImplemented, contracts.NewResponse("Hello world!"))
}
