package domain

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
