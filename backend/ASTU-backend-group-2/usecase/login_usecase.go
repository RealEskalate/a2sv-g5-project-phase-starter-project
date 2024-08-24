package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
)

type loginUsecase struct {
	userRepository entities.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository entities.UserRepository, timeout time.Duration) entities.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	user, err := lu.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return entities.User{}, err
	}
	return *user, nil
}
func (lu *loginUsecase) UpdateRefreshToken(c context.Context, userID string, refreshToken string) error {
	err := lu.userRepository.UpdateRefreshToken(c, userID, refreshToken)
	if err != nil {
		return err
	}
	return nil
}

func (lu *loginUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
