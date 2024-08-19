package dtos

// dtos/user.go

import "go.mongodb.org/mongo-driver/bson/primitive"

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

// Response represents a generic response structure
type Response struct {
	Message string `json:"message"`
}

// ProfileUpdateRequest represents the payload for updating a user's profile
type ProfileUpdateRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `bson:"password" json:"password"`
}
