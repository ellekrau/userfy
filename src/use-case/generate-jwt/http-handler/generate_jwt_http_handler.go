package http_handler

import (
	"github.com/ellekrau/mercafacil/server/middlewares/authentication"
	"github.com/ellekrau/mercafacil/use-case/generate-jwt/http-handler/contracts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateJWT(c *gin.Context) {
	// TODO add context timeout

	jwtToken, err := authentication.GenerateJWTToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, contracts.NewResponse("Generate token error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, contracts.NewResponse(jwtToken))
}
