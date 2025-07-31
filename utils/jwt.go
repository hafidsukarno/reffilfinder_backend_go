package utils

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Username string
	Role     string
	UserID   uint
	jwt.RegisteredClaims
}
