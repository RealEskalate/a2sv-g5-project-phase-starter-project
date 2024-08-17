package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
)

type authService struct {
	userService interfaces.UserService
	tokenRepo interfaces.RefreshTokenRepository
	passwordResetService interfaces.PasswordResetService
}

func NewAuthService(userService interfaces.UserService, tokenRepo interfaces.RefreshTokenRepository, passwordResetService interfaces.PasswordResetService) interfaces.AuthenticationService {
	return &authService{
		userService: userService,
		tokenRepo: tokenRepo,
		passwordResetService: passwordResetService,
	}
}


func (service *authService) RegisterUser(user *entities.User) (*entities.User, error) {

	//to be implemented
	return &entities.User{}, nil
}

func (service *authService) Login(emailOrUsername, password string) (*entities.RefreshToken,string, error) {
	//to be implemented
	return &entities.RefreshToken{}, "", nil
}

func (service *authService) Logout(userId string) error {

	//delete the token from database
	err := service.tokenRepo.DeleteRefreshTokenByUserId(userId)
	if err != nil {
		return err
	}
	return nil
	
}
