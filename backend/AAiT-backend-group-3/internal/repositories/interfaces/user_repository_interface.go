package repository_interface

import (
	"AAIT-backend-group-3/internal/domain/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type UserRepositoryInterface interface {
	SignUp(user *models.User) (*models.User, error)
	GetUserByID(id primitive.ObjectID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(id primitive.ObjectID) error
	UpdateProfile(id primitive.ObjectID, user *models.User) error
    PromoteUser(userID primitive.ObjectID) error
    DemoteUser(userID primitive.ObjectID) error
	UpdatePassword(userID string, hashedPassword string) error
}