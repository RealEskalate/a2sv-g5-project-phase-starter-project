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
	Role      string             `json:"role" bson:"role"` // e.g., "Admin" or "User"
	Active    bool               `json:"active" bson:"active"`
}

type SignupRequest struct {
	Firstname string `json:"firstname" bson:"firstname" binding:"required"`
	Lastname  string `json:"lastname" bson:"lastname" binding:"required"`
	Username  string `json:"username" bson:"username" binding:"required"`
	Password  string `json:"password" bson:"password" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
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
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	SignupUsecase(ctx context.Context, user *User) error
	// LogOutUser(userID primitive.ObjectID) error
	// ForgetPassword(email string) error
	// UpdateUser(userID primitive.ObjectID, updatedUser *User) error
}

type UserRepository interface {
  GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
}