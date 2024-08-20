package irepo

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
)

type UserRepository interface {
	Save(user *models.User) error
	FindById(id uuid.UUID) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

