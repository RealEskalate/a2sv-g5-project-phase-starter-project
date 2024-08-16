package domain

import "time"

type User struct {
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Bio        string    `json:"bio"`
	Avatar     string    `json:"avatar"`
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	Role       string    `json:"role" `
	Address    string    `json:"address"`
	JoinedDate time.Time `json:"joinedDate"`
}

type Token struct {
	Username  string `json:"username" bson:"username"`
	ExpiresAt int64  `json:"expires_at" bson:"expires_at"`
}

type UserRepository interface {
	CheckUsernameAndEmail(username, email string) error
	RegisterUser(user *User) error
	GetUserByUsernameorEmail(usernameoremail string) (*User, error)
	UpdateProfile(usernameoremail string, user *User) error
	Resetpassword(usernameoremail string, password string) error
	InsertToken(token *Token) error
	GetTokenByUsername(username string) (*Token, error)
	DeleteToken(username string) error
}

type UserUsecase interface {
	RegisterUser(user *User) error
	LoginUser(usernameoremail, password string) (string, string, error)
	UpdateProfile(usernameoremail string, user *User) error
	ResetPassword(usernameoremail, password string) error
	LogoutUser(username string) error
	ForgotPassword(email string) error
}
