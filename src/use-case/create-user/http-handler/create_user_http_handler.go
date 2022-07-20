package createuserhttphandler

import (
	"fmt"
	"github.com/ellekrau/userfy/use-case/create-user/http-handler/contracts"
	createuserservices "github.com/ellekrau/userfy/use-case/create-user/services"
	createuserservicecontracts "github.com/ellekrau/userfy/use-case/create-user/services/contracts"
	customerror "github.com/ellekrau/userfy/utils/custom-error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHttpHandler(c *gin.Context) {
	request, err := contracts.CreateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	clientKey, ok := c.Get("key")
	if !ok {
		c.JSON(http.StatusInternalServerError, customerror.CustomError{
			Code:    "invalid_client_key",
			Message: "Error in get client key from request",
		})
		return
	}

	serviceInput := createuserservicecontracts.NewCreateUserServiceInput(clientKey.(string), request.Name, request.Cellphone)
	user, serviceErr := createuserservices.CreateUser(serviceInput)
	if serviceErr != nil {
		c.JSON(http.StatusInternalServerError, serviceErr)
		return
	}

	c.JSON(http.StatusCreated, contracts.NewResponse(fmt.Sprint(user.ID)))
}
