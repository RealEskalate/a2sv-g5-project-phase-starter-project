package user_usecase

import (
	"time"

	"blog-api/domain/user"
)

type userUsecase struct {
	repo           user.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository user.UserRepository, timeout time.Duration) user.UserUsecase {
	return &userUsecase{
		repo:           userRepository,
		contextTimeout: timeout,
	}
}
