package userusecase

import (
	"blogs/bootstrap"
	"blogs/config"
	"blogs/domain"
	"time"
)

func (u *UserUsecase) AddRoot() error {
	err := u.UserRepo.CheckRoot()

	if err != nil && err.Error() == "root user already exists" {
		return nil
	}

	if err != nil {
		return err
	}

	rootUsername, err := bootstrap.GetEnv("ROOT_USERNAME")
	if err != nil {
		return err
	}

	rootPassword, err := bootstrap.GetEnv("ROOT_PASSWORD")
	if err != nil {
		return err
	}

	err = config.IsStrongPassword(rootPassword)
	if err != nil {
		return err
	}

	rootPassword, err = config.HashPassword(rootPassword)
	if err != nil {
		return err
	}

	rootEmail, err := bootstrap.GetEnv("ROOT_EMAIL")
	if err != nil {
		return err
	}

	err = config.IsValidEmail(rootEmail)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username:   rootUsername,
		Password:   rootPassword,
		Email:      rootEmail,
		Role:       "root",
		JoinedDate: time.Now(),
		IsVerified: true,
	}

	return u.UserRepo.RegisterUser(user)
}
