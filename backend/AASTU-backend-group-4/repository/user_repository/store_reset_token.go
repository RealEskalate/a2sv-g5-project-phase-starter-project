package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepository) StoreResetToken(ctx context.Context, userID primitive.ObjectID, resetToken string, expiryHour int) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"reset_token":        resetToken,
			"reset_token_expiry": expiryHour, // Example expiry time
		},
	}

	_, err := ur.collection.UpdateOne(ctx, filter, update)
	return err
}
