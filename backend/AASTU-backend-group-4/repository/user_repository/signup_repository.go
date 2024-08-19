package user_repository

import (
	"blog-api/domain/user"
	"context"
)

func (ur *UserRepository) SignupRepository(ctx context.Context, user *user.User) error {
	_, err := ur.collection.InsertOne(ctx, user)
	return err
}
