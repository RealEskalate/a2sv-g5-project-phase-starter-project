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
	userRepository  domain.UserRepository
	tokenRepository domain.TokenRepository
	contextTimeout  time.Duration
}

// checkRefreshToken implements domain.LoginUsecase.

func NewLoginUsecase(userRepository domain.UserRepository, tokenRepo domain.TokenRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository:  userRepository,
		tokenRepository: tokenRepo,
		contextTimeout:  timeout,
	}
}

func (lu *loginUsecase) AuthenticateUser(c context.Context, login *domain.AuthLogin) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	user, err := lu.userRepository.GetUserByEmail(ctx, login.Email)
	if err != nil {
		return nil, errors.New("invalid Email")
	}

	err = userutil.ComparePassword(user.Password, login.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateAccessToken(&domain.AuthSignup{Email: user.Email, Username: user.Username, UserID: user.ID}, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateRefreshToken(&domain.AuthSignup{Username: user.Username, Email: user.Email, UserID: user.ID}, secret, expiry)
}

func (lu *loginUsecase) SaveRefreshToken(c context.Context, token *domain.Token) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	return lu.tokenRepository.SaveToken(ctx, token)
}

func (lu *loginUsecase) CheckRefreshToken(c context.Context, refreshToken string) (*domain.Token, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	return lu.tokenRepository.FindTokenByRefreshToken(ctx, refreshToken)

}
