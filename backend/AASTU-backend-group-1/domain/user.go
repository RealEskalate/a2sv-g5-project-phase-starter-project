package domain

import "time"

type User struct {
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Bio        string    `json:"bio"`
	Avatar     string    `json:"avatar"`
	UserName   string    `json:"username" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	Role       string    `json:"role" `
	Address    string    `json:"address"`
	JoinedDate time.Time `json:"joinedDate"`
}

type UserRepository interface {
	CheckUsernameAndEmail(username, email string) ( error)
	Register(user *User) error
	GetUserByUsernameorEmail(usernameoremail string) (*User, error)
	UpdateProfile(usernameoremail string, user *User) error
	Resetpassword(usernameoremail string, password string) error
}