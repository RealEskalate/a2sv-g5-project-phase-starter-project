package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string             `bson:"username" json:"username" validate:"required,min=3,max=30"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	Email      string             `bson:"email" json:"email" validate:"required,email"`
	Password   string             `bson:"password" json:"password"`
	Role       Role               `bson:"role" json:"role"`
	IsVerified bool               `bson:"is_verified" json:"is_verified"`
}

// Session struct represents a session in the system
type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	AccessToken  string             `bson:"access_token" json:"access_token"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
}

// CreateAccountRequest represents the payload for creating a new user account
type CreateAccountRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// CreateAccountResponse represents the response after creating a new user account
type CreateAccountResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// PasswordResetRequest struct represents a password reset request
type PasswordResetRequest struct {
	Email string `bson:"email" json:"email" validate:"required,email"`
}

// SetUpPasswordRequest represents the payload for setting up a new password
type SetUpPasswordRequest struct {
	Password string             `json:"password" validate:"required,min=8"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
}

// SetUpPasswordResponse represents the response after setting up a new password

// LoginRequest represents the payload for user login
type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

// LoginResponse represents the response after a successful login
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// LogoutRequest represents the payload for logging out a user
type LogoutRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

// LogoutResponse represents the response after a successful logout
type Response struct {
	Message string `json:"message"`
}

// ProfileUpdateRequest represents the payload for updating a user's profile
type ProfileUpdateRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `bson:"password" json:"password"`
}
