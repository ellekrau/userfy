package http_handler

import (
	"github.com/ellekrau/mercafacil/use-case/get-jwt-token/http-handler/contracts"
	"github.com/ellekrau/mercafacil/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateJWT(c *gin.Context) {
	request, requestErr := contracts.CreateRequest(c)
	if requestErr != nil {
		c.JSON(http.StatusBadRequest, requestErr)
		return
	}

	jwtToken, err := jwt.GenerateJWTToken(request.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, contracts.NewResponse(jwtToken))
}
