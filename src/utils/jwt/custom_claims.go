package jwt

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	Key string
	jwt.StandardClaims
}
