package authmiddleware

import (
	"fmt"
	customerror "github.com/ellekrau/mercafacil/utils/custom-error"
	"github.com/ellekrau/mercafacil/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var errInvalidTokenCode = "invalid_token"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		splitToken := strings.Split(c.GetHeader("Authorization"), " ")

		// Token must have 2 pieces, like Bearer some.chars.here
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customerror.CustomError{
				Code:    errInvalidTokenCode,
				Message: "Invalid token",
			})
			return
		}

		// Token must start with Bearer property
		if splitToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customerror.CustomError{
				Code:    errInvalidTokenCode,
				Message: "Invalid token. Must be 'Bearer' token.",
			})
			return
		}

		if customClaims, err := jwt.ValidateJWTToken(splitToken[1]); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customerror.CustomError{
				Code:    errInvalidTokenCode,
				Message: fmt.Sprint("Invalid token. Error: ", err.Error()),
			})
			return
		} else {
			c.Set("key", customClaims.Key)
		}

		c.Next()
	}
}
