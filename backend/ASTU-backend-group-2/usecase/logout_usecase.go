package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
)

type logoutUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLogoutUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *logoutUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	user, err := lu.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return *user, nil
}

func (lu *logoutUsecase) RevokeRefreshToken(c context.Context, refreshToken string) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.RevokeRefreshToken(ctx, refreshToken)
}
