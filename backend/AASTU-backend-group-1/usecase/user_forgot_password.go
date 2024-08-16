package usecase

import (
	"blogs/config"
	"fmt"
)

// ForgotPassword generates a password reset token and sends it to the user
func (u *UserUsecase) ForgotPassword(email string) error {
	// Check if the user exists
	user, err := u.UserRepo.GetUserByUsernameorEmail(email)
	if err != nil {
		return err
	}

	// Generate a password reset token
	resetToken, _, err := config.GenerateToken(user.Username, user.Role, "password-reset")
	if err != nil {
		return err
	}

	// In a real application, send an email to the user with the resetToken
	fmt.Printf("Password reset token generated: %s\n", resetToken)

	return nil
}
