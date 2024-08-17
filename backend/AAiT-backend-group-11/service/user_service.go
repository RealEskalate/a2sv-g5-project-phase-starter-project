package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"time"
)

type UserService struct {
	userRepository interfaces.UserRepository
	contextTimeout time.Duration
}

// CreateUser implements interfaces.UserService.
func (us *UserService) CreateUser(user *entities.User) (*entities.User, error) {
	_, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	return us.userRepository.CreateUser(user)
}

// DeleteUser implements interfaces.UserService.
func (us *UserService) DeleteUser(userId string) error {
	_, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	return us.userRepository.DeleteUser(userId)
}

// DemoteUserToRegular implements interfaces.UserService.
func (us *UserService) DemoteUserToRegular(userId string) error {
	_, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	return us.userRepository.DemoteUserToRegular(userId)
}

// FindUserByEmail implements interfaces.UserService.
func (us *UserService) FindUserByEmail(email string) (*entities.User, error) {
	_, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()
	return us.userRepository.FindUserByEmail(email)
}

// FindUserById implements interfaces.UserService.
func (us *UserService) FindUserById(userId string) (*entities.User, error) {
	_, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()
	return us.userRepository.FindUserById(userId)
}

// PromoteUserToAdmin implements interfaces.UserService.
func (us *UserService) PromoteUserToAdmin(userId string) error {
	_, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	return us.userRepository.PromoteUserToAdmin(userId)
}

// UpdateUser implements interfaces.UserService.
func (us *UserService) UpdateUser(user *entities.User) (*entities.User, error) {
	_, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	return us.userRepository.UpdateUser(user)
}

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
