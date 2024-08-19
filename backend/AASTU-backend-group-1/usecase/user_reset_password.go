package usecase

import (
	"blogs/config"
	"blogs/domain"
)

func (u *UserUsecase) ResetPassword(tokenString string) error {
	claims, err := config.ValidateToken(tokenString, "password-reset")
	if err != nil {
		return err
	}

	resetClaims, ok := claims.(*domain.PasswordResetClaims)
	if !ok {
		return err
	}

	// Get the token from the repository
	_, err = u.UserRepo.GetTokenByUsername(resetClaims.Username)
	if err != nil {
		return err
	}

	// Get user by username or email
	_, err = u.UserRepo.GetUserByUsernameorEmail(resetClaims.Username)
	if err != nil {
		return err
	}

	// Update the password in the repository
	err = u.UserRepo.Resetpassword(resetClaims.Username, resetClaims.Password)
	if err != nil {
		return err
	}

	return nil
}
