package interfaces

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(user *domain.User) *domain.CustomError
	GetUserByEmail(email string) (*domain.User, *domain.CustomError)
	GetUserByID(id uuid.UUID) (*domain.User, *domain.CustomError)
	GetUserByUsername(username string) (*domain.User, *domain.CustomError)
	UpdateUser(user *dto.UserUpdate) *domain.CustomError
	PromoteUser(id uuid.UUID, isPromote bool) *domain.CustomError
	GetAllUsersWithName(name string) ([]uuid.UUID, *domain.CustomError)
	UpdateUserToken(user *domain.User) *domain.CustomError
	Count() (int64, *domain.CustomError)
}

