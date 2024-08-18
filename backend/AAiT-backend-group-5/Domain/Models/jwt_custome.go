package models

import "github.com/dgrijalva/jwt-go"

type JWTCustome struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
