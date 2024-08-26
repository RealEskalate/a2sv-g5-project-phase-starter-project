package user_repository

import (
	"blog-api/domain"
	"blog-api/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *userRepository) GetByUsernameOrEmail(ctx context.Context, identifier string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": identifier},
			{"email": identifier},
		},
	}

	// Use FindOne to get the single result
	result := r.collection.FindOne(ctx, filter)

	// Handle the result of FindOne
	err := result.Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No document found, return nil, nil
			return nil, nil
		}
		// Return any other errors encountered
		return nil, err
	}

	// Successfully decoded user
	return &user, nil
}
