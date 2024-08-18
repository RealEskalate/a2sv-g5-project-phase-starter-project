package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	IsAdmin   bool               `json:"isAdmin" bson:"isAdmin"`
	Active    bool               `json:"active" bson:"active"`
}

type LoginRequest struct {
	Email    string `form:"email" `
	Username string `form:"username" `
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserUsecase interface {
	SignupUsecase(ctx context.Context, user *User) error
	// UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
}

type UserRepository interface {
	SignupRepository(ctx context.Context, user *User) error
	// ForgetPassword(email string) error
	// UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
}
