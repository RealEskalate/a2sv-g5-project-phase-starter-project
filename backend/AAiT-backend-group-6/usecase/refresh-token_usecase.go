package usecase

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/utils"
	"time"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return utils.GenerateAccessToken(user, expiry, secret)
}

func (lu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.GenerateRefreshToken(user, expiry, secret)
}