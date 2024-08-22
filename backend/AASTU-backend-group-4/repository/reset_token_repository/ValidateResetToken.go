package reset_token_repository

import (
	"blog-api/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *resetTokenRepository) ValidateResetToken(ctx context.Context, token string) (string, error) {
	var result domain.PasswordResetToken
	filter := bson.M{"token": token, "used": false, "expiry": bson.M{"$gte": time.Now()}}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.Email, nil
}
