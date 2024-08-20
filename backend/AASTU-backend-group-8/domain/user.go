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
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Profile struct {
	ID        primitive.ObjectID `json:"id" gorm:"primaryKey"`
	UserID    primitive.ObjectID `json:"user_id"`
	Bio       string             `json:"bio"`
	AvatarURL string             `json:"avatar_url"`
}

type UserUsecaseInterface interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email *string) (*User, error)
	Register(user *User) error
	Login(user *AuthUser) (string, string, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(username , newPassword string) error
	DeleteUser(objectID primitive.ObjectID) error
}
	// RefreshToken(refreshToken *RefreshToken) (string, error)


type UserRepositoryInterface interface {
	//User operations
	Create(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(username string, user *User) error
	DeleteUser(id primitive.ObjectID) error
}

