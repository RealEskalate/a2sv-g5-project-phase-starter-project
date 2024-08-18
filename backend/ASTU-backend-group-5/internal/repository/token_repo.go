package repository

import (
	"blogApp/internal/domain"
	"context"
	"time"
)

type TokenRepository interface {
	BlacklistToken(ctx context.Context, token string, tokenType domain.TokenType, expiry time.Time) error
	IsTokenBlacklisted(ctx context.Context, token string, tokenType domain.TokenType) (bool, error)
	RemoveBlacklistedToken(ctx context.Context, token string, tokenType domain.TokenType) error
}
