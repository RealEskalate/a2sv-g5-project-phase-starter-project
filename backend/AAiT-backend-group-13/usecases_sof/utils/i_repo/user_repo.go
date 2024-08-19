package irepository

import (
	"github.com/google/uuid"
	usermodel "github.com/group13/blog/domain/models/user"
)

type UserRepository interface {
	Save(user *usermodel.User) error 
	FindById(id uuid.UUID) (*usermodel.User, error)
	FindByUsername(username string) (*usermodel.User, error)
	FindByEmail(email string) (*usermodel.User, error)
	
}