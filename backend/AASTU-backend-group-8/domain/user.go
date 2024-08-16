package domain

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Email struct {
	Email string `json:"email"`
}

type Claims struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
	jwt.StandardClaims
}

type User struct {
	ID       primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name     string             `json:"name" validate:"required,min=2,max=100"`
	Email    string             `json:"email" validate:"required,email"`
	Password string             `json:"-"`
	Role     string             `json:"role"`
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Profile struct {
	ID        primitive.ObjectID `json:"id" gorm:"primaryKey"`
	UserID    uint               `json:"user_id"`
	Bio       string             `json:"bio"`
	AvatarURL string             `json:"avatar_url"`
}

type UserUsecaseInterface interface {
	Register(user *User) error
	Login(user *AuthUser) (string, string, error)
	DeleteRefeshToken(userID primitive.ObjectID) error
	ForgotPassword(email *string) error
	GetProfile(objectID primitive.ObjectID) (*User, error)
	UpdateProfile(objectID primitive.ObjectID, user *Profile) (*User, error)
	GetAllUsers() ([]User, error)
	DeleteUser(objectID primitive.ObjectID) error
	RefreshToken(refreshToken *RefreshToken) (string, error)
}

type UserRepositoryInterface interface {
	Create(user *User) error
	FindByUsername(username *string) (*User, error)
	// FindTokenByUsername(username *string) (*time.Time, error)
	FindByID(objectID primitive.ObjectID) (*User, error)
	FindAll() ([]User, error)
	UpdateProfile(userID string, user *User) (*User, error)
	Delete(objectID primitive.ObjectID) error
	SaveToken(username string, token *RefreshToken) error
	DeleteToken(username string) error
	FindRefreshToken(token string) (*RefreshToken, error)
}
