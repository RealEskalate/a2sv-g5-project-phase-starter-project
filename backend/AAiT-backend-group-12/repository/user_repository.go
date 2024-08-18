package repository

import (
	"blog_api/domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection: collection}
}

func (r *UserRepository) CreateUser(c context.Context, user *domain.User) domain.CodedError {
	_, err := r.collection.InsertOne(c, user)
	// TODO: Handle index error types

	if err != nil {
		return *domain.NewError("error: failed to create user, "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}
