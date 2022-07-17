package authmiddleware

import (
	"github.com/ellekrau/mercafacil/server/middlewares/authentication/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const invalidTokenMessage = "invalid token"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		splitToken := strings.Split(c.GetHeader("Authorization"), " ")

		// Token must have 2 pieces, like Bearer some.chars.here
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, newResponse(invalidTokenMessage))
			return
		}

		// Token must start with Bearer property
		if splitToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, newResponse(invalidTokenMessage))
			return
		}

		if err := jwt.ValidateJWTToken(splitToken[1]); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, newResponse(err.Error()))
			return
		}

		c.Next()
	}
}
