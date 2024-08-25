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

func NewLogoutUsecase(userRepository entities.UserRepository, timeout time.Duration) entities.LogOutUsecase {
	return &logoutUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *logoutUsecase) LogOut(c context.Context, refreshid string)error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	err := lu.userRepository.DeleteRefreshToken(ctx, refreshid)

	return err
}
