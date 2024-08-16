package auth

import "time"

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
	CreateUser(user User) (string, error)
	UpdateUser(id string, user User) (User, error)
	GetUserByUsername(username string) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUsers() ([]User, error)
	DeleteUser(id string) error
	RegisterToken(token string) error
	GetToken(token string) (Token, error)
	DeleteToken(token string) error
}

type AuthServices interface {
	Login(info LoginForm) (string, error)
	RegisterUser(user User)
	UpdateProfile(user User)
	Activate(userID string, token string)
	Logout(userID string)
	GenerateToken(user User) (string, error)
}
