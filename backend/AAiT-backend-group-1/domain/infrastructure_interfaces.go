package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtService interface {
	GenerateAccessTokenWithPayload(user User) (string,  Error)
	GenerateRefreshTokenWithPayload(user User) (string,  Error)
	GenerateVerificationToken(user User) (string,  Error)
	GenerateResetToken(email string) (string,  Error)
	ValidateVerificationToken(token string) (*jwt.Token,  Error)
	ValidateAccessToken(token string) (*jwt.Token,  Error)
	ValidateRefreshToken(token string) (*jwt.Token,  Error)
	ValidateResetToken(token string) (*jwt.Token,  Error)
	RevokedToken(token string) Error
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
	SendPasswordResetEmail(to, name, resetLink, resetCode string) error
}

