package interfaces

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/dto"
	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
	UpdateUser(user *dto.UserUpdate) error
	PromoteUser(id uuid.UUID, makeAdmin bool) error
	GetAllUsersWithName(name string) ([]uuid.UUID, error)
}