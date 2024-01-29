package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("sdgsdgasegesg")

type JWTclaim struct {
	Username string
	UserID   string
	jwt.RegisteredClaims
}
