package user_usecase

import (
	"blog-api/infrastructure/auth"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid id")
	}
	return u.repo.UpdatePassword(ctx, ID, string(encryptedPassword))
}
