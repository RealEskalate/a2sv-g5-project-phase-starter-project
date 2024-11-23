package user_usecase

import (
	"context"
	"fmt"
)

func (u *userUsecase) RequestPasswordReset(ctx context.Context,  email, frontendBaseURL string) error {
	// Generate a password reset token
	token, err := u.authService.GeneratePasswordResetToken(ctx, email)
	if err != nil {
		return err
	}
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", frontendBaseURL, token)

	// Send the password reset email to the user
	return u.emailService.SendPasswordResetEmail(ctx, email, resetLink)
}
