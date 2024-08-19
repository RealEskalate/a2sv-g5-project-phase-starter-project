package user_repository

import (
	"blog-api/domain"
	"context"
)

func (ur *UserRepository) SignupRepository(ctx context.Context, user *domain.User) error {
	_, err := ur.collection.InsertOne(ctx, user)
	return err
}
