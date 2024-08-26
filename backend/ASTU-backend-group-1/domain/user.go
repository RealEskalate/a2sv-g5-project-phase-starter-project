package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// User represents a user in the system.
// @Description User model
// @Model User
// @Property ID string "ID of the user"
// @Property Username string "Username of the user"
// @Property Email string "Email of the user"
// @Property FirstName string "First name of the user"
// @Property LastName string "Last name of the user"
// @Property Password string "Password of the user"



type User struct {
	ID             string    `bson:"_id" json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Password       string    `json:"password"`
	VerifyToken    string    `json:"-"`
	RefreshToken   string    `json:"-"`
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
	IsAdmin   bool
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
	GetByUsername(username string) (User, error)
	Create(u *User) (User, error)
	UpdateUser(userId string, updateData User) (User, error)
	ChangePassword(email string, oldPassword string, newPassword string) (string, error)
	Delete(userId string) error
	AccountVerification(uemail string, confirmationToken string) error
	ResetPassword(email string, token string, password string) (string, error)
	ForgetPassword(email string) (string, error)
	LoginUser(uname string, password string,email string) (string, error)
	Logout(email string) error
	DemoteUser(userId string) (User, error)
	PromteUser(userId string) (User, error)
	PromteUserByEmail(email string) (User, error)
	DemoteUserByEmail(username string) (User, error)

}
type Claims struct {
	ID       string `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
	jwt.StandardClaims
}

type ChangePassword struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}