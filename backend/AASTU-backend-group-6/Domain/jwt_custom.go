package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	UserName string `json:"user_name"`
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
