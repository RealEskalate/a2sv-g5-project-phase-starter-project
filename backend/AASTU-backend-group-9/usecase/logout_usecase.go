package usecase

import (
	"blog/domain"
	"context"
	"errors"
	"time"
)

type LogoutUsecase struct {
	tokenRepository domain.TokenRepository
	contextTimeout  time.Duration
}



func NewLogoutUsecase(tokenRepository domain.TokenRepository, timeout time.Duration) domain.LogoutUsecase {
	return &LogoutUsecase{
		tokenRepository: tokenRepository,
		contextTimeout: timeout,
	}
}

func (lu *LogoutUsecase) Logout(cxt context.Context, refreshToken string) error {
	token, err := lu.tokenRepository.FindTokenByRefreshToken(cxt, refreshToken)
	if err != nil {
		return errors.New("token not found")
	}

	err = lu.tokenRepository.DeleteToken(cxt, token.ID)
	if err != nil {
		return errors.New("failed to delete token")
	}
	return nil
}