package usecase

import (
	"blogs/bootstrap"
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

	// Validate the password
	err = config.IsStrongPassword(user.Password)
	if err != nil {
		return err
	}

	// Hash the password before saving
	user.Password, err = config.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Setup verification email
	apiBase, err := bootstrap.GetEnv("API_BASE")
	if err != nil {
		return err
	}

	verifyToken, _, err := config.GenerateToken(
		&domain.RegisterClaims{
			Username: user.Username,
		}, "register")

	if err != nil {
		return err
	}

	emailHeader := "Welcome to Blogs!"
	emailBody := "Please verify your email by clicking the link below.\n" + apiBase + "/users/verify?token=" + verifyToken

	err = config.SendEmail(user.Email, emailHeader, emailBody)
	if err != nil {
		return err
	}

	err = u.UserRepo.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}
