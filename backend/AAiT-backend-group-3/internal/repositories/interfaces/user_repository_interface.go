package repository_interface

import (
	"AAIT-backend-group-3/internal/domain/models"
	"time"
)

type UserRepositoryInterface interface {
	SignUp(user *models.User) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, user *models.User) error
	UpdateUserProfile(userID string, updateData *models.User) error
	PromoteUser(userID string) error
	DemoteUser(userID string) error
	UpdatePassword(userID string, hashedPassword string) error
	BlacklistToken(token string, remainingTime time.Duration) error
}
