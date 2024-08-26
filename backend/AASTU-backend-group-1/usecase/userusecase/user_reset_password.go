package userusecase

import (
	"blogs/config"
	"blogs/domain"
)

func (u *UserUsecase) ResetPassword(tokenString string) error {
	claims := &domain.RegisterClaims{}
	err := config.ValidateToken(tokenString, claims)
	if err != nil {
		return err
	}

	// Get the token from the repository
	_, err = u.UserRepo.GetTokenByUsername(claims.Username)
	if err != nil {
		return err
	}

	// Get user by username or email
	_, err = u.UserRepo.GetUserByUsernameorEmail(claims.Username)
	if err != nil {
		return err
	}

	// Update the password in the repository
	err = u.UserRepo.Resetpassword(claims.Username, claims.Password)
	if err != nil {
		return err
	}

	return nil
}
