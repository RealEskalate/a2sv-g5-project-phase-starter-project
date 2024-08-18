package auth

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"isactive"`
	IsAdmin   bool      `json:"isadmin"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type Token struct {
	ID          string `json:"id" bson:"_id"`
	TokenString string `json:"tokenstring"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user User) (string, error)
	UpdateUser(ctx context.Context, id string, user User) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	DeleteUser(ctx context.Context, id string) error
	RegisterToken(ctx context.Context, token string) error
	GetToken(ctx context.Context, token string) (Token, error)
	DeleteToken(ctx context.Context, token string) error
}

type AuthServices interface {
	Login(info LoginForm) (string, error)
	RegisterUser(user User)
	UpdateProfile(user User)
	Activate(userID string, token string)
	Logout(userID string)
	GenerateToken(user User) (string, error)
}
