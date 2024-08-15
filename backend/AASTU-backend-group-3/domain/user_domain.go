package domain

import (
	"context"

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
	Role      Role                `bson:"role" json:"role"`
	CreatedAt primitive.Timestamp `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.Timestamp `bson:"updatedAt" json:"updatedAt"`
}

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, userID string) (User, error)
	GetUserBlogs(ctx context.Context, userID string) ([]Blog, error)
	UpdateProfile(ctx context.Context, user User) error
	UploadImage(ctx context.Context, userID string, imagePath string) error
}

type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, user User) error
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
	RefreshToken(ctx context.Context, token string) (string, error)
	Logout(ctx context.Context, userID string) error
}

type UserRepository interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, userID string) (User, error)
	GetUserBlogs(ctx context.Context, userID string) ([]Blog, error)
	UpdateProfile(ctx context.Context, user User) error
	UploadImage(ctx context.Context, userID string, imagePath string) error
}

type AuthRepository interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, user User) error
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
	RefreshToken(ctx context.Context, token string) (string, error)
	Logout(ctx context.Context, userID string) error
}
