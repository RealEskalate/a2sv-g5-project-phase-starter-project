package usecase

import (
	"blog/domain"
	"blog/internal/tokenutil"
	"blog/internal/userutil"
	"context"
	"errors"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) AuthenticateUser(c context.Context, login *domain.AuthLogin) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	user, err := lu.userRepository.GetUserByEmail(ctx, login.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = userutil.ComparePassword(user.Password, login.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateAccessToken(&domain.AuthSignup{Email: user.Email, Username: user.Username}, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateRefreshToken(&domain.AuthSignup{Username: user.Username}, secret, expiry)
}
