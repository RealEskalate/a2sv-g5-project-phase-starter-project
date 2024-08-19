package irepo

import (
	"github.com/google/uuid"
	usermodel "github.com/group13/blog/domain/models/user"
)

type UserRepository interface {
	Save(*usermodel.User) error
	ByUsername(string) (*usermodel.User, error)
	ById(uuid.UUID) (bool, error)
}
