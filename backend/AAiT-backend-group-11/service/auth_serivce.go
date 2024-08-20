package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"errors"
)

type authService struct {
	userService          interfaces.UserService
	tokenRepo            interfaces.RefreshTokenRepository
	passwordResetService interfaces.PasswordResetService
	passwordService      interfaces.PasswordService
	tokenService         interfaces.TokenService
}

func NewAuthService(userService interfaces.UserService, tokenRepo interfaces.RefreshTokenRepository, passwordResetService interfaces.PasswordResetService,
	passService interfaces.PasswordService, tokenService interfaces.TokenService) interfaces.AuthenticationService {
	return &authService{
		userService:          userService,
		tokenRepo:            tokenRepo,
		tokenService:         tokenService,
		passwordResetService: passwordResetService,
		passwordService:      passService,
	}
}

func (service *authService) RegisterUser(user *entities.User) (*entities.User, error) {

	//to be implemented
	return &entities.User{}, nil
}

func (service *authService) Login(emailOrUsername, password string) (*entities.RefreshToken, string, error) {
	user, _ := service.userService.FindUserByEmail(emailOrUsername)
	err := service.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return nil, "", errors.New("Invalid password")
	}
	token, err := service.tokenService.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}
	refresh_tok, err := service.tokenService.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", err
	}
	return refresh_tok, token, nil
}

func (service *authService) Logout(userId string) error {

	//delete the token from database
	err := service.tokenRepo.DeleteRefreshTokenByUserId(userId)
	if err != nil {
		return err
	}
	return nil

}
