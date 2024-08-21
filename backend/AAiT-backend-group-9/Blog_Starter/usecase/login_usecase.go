package usecase

import (
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)


type LoginUseCase struct {
	UserRepository domain.UserRepository
	TokenManager utils.TokenManager
	ContextTimeout  time.Duration
}


func NewLoginUseCase(loginRepository , userRepository domain.UserRepository,tokenManager utils.TokenManager ,timeout time.Duration) domain.LoginUsecase {
	return &LoginUseCase{
		UserRepository: userRepository,
		TokenManager: tokenManager,
		ContextTimeout:  timeout,
	}
}
// Login implements domain.LoginUsecase.
func (l *LoginUseCase) Login(c context.Context, req *domain.UserLogin) (*domain.LoginResponse, error) {
    ctx, cancel := context.WithTimeout(c, l.ContextTimeout)
    defer cancel()
    env := config.NewEnv()
    user, err := l.UserRepository.GetUserByEmail(ctx, req.Email)
    if err != nil {
        return nil, err
    }

    // Compare the hashed password with the plain text password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return nil, fmt.Errorf("password incorrect")
    }

    accessToken, err := l.TokenManager.CreateAccessToken(user, env.AccessTokenSecret, 1)
    if err != nil {
        return nil, err
    }
    refreshToken, err := l.TokenManager.CreateRefreshToken(user, env.RefreshTokenSecret, 24)
    if err != nil {
        return nil, err
    }

    _, err = l.UserRepository.UpdateToken(ctx, accessToken, refreshToken, user.UserID.String())
    if err != nil {
        return nil, err
    }

    var loginResponse domain.LoginResponse
    loginResponse.AccessToken = accessToken
    loginResponse.RefreshToken = refreshToken
    loginResponse.UserID = user.UserID.String()

    return &loginResponse, nil
}
// UpdatePassword implements domain.LoginUsecase.
func (l *LoginUseCase) UpdatePassword(c context.Context, req domain.ChangePasswordRequest, userID string) error {
    ctx, cancel := context.WithTimeout(c, l.ContextTimeout)
    defer cancel()

    // Hash the new password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Update the password in the repository
    _, err = l.UserRepository.UpdatePassword(ctx, string(hashedPassword), userID)
    if err != nil {
        return err
    }

    return nil
}
