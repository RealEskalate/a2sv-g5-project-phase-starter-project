package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepository) DeleteRefreshTokenByUserID(ctx context.Context, userID primitive.ObjectID) error {
	filter := bson.M{"user_id": userID}
	_, err := ur.collection.DeleteOne(ctx, filter)
	return err
}
