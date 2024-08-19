package repository_interface

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"AAIT-backend-group-3/internal/domain/models"
)
type UserRepositoryInterface interface {
	SignUp(user *models.User) error
	GetUserByID(id primitive.ObjectID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(id primitive.ObjectID) error
	UpdateProfile(id primitive.ObjectID, user *models.User) error
	Login(user *models.User) (*models.User, error)
	ForgetPassword(email string) error
    PromoteUser(userID primitive.ObjectID) error
    DemoteUser(userID primitive.ObjectID) error
}