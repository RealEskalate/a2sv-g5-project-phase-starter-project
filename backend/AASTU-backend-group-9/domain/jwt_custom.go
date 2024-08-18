package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Email    string `json:"email"`
	Username   string `json:"username"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	Username string `json:"usernamae"`
	jwt.StandardClaims
}