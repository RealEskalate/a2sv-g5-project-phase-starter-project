package usecases

import (
	"time"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"aait.backend.g10/usecases/interfaces"
	"github.com/google/uuid"
)

type IUserUseCase interface {
	CreateUser(user *domain.User) (*domain.User, *domain.CustomError)
	// GetUserByEmail(email string) (*domain.User, *domain.CustomError)
	GetUserByID(id uuid.UUID) (*domain.User, *domain.CustomError)
	UpdateUser(user *dto.UserUpdate) *domain.CustomError
	PromoteUser(id uuid.UUID, isPromote bool) *domain.CustomError
	// GetUserByName(name string) (*domain.User, *domain.CustomError)
}

type UserUseCase struct {
	userRepo interfaces.IUserRepository
}

func NewUserUseCase(repo interfaces.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (u *UserUseCase) CreateUser(user *domain.User) (*domain.User, *domain.CustomError) {
	return  user, u.userRepo.CreateUser(user)
}

// func (u *UserUseCase) GetUserByEmail(email string) (*domain.User, *domain.CustomError) {
// 	return u.userRepo.GetUserByEmail(email)
// }

func (u *UserUseCase) GetUserByID(id uuid.UUID) (*domain.User, *domain.CustomError) {
	return u.userRepo.GetUserByID(id)
}

func (u *UserUseCase) UpdateUser(user *dto.UserUpdate) *domain.CustomError {
	user.UpdatedAt = time.Now().UTC()
	return u.userRepo.UpdateUser(user)
}

func (u *UserUseCase) PromoteUser(id uuid.UUID, isPromote bool) *domain.CustomError {
	return u.userRepo.PromoteUser(id, isPromote)
}
