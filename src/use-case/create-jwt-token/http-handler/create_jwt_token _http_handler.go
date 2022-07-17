package http_handler

import (
	"github.com/ellekrau/mercafacil/server/middlewares/authentication/jwt"
	"github.com/ellekrau/mercafacil/use-case/create-jwt-token/http-handler/contracts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateJWT(c *gin.Context) {
	// TODO add context timeout

	jwtToken, err := jwt.GenerateJWTToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, contracts.NewResponse("Generate token error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, contracts.NewResponse(jwtToken))
}
