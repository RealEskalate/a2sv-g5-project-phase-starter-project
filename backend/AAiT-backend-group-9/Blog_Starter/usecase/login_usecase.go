package usecase

import (
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)


type LoginUseCase struct {
	UserRepository domain.UserRepository
	TokenManager utils.TokenManager
	ContextTimeout  time.Duration
    Env  *config.Env

}


func NewLoginUseCase(  userRepository domain.UserRepository,tokenManager utils.TokenManager ,timeout time.Duration, env *config.Env) domain.LoginUsecase {
	return &LoginUseCase{
		UserRepository: userRepository,
		TokenManager: tokenManager,
		ContextTimeout:  timeout,
        Env: env,
	}
}
// Login implements domain.LoginUsecase.
func (l *LoginUseCase) Login(c context.Context, req *domain.UserLogin) (*domain.LoginResponse, error) {
    ctx, cancel := context.WithTimeout(c, l.ContextTimeout)
    defer cancel()

    user, err := l.UserRepository.GetUserByEmail(ctx, req.Email)
    if err != nil {
        return nil, err
    }

    if !user.IsActivated{
        return nil,errors.New("user is not activated, Verify your email")
    }

    // Compare the hashed password with the plain text password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return nil, fmt.Errorf("password incorrect")
    }

    accessToken, err := l.TokenManager.CreateAccessToken(user, l.Env.AccessTokenSecret, l.Env.AccessTokenExpiryHour)
    if err != nil {
        return nil, err
    }
    refreshToken, err := l.TokenManager.CreateRefreshToken(user, l.Env.RefreshTokenSecret, l.Env.RefreshTokenExpiryHour)
    if err != nil {
        return nil, err
    }
    userID :=user.UserID.Hex() 
    _, err = l.UserRepository.UpdateToken(ctx, accessToken, refreshToken, userID)
    if err != nil {
        return nil, err
    }

    var loginResponse domain.LoginResponse
    loginResponse.AccessToken = accessToken
    loginResponse.RefreshToken = refreshToken
    loginResponse.UserID = userID

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
    _, err = l.UserRepository.UpdateToken(ctx, "", "", userID)
    if err != nil {
        return err
    }

    return nil
}




// LogOut implements domain.LogoutUsecase.
func (l *LoginUseCase) LogOut(c context.Context, userID string) error {
	ctx,cancel := context.WithTimeout(c, l.ContextTimeout)
	defer cancel()
	_,err := l.UserRepository.UpdateToken(ctx, "", "" ,userID)
	if err != nil {
		return err
	}
	return nil
}


