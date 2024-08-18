package domain

import "time"

type User struct {
	ID             string    `bson:"_id" json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Password       string    `json:"password"`
	VerifyToken    string    `json:"-"`
	ExpirationDate time.Time `json:"expirationtoken"`
	IsAdmin        bool      `json:"is_admin"`
	IsActive       bool      `json:"is_active"`
}
type UserFilter struct {
	UserId    string
	Username  string
	Email     string
	FirstName string
	LastName  string

	IsAdmin bool
}
type UserFilterOption struct {
	Filter     UserFilter
	Pagination int
}
type UserRepository interface {
	Get(opts UserFilterOption) ([]User, error)
	Create(u *User) (User, error)
	Update(userId string, updateData User) (User, error)
	Delete(userId string) error
}
type UserUsecase interface {
	Get() ([]User, error)
	GetByID(userID string) (User, error)
	GetByEmail(email string) (User, error)
	GetByUsername(email string) (User, error)
	Create(u *User) (User, error)
	Update(userId string, updateData User) (User, error)
	Delete(userId string) error
}
