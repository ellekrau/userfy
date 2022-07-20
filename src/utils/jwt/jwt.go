package jwt

import (
	"errors"
	"github.com/ellekrau/mercafacil/config"
	"github.com/golang-jwt/jwt"
)

var errInvalidToken = errors.New("invalid token error")

func GenerateJWTToken(key string) (string, error) {
	c := CustomClaims{
		Key: key,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Authentication.JWTKey))
}

func ValidateJWTToken(token string) (CustomClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Authentication.JWTKey), nil
	}

	customClaims := &CustomClaims{}
	extractedToken, err := jwt.ParseWithClaims(token, customClaims, keyFunc)
	if err != nil {
		return CustomClaims{}, errInvalidToken
	}

	if !extractedToken.Valid {
		return CustomClaims{}, errInvalidToken
	}

	customClaims, ok := extractedToken.Claims.(*CustomClaims)
	if !ok {
		return CustomClaims{}, errInvalidToken
	}

	return *customClaims, nil
}
