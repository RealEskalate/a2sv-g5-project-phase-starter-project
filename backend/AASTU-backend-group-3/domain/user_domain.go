package domain

import (
	"encoding/json"
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
	Role      string              `bson:"role" json:"role"`
	CreatedAt primitive.Timestamp `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.Timestamp `bson:"updatedAt" json:"updatedAt"`
	Image     string              `bson:"image,omitempty" json:"image,omitempty"`

	ActivationToken string         `bson:"activation_token,omitempty" json:"activation_token,omitempty"`
	TokenCreatedAt  time.Time      `bson:"token_created_at"`
	IsActive        bool           `bson:"is_active"`
	RefreshTokens   []RefreshToken `bson:"refresh_tokens" json:"refresh_tokens"`

	GoogleID           string `bson:"google_id,omitempty" json:"google_id,omitempty"`
	PasswordResetToken string `bson:"password_reset_token,omitempty" json:"password_reset_token,omitempty"`
}

type ReturnUser struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string              `bson:"username" json:"username"`
	Email     string              `bson:"email" json:"email"`
	Bio       string              `bson:"bio,omitempty" json:"bio,omitempty"`
	Role      string              `bson:"role" json:"role"`
	CreatedAt primitive.Timestamp `bson:"createdAt" json:"createdAt"`
	Image     string              `bson:"image,omitempty" json:"image,omitempty"`
	IsActive  bool                `bson:"is_active"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Bio      string `bson:"bio,omitempty" json:"bio,omitempty"`
		Role     string `bson:"role" json:"role"`
		Image    string `bson:"image,omitempty" json:"image,omitempty"`
		IsActive bool   `bson:"is_active"`
	}{
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Username,
		Bio:      u.Bio,
		Role:     u.Role,
		IsActive: u.IsActive,
	})

}

type OAuthProvider string

const (
	Google OAuthProvider = "google"
)

type OAuthUserInfo struct {
	Provider   OAuthProvider
	ProviderID string
	Email      string
	FirstName  string
	LastName   string
	Name       string
	Picture    string
}

type RefreshToken struct {
	Token     string    `bson:"token" json:"token"`
	DeviceID  string    `bson:"device_id" json:"device_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type LogInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User		 ReturnUser `json:"user"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type TokenGenerator interface {
	GenerateToken(user User) (string, error)
	GenerateRefreshToken(user User) (string, error)
	RefreshToken(token string) (string, error)
}

type TokenVerifier interface {
	VerifyToken(token string) (*User, error)
	VerifyRefreshToken(token string) (*User, error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type UserUsecase interface {
	// for every user
	Login(user *User, deviceID string) (LogInResponse, *CustomError)
	RefreshToken(userID, deviceID, token string) (LogInResponse, *CustomError)
	Register(user User) *CustomError
	GetUserByUsernameOrEmail(username, email string) (User, *CustomError)
	AccountActivation(token string, email string) *CustomError
	Logout(userID, deviceID, token string) *CustomError
	LogoutAllDevices(userID string) *CustomError
	LogoutDevice(userID, deviceID string) *CustomError
	GetDevices(userID string) ([]string, *CustomError)

	ActivateAccountMe(Email string) *CustomError

	// for google oauth
	OAuthLogin(oauthUserInfo OAuthUserInfo, deviceID string) (LogInResponse, *CustomError)

	// reset password
	ResetPassword(token, newPassword string) *CustomError
	SendPasswordResetLink(email string) *CustomError

	// for user profile
	GetMyProfile(userID string) (User, *CustomError)
	GetUsers() ([]User, *CustomError)
	DeleteUser(userID string) (User, *CustomError)
	DeleteMyAccount(userID string) *CustomError
	UploadImage(userID string, imagePath string) *CustomError
	UpdateMyProfile(user User, UserID string) *CustomError

	UpdateUserRole(userID, role string) (User, *CustomError)
}

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
	FindOrCreateUserByGoogleID(oauthUserInfo OAuthUserInfo, deviceID string) (*User, error)
	GetUserByResetToken(token string) (User, error)
	GetUserByEmail(email string) (User, error)

	// ActivateAccountMe(Email string) error

	// // for user profile
	GetMyProfile(userID string) (User, error)
	GetUsers() ([]User, error)
	DeleteUser(userID string) (User, error)
	DeleteMyAccount(userID string) error
	// GetUser( userID string) (User, error)
	UploadImage(userID string, imagePath string) error
	UpdateMyProfile(user User, userID string) error

	// // GetUserBlogs(ctx context.Context, userID string) ([]Blog, error)

	// // Admin only
	UpdateUserRole(userID, role string) (User, error)
}
