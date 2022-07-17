package createuserhttphandler

import (
	"github.com/ellekrau/mercafacil/use-case/create-user/http-handler/contracts"
	createuserservices "github.com/ellekrau/mercafacil/use-case/create-user/services"
	createuserservicecontracts "github.com/ellekrau/mercafacil/use-case/create-user/services/contracts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHttpHandler(c *gin.Context) {
	request, err := contracts.CreateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, contracts.NewResponse(err.Error()))
		return
	}

	serviceInput := createuserservicecontracts.NewCreateUserServiceInput(request.Name, request.Cellphone)
	createuserservices.CreateUser(c, serviceInput)

	c.JSON(http.StatusNotImplemented, contracts.NewResponse("Hello world!"))
}
