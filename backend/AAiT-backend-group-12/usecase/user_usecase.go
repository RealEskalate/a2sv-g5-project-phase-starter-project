package usecase

import (
	"blog_api/domain"
	"context"
)

type UserUsecase struct {
	userRepository domain.UserRepositoryInterface
}

func NewUserUsecase(userRepository domain.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) Signup(c context.Context, user *domain.User) domain.CodedError {
	// TODO: Validate user input
	return u.userRepository.CreateUser(c, user)
}
