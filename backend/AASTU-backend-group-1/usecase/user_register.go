package usecase

import (
	"blogs/config"
	"blogs/domain"
)

func (u *UserUsecase) RegisterUser(user *domain.User) error {
	// Validate the username
	err := config.IsValidUsername(user.Username)
	if err != nil {
		return err
	}

	// Validate the email
	err = config.IsValidEmail(user.Email)
	if err != nil {
		return err
	}

	// Check if the username and email are unique
	err = u.UserRepo.CheckUsernameAndEmail(user.Username, user.Email)
	if err != nil {
		return err
	}

	err = config.IsStrongPassword(user.Password)
	if err != nil {
		return err
	}

	// Hash the password before saving
	user.Password, err = config.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Save the new user in the repository
	err = u.UserRepo.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}
