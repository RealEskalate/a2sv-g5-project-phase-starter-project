package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaim struct {
	UserID  string `json:"userId"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Isadmin bool   `json:"isadmin"`
	Exp     int64  `json:"exp"`
	jwt.StandardClaims
}
