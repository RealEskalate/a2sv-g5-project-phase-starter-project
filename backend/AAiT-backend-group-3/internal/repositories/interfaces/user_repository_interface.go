package repository_interface

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"AAIT-backend-group-3/internal/domain/models"
)
type UserRepository interface {
	SignUp(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
	UpdateProfile(ctx context.Context, id primitive.ObjectID, user *models.User) error
	Login(ctx context.Context, user *models.User) (*models.User, error)
	ForgetPassword(ctx context.Context, email string) error
    PromoteUser(ctx context.Context, userID primitive.ObjectID) error
    DemoteUser(ctx context.Context, userID primitive.ObjectID) error
}