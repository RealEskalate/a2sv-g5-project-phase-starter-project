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

func (lu *LogoutUsecase) Logout(cxt context.Context, refreshToken string, deviceFingerprint string) error {
	ctx, cancel := context.WithTimeout(cxt, lu.contextTimeout)
	defer cancel()

	token, err := lu.tokenRepository.FindTokenByRefreshToken(ctx, refreshToken)
	if err != nil {
		return errors.New("token not found")
	}

	// Check if the device fingerprint matches
	if token.DeviceFingerprint != deviceFingerprint {
		return errors.New("device fingerprint does not match")
	}

	err = lu.tokenRepository.DeleteToken(ctx, token.ID)
	if err != nil {
		return errors.New("failed to delete token")
	}

	return nil
}