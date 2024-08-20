package user_usecase

import (
	"blog-api/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *UserUsecase) UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *domain.UpdateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.repo.UpdateUser(ctx, userID, updatedUser)
	if err != nil {
		return err
	}

	return nil
}
