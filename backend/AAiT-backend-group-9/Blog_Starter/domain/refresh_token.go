package domain

import (
	"context"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	CheckRefreshToken(ctx context.Context, userID string, refreshToken string) error
	UpdateTokens(ctx context.Context, userID string) (*RefreshTokenResponse, error)
}

