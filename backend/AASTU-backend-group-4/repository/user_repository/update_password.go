package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepository) UpdatePassword(ctx context.Context, userID primitive.ObjectID, newPassword string) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"password":           newPassword,
			"reset_token":        nil,
			"reset_token_expiry": nil,
		},
	}
	_, err := ur.collection.UpdateOne(ctx, filter, update)
	return err
}
