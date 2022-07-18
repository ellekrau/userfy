package http_handler

import (
	"github.com/ellekrau/mercafacil/server/middlewares/auth/jwt"
	"github.com/ellekrau/mercafacil/use-case/get-jwt-token/http-handler/contracts"
	customerror "github.com/ellekrau/mercafacil/utils/custom-error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateJWT(c *gin.Context) {
	// TODO add context timeout

	jwtToken, err := jwt.GenerateJWTToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, customerror.CustomError{
			Code:    "generate_token",
			Message: "Generate token error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, contracts.NewResponse(jwtToken))
}
