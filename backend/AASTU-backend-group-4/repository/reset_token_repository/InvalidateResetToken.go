package reset_token_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *resetTokenRepository) InvalidateResetToken(ctx context.Context, token string) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"token": token}, bson.M{"$set": bson.M{"used": true}})
	return err
}
