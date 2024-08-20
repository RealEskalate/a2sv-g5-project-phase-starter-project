package usecases

import (
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type IUserUseCase interface {
	CreateUser(user *domain.User) (*domain.User, error)
	// GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
	UpdateUser(user *dto.UserUpdate) error
	PromoteUser(id uuid.UUID, isPromote bool) error
	// GetUserByName(name string) (*domain.User, error)
}

type UserUseCase struct {
	userRepo interfaces.IUserRepository
}

func NewUserUseCase(repo interfaces.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (u *UserUseCase) CreateUser(user *domain.User) (*domain.User, error) {
	return  user, u.userRepo.CreateUser(user)
}

// func (u *UserUseCase) GetUserByEmail(email string) (*domain.User, error) {
// 	return u.userRepo.GetUserByEmail(email)
// }

func (u *UserUseCase) GetUserByID(id uuid.UUID) (*domain.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *UserUseCase) UpdateUser(user *dto.UserUpdate) error {
	user.UpdatedAt = time.Now().UTC()
	return u.userRepo.UpdateUser(user)
}

func (u *UserUseCase) PromoteUser(id uuid.UUID, isPromote bool) error {
	return u.userRepo.PromoteUser(id, isPromote)
}
