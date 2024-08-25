package usecase

import (
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type RefreshTokenUsecase struct {
	UserRepository domain.UserRepository
	TokenManager   utils.TokenManager
	ContextTimeout time.Duration
	Env            *config.Env
}

// CheckRefreshToken implements domain.RefreshTokenUsecase.

func (r *RefreshTokenUsecase) CheckRefreshToken(ctx context.Context, userID string, refreshToken string) error {
	ctx, cancel := context.WithTimeout(ctx, r.ContextTimeout)
	defer cancel()

	// Fetch user from repository
	user, err := r.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Check if the stored refresh token matches the provided one
	if user.RefreshToken != refreshToken {
		return errors.New("refresh token is not correct")
	}

	// Parse the JWT token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm used for signing the token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return the secret key used for signing the token
		return []byte(r.Env.RefreshTokenSecret), nil 
	})

	if err != nil {
		return err
	}

	// Extract the claims from the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration time
		exp := int64(claims["exp"].(float64))
		if exp < time.Now().Unix() {
			return errors.New("refresh token has expired")
		}
	} else {
		return errors.New("invalid token claims")
	}

	return nil
}

// RefreshToken implements domain.RefreshTokenUsecase.
func (r *RefreshTokenUsecase) UpdateTokens(ctx context.Context, userID string) (*domain.RefreshTokenResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, r.ContextTimeout)
	defer cancel()

	user, err := r.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	accessToken, err := r.TokenManager.CreateAccessToken(user, r.Env.AccessTokenSecret, r.Env.AccessTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := r.TokenManager.CreateRefreshToken(user, r.Env.RefreshTokenSecret, r.Env.RefreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	userId := user.UserID.Hex()
	user, err = r.UserRepository.UpdateToken(ctx, accessToken, refreshToken, userId)
	if err != nil {
		return nil, err
	}

	var refreshTokenResponse domain.RefreshTokenResponse
	refreshTokenResponse.AccessToken = user.AccessToken
	refreshTokenResponse.RefreshToken = user.RefreshToken

	return &refreshTokenResponse, nil

}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, tokenManager utils.TokenManager, timeout time.Duration, env *config.Env) domain.RefreshTokenUsecase {
	return &RefreshTokenUsecase{
		UserRepository: userRepository,
		TokenManager:   tokenManager,
		ContextTimeout: timeout,
		Env:            env,
	}
}
