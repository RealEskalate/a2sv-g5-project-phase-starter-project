package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname          string             `json:"firstname" bson:"firstname"`
	Lastname           string             `json:"lastname" bson:"lastname"`
	Username           string             `json:"username" bson:"username"`
	Password           string             `json:"password" bson:"password"`
	Email              string             `json:"email" bson:"email"`
	Bio                string             `json:"bio" bson:"bio"`
	ProfilePicture     string             `json:"profile_picture" bson:"profile_picture"` // URL to the profile picture
	ContactInformation string             `json:"contact_information" bson:"contact_information"`
	IsAdmin            bool               `json:"isAdmin" bson:"isAdmin"`
	Active             bool               `json:"active" bson:"active"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
}

type UpdateRequest struct {
	Firstname          string `json:"firstname" bson:"firstname"`
	Lastname           string `json:"lastname" bson:"lastname"`
	Username           string `json:"username" bson:"username"`
	Bio                string `json:"bio" bson:"bio"`
	ProfilePicture     string `json:"profile_picture" bson:"profile_picture"`
	ContactInformation string `json:"contact_information" bson:"contact_information"`
	RefreshToken       string `json:"refresh_token" bson:"refresh_token"`
}

// type ForgotPasswordRequest struct {
// 	Email string `json:"email" binding:"required,email"`
// }

// type ResetPasswordRequest struct {
// 	ResetToken      string `json:"reset_token" binding:"required"`
// 	NewPassword     string `json:"new_password" binding:"required,min=8"`
// 	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
// }

type UserUsecase interface {
	// SignupUsecase(ctx context.Context, user *User) error
	// UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *User) error

	// GetByUsername(ctx context.Context, username string) (User, error)
	// DeleteRefreshTokenByUserID(ctx context.Context, userID string) error
	// GeneratePasswordResetToken(ctx context.Context, email, resetTokenSecret string, expiryHour int) error
	// ResetPassword(ctx context.Context, resetToken, newPassword, resetTokenSecret string) error
	// LoginUser(ctx context.Context, loginRequest LoginRequest, Env *bootstrap.Env) (LoginResponse, error)

	// New methods for user usecase
	UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *UpdateRequest) error
	SignUp(ctx context.Context, req SignupRequest) (SignupResponse, error)
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	Logout(ctx context.Context, userID string) error
	RequestPasswordReset(ctx context.Context, email, frontendBaseURL string) error
	ResetPassword(ctx context.Context, req ResetPasswordRequest) error
	PromoteDemote(ctx context.Context, userID primitive.ObjectID, action string) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
}

type UserRepository interface {
	// SignupRepository(ctx context.Context, user *User) error
	// ForgetPassword(email string) error
	// UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *UpdateRequest) error
	// GetByEmail(ctx context.Context, email string) (User, error)
	// GetByUsername(ctx context.Context, username string) (User, error)
	// DeleteRefreshTokenByUserID(ctx context.Context, userID string) error // used in logout
	// StoreResetToken(ctx context.Context, userID string, resetToken string, expiryHour int) error
	// UpdatePassword(ctx context.Context, userID string, newPassword string) error

	// New methods for user repository

	CreateUser(ctx context.Context, user *User) error
	GetByUsernameOrEmail(ctx context.Context, identifier string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	UpdatePasswordByEmail(ctx context.Context, email, newPassword string) error
	PromoteDemote(ctx context.Context, userID primitive.ObjectID, action string) error
	UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *UpdateRequest) error
}
