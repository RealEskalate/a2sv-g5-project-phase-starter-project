package refresh_token_repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *refreshTokenRepository) StoreRefreshToken(ctx context.Context, userID string, tokenString string, expiresAt time.Time) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"refresh_token": tokenString,
			"expires_at":    expiresAt,
		},
	}
	opts := options.Update().SetUpsert(true) // Upsert will create a new document if one doesn't already exist

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}
