package user_usecase

import (
	"blog-api/domain"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) ResetPassword(ctx context.Context, req domain.ResetPasswordRequest) error {
	// Check if the new password and confirmation match
	if req.NewPassword != req.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}

	// Validate the reset token using AuthService
	email, err := u.authService.ValidateResetToken(ctx, req.Token)
	if err != nil {
		return err
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	// Update the user's password in the database
	err = u.userRepo.UpdatePasswordByEmail(ctx, email, string(hashedPassword))
	if err != nil {
		return errors.New("failed to update password")
	}

	// Invalidate the used reset token to prevent reuse
	err = u.authService.InvalidateResetToken(ctx, req.Token)
	if err != nil {
		return errors.New("failed to invalidate password reset token")
	}

	return nil
}
