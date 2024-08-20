package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtService interface {
	GenerateAccessTokenWithPayload(user User) (string, error)
	GenerateRefreshTokenWithPayload(user User) (string, error)
	ValidateAccessToken(token string) (*jwt.Token, error)
	ValidateRefreshToken(token string) (*jwt.Token, error)
	RevokedToken(token string) error
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) (bool, error)
}

type MiddlewareService interface {
	Authenticate() gin.HandlerFunc
	Authorize() gin.HandlerFunc
}

