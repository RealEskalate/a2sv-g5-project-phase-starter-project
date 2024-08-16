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

// DeleteUser implements UserUseCaseInterface.
func (u *UserUsecase) DeleteUser(id string) error {
	panic("unimplemented")
}

// FilterUsers implements UserUseCaseInterface.
func (u *UserUsecase) FilterUsers(filter map[string]interface{}) ([]*domain.User, error) {
	panic("unimplemented")
}

// GetAllUsers implements UserUseCaseInterface.
func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	panic("unimplemented")
}

// PromoteToAdmin implements UserUseCaseInterface.
func (u *UserUsecase) PromoteToAdmin(UserId string) error {
	panic("unimplemented")
}

// UpdateUser implements UserUseCaseInterface.
func (u *UserUsecase) UpdateUser(user *domain.User) error {
	panic("unimplemented")
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) RegisterUser(user *domain.User) (*domain.User, error) {
	email := user.Email
	exists, err := u.repo.FindUserByEmail(context.Background(), email)
	if exists != nil {
		return nil, errors.New("user already exists")
	}
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}
	user.Role = "user"

	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	err = u.repo.CreateUser(context.Background(), user)
	return user, err
}
