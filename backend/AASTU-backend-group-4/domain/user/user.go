package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	Role      string             `json:"role" bson:"role"` // e.g., "Admin" or "User"
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
	// RegisterUser(user User) error
	// LoginUser(ctx context.Context, request User) (LoginResponse, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	// LogOutUser(userID primitive.ObjectID) error
	// ForgetPassword(email string) error
	// UpdateUser(userID primitive.ObjectID, updatedUser *User) error
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
}

type UserDatabase interface {
	// GetByEmail(email string) (User, error)
	// GetByUsername(username string) (User, error)
}
