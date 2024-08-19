package usecases

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"context"
	"time"
)

type refreshTokenUsecase struct {
	userRepository       domain.UserRepository
	activeUserRepository domain.ActiveUserRepository
	contextTimeout       time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

// removeActiveUser implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) RemoveActiveUser(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	return r.activeUserRepository.DeleteActiveUserById(id, ctx)
}

// checkActiveUser implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) CheckActiveUser(c context.Context, id string) (domain.ActiveUser, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	return r.activeUserRepository.FindActiveUserById(id, ctx)
}

// checkActiveUser implements domain.RefreshTokenUsecase.

// CreateAccessToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return infrastructure.CreateAccessToken(user, secret, expiry)
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
	print(ctx)
	return r.userRepository.FindUserByID(ctx, id)
}
