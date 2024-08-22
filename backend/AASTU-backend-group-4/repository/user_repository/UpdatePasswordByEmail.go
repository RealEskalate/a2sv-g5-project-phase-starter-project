package user_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *userRepository) UpdatePasswordByEmail(ctx context.Context, email, newPassword string) error {
	// Define the filter to find the user by email
	filter := bson.M{"email": email}

	// Define the update operation to set the new password
	update := bson.M{
		"$set": bson.M{
			"password": newPassword,
		},
	}

	// Perform the update operation
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Check if any document was modified
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
