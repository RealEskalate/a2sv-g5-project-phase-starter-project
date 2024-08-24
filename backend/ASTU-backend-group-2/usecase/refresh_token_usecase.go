package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
)

type refreshTokenUsecase struct {
	userRepository entities.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository entities.UserRepository, timeout time.Duration) entities.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, id string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	user, err := rtu.userRepository.GetUserById(ctx, id)
	if err != nil {
		return entities.User{}, err
	}
	return *user, nil
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	userClaims, err := tokenutil.ExtractUserClaimsFromToken(requestToken, secret)
	if err != nil {
		return "", err
	}
	return userClaims["id"].(string), nil
}
