package user_usecase

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *UserUsecase) PromoteDemote(ctx context.Context, userID primitive.ObjectID, action string) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.repo.PromoteDemote(ctx, userID, action)

	return err
}
