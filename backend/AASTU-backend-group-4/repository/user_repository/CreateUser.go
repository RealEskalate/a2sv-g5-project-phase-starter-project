package user_repository

import (
	"blog-api/domain"
	"context"
)

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
