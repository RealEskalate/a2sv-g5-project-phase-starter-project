package usecase

import (
	"blogs/config"
	"errors"
)

func (u *UserUsecase) LoginUser(usernameoremail string, password string) (string, string, error) {
	user, err := u.UserRepo.GetUserByUsernameorEmail(usernameoremail)
	if err != nil {
		return "", "", err
	}

	// Compare the hashed password
	err = config.ComparePassword(user.Password, password)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	// Generate access token
	accessToken, _, err := config.GenerateToken(user.Username, user.Role, "access")
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken, tokenEntry, err := config.GenerateToken(user.Username, user.Role, "refresh")
	if err != nil {
		return "", "", err
	}

	// Save the refresh token in the repository
	err = u.UserRepo.InsertToken(tokenEntry)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
