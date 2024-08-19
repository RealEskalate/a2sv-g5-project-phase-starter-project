package irepo

import (
	"github.com/google/uuid"
	usermodel "github.com/group13/blog/domain/models/user"

	ihash "github.com/group13/blog/domain/i_hash"
)

type UserRepository interface {
	Save(user *usermodel.User) error
	FindById(id uuid.UUID) (*usermodel.User, error)
	FindByUsername(username string) (*usermodel.User, error)
	CheckUsernameAvailability(username string) error
	CheckEmailAvailability(email string) error
	MatchPassword(password string, hashedPassword string, hash ihash.Service) bool
	GenerateValidationLink(user usermodel.User) (string, error)
}
