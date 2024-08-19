package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (ur *UserRepository) UpdatePassword(ctx context.Context, userID string, newPassword string) error {
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"password":           newPassword,
			"reset_token":        nil,
			"reset_token_expiry": nil,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
