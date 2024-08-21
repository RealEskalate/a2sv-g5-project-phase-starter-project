package user_usecase

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *UserUsecase) DeleteRefreshTokenByUserID(ctx context.Context, userID string) error {
	ID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid id")
	}
	return u.repo.DeleteRefreshTokenByUserID(ctx, ID)
}
