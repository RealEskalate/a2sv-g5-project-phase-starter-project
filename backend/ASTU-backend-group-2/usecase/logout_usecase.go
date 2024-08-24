package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
)

type logoutUsecase struct {
	userRepository entities.UserRepository
	contextTimeout time.Duration
}

func NewLogoutUsecase(userRepository entities.UserRepository, timeout time.Duration) entities.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *logoutUsecase) GetUserByEmail(c context.Context, email string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	user, err := lu.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return entities.User{}, err
	}
	return *user, nil
}

func (lu *logoutUsecase) RevokeRefreshToken(c context.Context, userID string, refreshToken string) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.RevokeRefreshToken(ctx, userID, refreshToken)
}
