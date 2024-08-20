package utils

import (
	"Blog_Starter/domain"

	"github.com/golang-jwt/jwt"
)

// JwtCustomClaims defines the custom claims for the access token.
type JwtCustomClaims struct {
    UserID string `json:"user_id"`
    Email  string `json:"email"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

// JwtCustomRefreshClaims defines the custom claims for the refresh token.
type JwtCustomRefreshClaims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

type TokenManager interface {
    CreateAccessToken(user *domain.User, secret string, expiry int) (string, error)
    CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error)
    ExtractIDFromToken(requestToken string, secret string) (string, error)
    ExtractRoleFromToken(requestToken string, secret string) (string, error)
}

