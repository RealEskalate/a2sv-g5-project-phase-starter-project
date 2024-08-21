package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtService interface {
	GenerateAccessTokenWithPayload(user User) (string, error)
	GenerateRefreshTokenWithPayload(user User) (string, error)
	GenerateVerificationToken(user User) (string, error)
	GenerateResetToken(email string) (string, error)
	ValidateVerificationToken(token string) (*jwt.Token, error)
	ValidateAccessToken(token string) (*jwt.Token, error)
	ValidateRefreshToken(token string) (*jwt.Token, error)
	ValidateResetToken(token string) (*jwt.Token, error)
	RevokedToken(token string) error
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) (bool, error)
}

type MiddlewareService interface {
	Authenticate() gin.HandlerFunc
	Authorize(role ...string) gin.HandlerFunc
}

type CacheService interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

type EmailService interface {
	SendMail(to, subject, templateName string, body interface{}) error
	SendVerificationEmail(to, name, verificationLink string) error
	SendPasswordResetEmail(to, name, resetLink string) error
}