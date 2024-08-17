package user_usecase

import (
	"context"
	"time"

	"blog-api/domain/user"
)

type userUsecase struct {
	repo           user.UserRepository
	contextTimeout time.Duration
	secret         string
}

func NewUserUsecase(userRepository user.UserRepository, timeout time.Duration, secret string) user.UserUsecase {
	return &userUsecase{
		repo:           userRepository,
		contextTimeout: timeout,
		secret:         secret,
	}
}

func (uc *userUsecase) GetByEmail(ctx context.Context, email string) (user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	u, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (uc *userUsecase) GetByUsername(ctx context.Context, username string) (user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	u, err := uc.repo.GetByUsername(ctx, username)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
