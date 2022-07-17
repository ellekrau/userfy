package createuserhttphandler

import (
	"fmt"
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
	user, serviceErr := createuserservices.CreateUser(serviceInput)
	if serviceErr != nil {
		c.JSON(http.StatusInternalServerError, serviceErr)
		return
	}

	c.JSON(http.StatusOK, contracts.NewResponse(fmt.Sprint(user.ID)))
}
