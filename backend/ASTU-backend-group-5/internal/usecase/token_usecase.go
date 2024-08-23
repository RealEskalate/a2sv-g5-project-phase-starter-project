package usecase

import (
	"blogApp/internal/domain"
	"blogApp/internal/repository"
	"context"
	"time"
)

type TokenUseCaseInterface interface {
	BlacklistToken(ctx context.Context, token string, tokenType domain.TokenType, expiry time.Time) error
	IsTokenBlacklisted(ctx context.Context, token string, tokenType domain.TokenType) (bool, error)
}

type TokenUsecase struct {
	repo repository.TokenRepository
}

func NewTokenUsecase(repo repository.TokenRepository) *TokenUsecase {
	return &TokenUsecase{repo: repo}
}

func (uc *TokenUsecase) BlacklistToken(ctx context.Context, token string, tokenType domain.TokenType, expiry time.Time) error {
	return uc.repo.BlacklistToken(ctx, token, tokenType, expiry)
}

func (uc *TokenUsecase) IsTokenBlacklisted(ctx context.Context, token string, tokenType domain.TokenType) (bool, error) {
	return uc.repo.IsTokenBlacklisted(ctx, token, tokenType)
}
