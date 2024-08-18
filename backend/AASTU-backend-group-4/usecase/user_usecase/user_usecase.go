package user_usecase

import (
	"time"

	"blog-api/domain/user"
)

type UserUsecase struct {
	repo           user.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository user.UserRepository, timeout time.Duration) user.UserUsecase {
	return &UserUsecase{
		repo:           userRepository,
		contextTimeout: timeout,
	}
}
