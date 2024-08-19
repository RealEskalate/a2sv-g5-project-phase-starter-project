package user_usecase

import (
	"blog-api/infrastructure/auth"
	"context"
)

func (u *UserUsecase) GeneratePasswordResetToken(ctx context.Context, email, resetTokenSecret string, expiryHour int) error {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	// Generate the reset token (you might use JWT or another method)
	resetToken, err := auth.GenerateResetToken(user.ID.Hex(), resetTokenSecret, expiryHour)
	if err != nil {
		return err
	}

	// Store the reset token in the database
	err = u.repo.StoreResetToken(ctx, user.ID.Hex(), resetToken, expiryHour)
	if err != nil {
		return err
	}

	// Send the reset token via email (implement an email service to handle this)
	// err = u.emailService.SendPasswordResetEmail(user.Email, resetToken)
	// if err != nil {
	//     return err
	// }

	return nil
}
