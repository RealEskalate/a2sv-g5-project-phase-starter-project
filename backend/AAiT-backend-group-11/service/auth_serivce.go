package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"backend-starter-project/utils"
	"context"
	"time"
)

type authService struct {
	userService interfaces.UserService
	tokenRepo interfaces.RefreshTokenRepository
}

func NewAuthService(userService interfaces.UserService, tokenRepo interfaces.RefreshTokenRepository) interfaces.AuthenticationService {
	return &authService{
		userService: userService,
		tokenRepo: tokenRepo,
	}
}


func (service *authService) RegisterUser(user *entities.User) (*entities.User, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//check if the user already exists
	_, err := service.userService.FindUserByEmail(user.Email)
	if err == nil {
		return nil, err
	}

	//hash the password
	hashedPassword, err := utils.NewPasswordService().HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	//set the hashed password
	user.Password = hashedPassword

	//create the user
	createdUser, err := service.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
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
