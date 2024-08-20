package interfaces

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	UpdateUser(user *dto.UserUpdate) error
	PromoteUser(id uuid.UUID, isPromote bool) error
	GetAllUsersWithName(name string) ([]uuid.UUID, error)
	UpdateUserToken(user *domain.User) error
}

