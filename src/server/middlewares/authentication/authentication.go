package authentication

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateJWTToken() (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute) // TODO add to env
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Authentication.Key))
}

func ValidateJWTToken() {

}
