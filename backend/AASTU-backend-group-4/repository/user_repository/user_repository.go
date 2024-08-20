package user_repository

import (
	"blog-api/domain"
	"blog-api/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	collection mongo.Collection
}

func NewUserRepository(collection mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var u domain.User
	filter := bson.M{"email": email}
	err := ur.collection.FindOne(ctx, filter).Decode(&u)
	return u, err
}

func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var u domain.User

	filter := bson.M{"username": username}
	err := ur.collection.FindOne(ctx, filter).Decode(&u)
	return u, err
}
