package userusecase

import (
	"blogs/config"
	"blogs/domain"
	"log"
)

func (u *UserUsecase) LoginUser(usernameoremail string, password string) (string, string, error) {
	user, err := u.UserRepo.GetUserByUsernameorEmail(usernameoremail)
	if err != nil {
		log.Println(err, "email or username not found")
		return "", "", config.ErrIncorrectPassword
	}

	// Compare the hashed password
	err = config.ComparePassword(user.Password, password)
	if err != nil {
		log.Println(err, "password incorrect")
		return "", "", config.ErrIncorrectPassword
	}

	// Generate access token
	accessToken, err := config.GenerateToken(
		&domain.LoginClaims{
			Username: user.Username,
			Role:     user.Role,
			Type:     "access",
		})

	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshClaims := &domain.LoginClaims{
		Username: user.Username,
		Role:     user.Role,
		Type:     "refresh",
	}

	refreshToken, err := config.GenerateToken(refreshClaims)
	if err != nil {
		return "", "", err
	}

	// Save the refresh token in the repository
	err = u.UserRepo.InsertToken(refreshClaims.ToToken())
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
