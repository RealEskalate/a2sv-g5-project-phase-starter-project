package user_usecase

import (
	"context"

	"blog-api/domain/user"
)

func (u *UserUsecase) SignupUsecase(ctx context.Context, user *user.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.repo.SignupRepository(ctx, user)
}
