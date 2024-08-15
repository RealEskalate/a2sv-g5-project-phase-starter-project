package user

import (
	"blogApp/internal/domain"
	"blogApp/internal/repository"
	"blogApp/pkg/hash"
	"context"
	"errors"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) RegisterUser(user *domain.User) (*domain.User, error) {
	email := user.Email
	exists, err := u.repo.FindUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return nil, errors.New("user already exists")
	}
	user.Role = "user"

	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	err = u.repo.CreateUser(context.Background(), user)
	return user, err
}
