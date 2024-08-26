package userusecase

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

	verifyToken, err := config.GenerateToken(
		&domain.RegisterClaims{
			User: *user,
		})

	if err != nil {
		return err
	}

	emailHeader := "Welcome to Blogs!"
	emailBody := "Hello " + user.Username + ", please verify your email by clicking <a href=\"" + apiBase + "/users/verify?token=" + verifyToken + "\">here</a>."

	err = config.SendEmail(user.Email, emailHeader, emailBody, true)
	if err != nil {
		return err
	}

	return nil
}
