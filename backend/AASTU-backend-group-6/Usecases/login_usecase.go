package usecases

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"context"
	"time"
)

type loginUsecase struct {
	userRepository       domain.UserRepository
	activeUserRepository domain.ActiveUserRepository
	contextTimeout       time.Duration
}

// SaveAsActiveUser implements domain.LoginUsecase.
func NewLoginUsecase(userRepository domain.UserRepository, activeUserRepository domain.ActiveUserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository:       userRepository,
		activeUserRepository: activeUserRepository,
		contextTimeout:       timeout,
	}
}
func (lu *loginUsecase) SaveAsActiveUser(user domain.User, refreshToken string, c context.Context) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.activeUserRepository.CreateActiveUser(domain.ActiveUser{
		ID:       user.ID,
		RefreshToken: refreshToken,
	}, ctx)
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.FindUserByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return infrastructure.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return infrastructure.CreateRefreshToken(user, secret, expiry)
}
