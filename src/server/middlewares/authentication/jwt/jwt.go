package jwt

import (
	"errors"
	"github.com/ellekrau/mercafacil/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, nil)
	return token.SignedString([]byte(config.Authentication.Key))
}

func ValidateJWTToken(token string) error {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Authentication.Key), nil
	}

	_, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return errors.New("invalid token")
	}

	return nil
}
