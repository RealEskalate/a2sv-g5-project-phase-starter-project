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

// UpdateRefreshToken implements entities.LoginUsecase.

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
func (lu *loginUsecase) CreateRefreshData(c context.Context, refreshData entities.RefreshData) error {
	return lu.CreateRefreshData(c, refreshData)
}

func (lu *loginUsecase) CreateAccessToken(user *entities.User, secret string, expiry int,refreshDataId string) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry,refreshDataId)
}

func (lu *loginUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int,refreshDataId string) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry,refreshDataId)
}
