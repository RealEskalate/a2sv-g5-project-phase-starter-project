package user_usecase

import (
	"blog-api/infrastructure/auth"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) ResetPassword(ctx context.Context, resetToken, newPassword, resetTokenSecret string) error {
	claims, err := auth.VerifyResetToken(resetToken, resetTokenSecret)
	if err != nil {
		return err
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return errors.New("invalid or missing user_id in token claims")
	}

	return u.repo.UpdatePassword(ctx, userID, string(encryptedPassword))
}
