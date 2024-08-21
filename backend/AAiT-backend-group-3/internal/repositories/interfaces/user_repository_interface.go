package repository_interface

import (
	"AAIT-backend-group-3/internal/domain/models"
)
type UserRepositoryInterface interface {
	SignUp(user *models.User) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(id string) error
	UpdateProfile(id string, user *models.User) error
    PromoteUser(userID string) error
    DemoteUser(userID string) error
	UpdatePassword(userID, hashedPassword string) error
}