package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string              `bson:"username" json:"username"`
	Email     string              `bson:"email" json:"email"`
	Password  string              `bson:"password" json:"password"`
	Bio       string              `bson:"bio,omitempty" json:"bio,omitempty"`
	Role      string                `bson:"role" json:"role"`
	CreatedAt primitive.Timestamp `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.Timestamp `bson:"updatedAt" json:"updatedAt"`

	ActivationToken string             `bson:"activation_token"`
	TokenCreatedAt time.Time          `bson:"token_created_at"`
	IsActive       bool               `bson:"is_active"`
	RefreshTokens   []RefreshToken      `bson:"refresh_tokens" json:"refresh_tokens"`
}


type RefreshToken struct {
    Token     string    `bson:"token" json:"token"`
    DeviceID  string    `bson:"device_id" json:"device_id"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type LogInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type UserUsecase interface {
	// for every user
	Login(user *User, deviceID string) (LogInResponse, error)
	RefreshToken(userID, deviceID, token string) (LogInResponse, error)
	Register(user User) error
	GetUserByUsernameOrEmail(username, email string) (User, error)
	AccountActivation(token string, email string) error
	// ForgotPassword(email string) error
	// ResetPassword(token, newPassword string) error
	// RefreshToken(token string) (string, error)
	// Logout(userID string) error

	// // for user profile
	// GetUsers(ctx context.Context) ([]User, error)
	// GetUser(ctx context.Context, userID string) (User, error)
	// GetUserBlogs(ctx context.Context, userID string) ([]Blog, error)
	// UpdateProfile(ctx context.Context, user User) error
	// UploadImage(ctx context.Context, userID string, imagePath string) error
	// DeleteMyAccount(ctx context.Context, userID string) error
	
	// // Admin only
	// DeleteUser(ctx context.Context, userID string) (User, error)
	// UpdateUserRole(ctx context.Context, userID, role string) (User, error)
}

// type AuthUsecase interface {
// 	Login(ctx context.Context, email, password string) (string, error)
// 	Register(ctx context.Context, user User) error
// 	ForgotPassword(ctx context.Context, email string) error
// 	ResetPassword(ctx context.Context, token, newPassword string) error
// 	RefreshToken(ctx context.Context, token string) (string, error)
// 	Logout(ctx context.Context, userID string) error
// }

type UserRepository interface {
	// for every user
	Login(user *User) (*User, error)
	Register(user User) error
	GetUserByUsernameOrEmail(username, email string) (User, error)
	AccountActivation(token string, email string) error
	UpdateUser(user *User) error
    DeleteRefreshToken(user *User, refreshToken string) error
    DeleteAllRefreshTokens(user *User) error
	GetUserByID(id string) (User, error)
	// ForgotPassword(email string) error
	// ResetPassword(token, newPassword string) error
	// RefreshToken(token string) (string, error)
	// Logout(userID string) error

	// // for user profile
	// GetUsers(ctx context.Context) ([]User, error)
	// GetUser(ctx context.Context, userID string) (User, error)
	// GetUserBlogs(ctx context.Context, userID string) ([]Blog, error)
	// UpdateProfile(ctx context.Context, user User) error
	// UploadImage(ctx context.Context, userID string, imagePath string) error
	// DeleteMyAccount(ctx context.Context, userID string) error

	// // Admin only
	// DeleteUser(ctx context.Context, userID string) (User, error)
	// UpdateUserRole(ctx context.Context, userID, role string) (User, error)

}

// type AuthRepository interface {
// 	Login(ctx context.Context, email, password string) (string, error)
// 	Register(ctx context.Context, user User) error
// 	ForgotPassword(ctx context.Context, email string) error
// 	ResetPassword(ctx context.Context, token, newPassword string) error
// 	RefreshToken(ctx context.Context, token string) (string, error)
// 	Logout(ctx context.Context, userID string) error
// }
