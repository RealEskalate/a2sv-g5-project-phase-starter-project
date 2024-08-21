package usecase

import (
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
	ctx, cancel:= context.WithTimeout(c, l.ContextTimeout)
	defer cancel()
	user,err:= l.UserRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	hashedPassword,err :=  bcrypt.GenerateFromPassword([]byte (req.Password),bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	if user.Password!= string(hashedPassword){
		return nil, fmt.Errorf("password incorrect")
	}
	accessToken,err:= l.TokenManager.CreateAccessToken(user,"secret", 1)
	if err!=nil{
		return nil, err
	}
	refreshToken,err := l.TokenManager.CreateRefreshToken(user, "secret", int(24))
	if err!=nil{
		return nil, err
	}
	
	_, err2:= l.UserRepository.UpdateToken(ctx,accessToken, refreshToken, user.UserID.String())
	if err2!=nil{
		return nil, err
	}

	var loginResponse domain.LoginResponse
	loginResponse.AccessToken = user.AccessToken
	loginResponse.RefreshToken = user.RefreshToken
	loginResponse.UserID = user.UserID.String()

	return &loginResponse, nil


}

// UpdatePassword implements domain.LoginUsecase.
func (l *LoginUseCase) UpdatePassword(c context.Context, req domain.ChangePasswordRequest, userID string) error {
	ctx, cancel := context.WithTimeout(c, l.ContextTimeout )
	defer cancel()
	_,err:=  l.UserRepository.UpdatePassword(ctx, req.Password, userID)
	if err!=nil{
		return err
	}
	return nil
}

