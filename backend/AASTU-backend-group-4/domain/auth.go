package domain

import (
	"context"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateAccessToken(ctx context.Context, user User) (string, error)
	GenerateAndStoreRefreshToken(ctx context.Context, user User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*RefreshResponse, error)
	DeleteRefreshToken(ctx context.Context, userID string) error
	GeneratePasswordResetToken(ctx context.Context, email string) (string, error)
	ValidateResetToken(ctx context.Context, token string) (string, error)
	InvalidateResetToken(ctx context.Context, token string) error
}

type RefreshRequest struct {
	AccessToken string `json:"accessToken"`
}

type RefreshResponse struct {
	AccessToken string `json:"accessToken"`
}
