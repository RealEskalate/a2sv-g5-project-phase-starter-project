package usecases

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type IUserUseCase interface {
	CreateUser(user *domain.User) (*domain.User, error)
	// GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id uuid.UUID) error
	PromoteUser(id uuid.UUID, makeAdmin bool) error
	// GetUserByName(name string) (*domain.User, error)
}

type UserUseCase struct {
	userRepo interfaces.IUserRepository
}

func NewUserUseCase(repo interfaces.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (u *UserUseCase) CreateUser(user *domain.User) (*domain.User, error) {
	return u.userRepo.CreateUser(user)
}

// func (u *UserUseCase) GetUserByEmail(email string) (*domain.User, error) {
// 	return u.userRepo.GetUserByEmail(email)
// }

func (u *UserUseCase) GetUserByID(id uuid.UUID) (*domain.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *UserUseCase) UpdateUser(user *domain.User) error {
	return u.userRepo.UpdateUser(user)
}

func (u *UserUseCase) DeleteUser(id uuid.UUID) error {
	return u.userRepo.DeleteUser(id)
}

func (u *UserUseCase) PromoteUser(id uuid.UUID, makeAdmin bool) error {
	return u.userRepo.PromoteUser(id, makeAdmin)
}
