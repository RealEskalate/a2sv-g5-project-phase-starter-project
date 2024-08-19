package userusecase

import (
	"blogs/bootstrap"
	"blogs/config"
	"blogs/domain"
)

// ForgotPassword generates a password reset token and sends it to the user
func (u *UserUsecase) ForgotPassword(email string, newPassword string) error {
	// Check if the user exists
	user, err := u.UserRepo.GetUserByUsernameorEmail(email)
	if err != nil {
		return err
	}

	err = config.IsStrongPassword(newPassword)
	if err != nil {
		return err
	}

	newPassword, err = config.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Generate a password reset token
	resetToken, _, err := config.GenerateToken(
		&domain.PasswordResetClaims{
			Username: user.Username,
			Password: newPassword,
		},
		"password-reset")

	if err != nil {
		return err
	}

	// Get the API base URL
	apiBase, err := bootstrap.GetEnv("API_BASE")
	if err != nil {
		return err
	}

	// Prepare the email
	emailSubject := "Password Reset"
	emailBody := "Reset your password by clicking the link below:\n" +
		apiBase + "/users/reset-password?token=" + resetToken

	// Send the email
	err = config.SendEmail(email, emailSubject, emailBody)
	if err != nil {
		return err
	}

	return nil
}
