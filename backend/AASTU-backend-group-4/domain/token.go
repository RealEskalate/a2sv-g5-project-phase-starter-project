package domain

import (
	"context"
	"time"
)

type RefreshTokenRepository interface {
	StoreRefreshToken(ctx context.Context, userID string, tokenString string, expiresAt time.Time) error
	GetRefreshToken(ctx context.Context, userID string) (string, error)
	DeleteRefreshToken(ctx context.Context, userID string) error
}

type PasswordResetToken struct {
	Token  string    `json:"token"`
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
	Used   bool      `json:"used"`
}

type ResetTokenRepository interface {
	StoreResetToken(ctx context.Context, token PasswordResetToken) error
	ValidateResetToken(ctx context.Context, token string) (string, error) // returns email
	InvalidateResetToken(ctx context.Context, token string) error
}
