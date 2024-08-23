package usecases

import (
	domain "blogs/Domain"
	infrastructure "blogs/Infrastructure"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type refreshTokenUsecase struct {
	userRepository       domain.UserRepository
	activeUserRepository domain.ActiveUserRepository
	contextTimeout       time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, activeUserrepo domain.ActiveUserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository:       userRepository,
		activeUserRepository: activeUserrepo,
		contextTimeout:       timeout,
	}
}

// removeActiveUser implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) RemoveActiveUser(c context.Context, id string, user_agent string) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	return r.activeUserRepository.DeleteActiveUser(id, user_agent, ctx)
}

// checkActiveUser implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) CheckActiveUser(c context.Context, id string, user_agent string) (domain.ActiveUser, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	return r.activeUserRepository.FindActiveUser(id, user_agent, ctx)
}

// checkActiveUser implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
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

// CreateRefreshToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
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

// ExtractIDFromToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return infrastructure.ExtractIDFromToken(requestToken, secret)
}

// GetUserByID implements domain.RefreshTokenUsecase.
func (r *refreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	print(ctx)
	return r.userRepository.FindUserByID(ctx, id)
}
