package usecase

import (
	"errors"
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"time"
)

func (au *UserUsecase) Login(user domain.User) (domain.User, error) {
	return au.UserRepo.Login(user)
}
func (au *UserUsecase) Register(user domain.User) error {
	

	if !infrastracture.IsValidEmail(user.Email) {
		return errors.New("invalid email format")
	}


	if !infrastracture.IsValidPassword(user.Password) {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one digit, one special character and minimum length of 8 characters")
	}

	
	_, err := au.UserRepo.GetUserByUsernameOrEmail(user.Username, user.Email)
	
	
	if err == nil {
		return errors.New("username or email already exists")
	}
	
		
	

	
	user.Role = "user"

	// Hash password
	hashedPassword, err := infrastracture.HashPassword(user.Password)
	token, err2 := infrastracture.GenerateActivationToken()


	if err != nil {
		return errors.New("could not hash password")	
	}
	if err2 != nil {
		return errors.New("could not generate activation token")
	}
	user.Password = hashedPassword
	user.ActivationToken = token
	user.TokenCreatedAt = time.Now()


	// Create user account in the database
	err = au.UserRepo.Register(user)
	if err != nil {
		return err
	}

	// Send activation email or link to the user
	err = infrastracture.SendActivationEmail(user.Email, token)
	if err != nil {
		return err
	}

	return nil
}


func (au *UserUsecase) GetUserByUsernameOrEmail(username, email string) (domain.User, error) {
	return au.UserRepo.GetUserByUsernameOrEmail(username, email)
}

func (au *UserUsecase) AccountActivation(token string, email string) error {
	return au.UserRepo.AccountActivation(token, email)
}