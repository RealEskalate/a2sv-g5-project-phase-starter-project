package user_usecase

import (
	"context"
	"time"

	"blog-api/domain"
)

type UserUsecase struct {
	repo           domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) *UserUsecase {
	return &UserUsecase{
		repo:           userRepository,
		contextTimeout: timeout,
	}
}

func (uc *UserUsecase) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	u, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (uc *UserUsecase) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	u, err := uc.repo.GetByUsername(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}
