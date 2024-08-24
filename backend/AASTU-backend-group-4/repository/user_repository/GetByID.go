package user_repository

import (
	"blog-api/domain"
	"blog-api/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	err = r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
