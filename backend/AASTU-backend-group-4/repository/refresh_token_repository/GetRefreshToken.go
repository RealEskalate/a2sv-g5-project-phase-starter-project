package refresh_token_repository

import (
	"blog-api/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *refreshTokenRepository) GetRefreshToken(ctx context.Context, userID string) (string, error) {
	filter := bson.M{"user_id": userID}
	var result struct {
		RefreshToken string `bson:"refresh_token"`
	}

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil // No token found for this user
		}
		return "", err
	}

	return result.RefreshToken, nil
}
