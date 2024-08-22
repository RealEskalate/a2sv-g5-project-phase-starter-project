package usecase

import (
	"AAiT-backend-group-6/domain"
	"context"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}


func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) CreateUser(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.CreateUser(ctx, user)
}
func (uu *userUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserByEmail(ctx, email)
}

func (uu *userUsecase) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserByUsername(ctx, username)
}

func (uu *userUsecase) GetUserByID(c context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserByID(ctx, id)
}

func (uu *userUsecase) DeleteUser(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.DeleteUser(ctx, id)
}

// GetUsers implements domain.UserUsecase.
func (u *userUsecase) GetUsers(c context.Context) ([]*domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserUsecase.
func (uu *userUsecase) UpdateUser(c context.Context, updatedUser *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.UpdateUser(ctx, updatedUser)
}
