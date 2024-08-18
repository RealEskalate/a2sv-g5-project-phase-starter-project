package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (ur *UserRepository) DeleteRefreshTokenByUserID(ctx context.Context, userID string) error {
	filter := bson.M{"user_id": userID}
	collection := ur.database.Collection(ur.collection)
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
