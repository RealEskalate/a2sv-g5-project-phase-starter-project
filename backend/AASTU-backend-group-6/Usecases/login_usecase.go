package usecases

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
func (lu *loginUsecase) SaveAsActiveUser(user domain.ActiveUser, refreshToken string, c context.Context) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	_, err := lu.activeUserRepository.FindActiveUser(user.ID.Hex(), user.UserAgent, ctx)
	if err == nil {
		return errors.New("User already logged in")
	}
	return lu.activeUserRepository.CreateActiveUser(user, ctx)
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.FindUserByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp), // Convert expiration time to *jwt.NumericDate
		},
	}
	return infrastructure.CreateToken(claims, secret)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))

	// Create claims
	claims := &domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp), // Convert expiration time to *jwt.NumericDate
		},
	}
	return infrastructure.CreateToken(claims, secret)
}
