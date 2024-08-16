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
	dbUser, err := u.repo.FindUserByEmail(context.Background(), email)

	if err != nil {
		return nil, err
	}
	if dbUser != nil {
		return nil, errors.New("user already exists")
	}

	user.Role = "user"

	isEmpty, err := u.repo.IsEmptyCollection(context.Background())
	if err != nil {
		return nil, err
	}

	if isEmpty {
		user.Role = "owner"
	}

	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	err = u.repo.CreateUser(context.Background(), user)
	return user, err
}
