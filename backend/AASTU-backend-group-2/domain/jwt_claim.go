package domain

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
