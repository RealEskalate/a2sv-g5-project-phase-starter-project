package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AIServicesInterface interface {
	GenerateContent(topics []string) (string, error)
	ReviewContent(blogContent string) (string, error)
	CleanText(value interface{}) string
	ExtractText(value interface{}) string
	GenerateTrendingTopics(keywords []string) ([]string, error)
}

type JWTServiceInterface interface {
	SignJWTWithPayload(username string, role string, tokenType string, tokenLifeSpan time.Duration) (string, CodedError)
	ValidateAndParseToken(rawToken string) (*jwt.Token, error)
	GetExpiryDate(token *jwt.Token) (time.Time, CodedError)
	GetUsername(token *jwt.Token) (string, CodedError)
	GetRole(token *jwt.Token) (string, CodedError)
	GetTokenType(token *jwt.Token) (string, CodedError)
}

type MailServiceInterface interface {
	SendMail(from string, to string, mailContent string) error
	EmailVerificationTemplate(hostUrl string, username string, token string) string
	PasswordResetTemplate(hostUrl string, username string, token string) string
}
