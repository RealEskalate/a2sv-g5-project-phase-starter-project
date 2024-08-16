package usecases

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"context"
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

// CreateAccessToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	panic("unimplemented")
}

// CreateRefreshToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return infrastructure.CreateRefreshToken(user, secret, expiry)
}

// ExtractIDFromToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return infrastructure.ExtractIDFromToken(requestToken, secret)
}

// GetUserByID implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	r.userRepository.GetUserByID(ctx, id)
}
