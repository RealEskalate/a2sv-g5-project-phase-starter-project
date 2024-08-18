package user_repository

import (
	"blog-api/domain/user"
	"context"
)

func (ur *UserRepository) SignupRepository(ctx context.Context, user user.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(ctx, user)

	return err
}
