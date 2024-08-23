package domain

import (
	"context"
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	FirstName  string             `json:"first_name" bson:"first_name" binding:"required,min=3,max=30"`
	LastName   string             `json:"last_name" bson:"last_name" binding:"max=30"`
	Email      string             `json:"email" bson:"email" binding:"required,email"`
	Active     bool               `json:"is_active" bson:"is_active"`
	Bio        string             `json:"bio" bson:"bio"`
	ProfileImg string             `json:"profile_img" bson:"profile_img"`
	Password   string             `json:"password" bson:"password" binding:"required,min=4,max=30,StrongPassword"`
	IsOwner    bool               `json:"is_owner" bson:"is_owner"`
	Role       string             `json:"role" bson:"role"`
	Tokens     []string           `json:"tokens" bson:"tokens"`
	VerToken   string             `json:"verify_token" bson:"verfiy_token"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at"`
	LastLogin  primitive.DateTime `json:"last_login" bson:"last_login"`
}

// this structure defined for data sent as a response
type UserOut struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	FirstName  string             `json:"first_name" bson:"first_name"`
	LastName   string             `json:"last_name" bson:"last_name"`
	Email      string             `json:"email" bson:"email"`
	Bio        string             `json:"bio" bson:"bio"`
	ProfileImg string             `json:"profile_img" bson:"profile_img"`
	IsOwner    bool               `json:"is_owner" bson:"is_owner"`
	Role       string             `json:"role" bson:"role"` //may make only tobe admin or user
}

func (u *User) ToUserOut() *UserOut {
	return &UserOut{
		ID:         u.ID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
		Bio:        u.Bio,
		ProfileImg: u.ProfileImg,
		IsOwner:    u.IsOwner,
		Role:       u.Role,
	}
}

type UserUpdate struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Bio       string `json:"bio" bson:"bio"`
}

func (u *UserUpdate) ToUser() *User {
	return &User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Bio:       u.Bio,
	}
}

// user knows the password and wants to update
type UpdatePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// user forgot the password and wants to reset
// reset passowrd token will be exreacted from the url /reset-password/:user_id/:<reset password token>

type UserFilter struct {
	Email     string
	DateFrom  time.Time
	DateTo    time.Time
	Limit     int64
	Pages     int64
	FirstName string
	LastName  string
	Role      string
	IsOwner   string
	Active    string
	Bio       string
	sort      string
}

type UserUsecase interface {
	CreateUser(c context.Context, user *User) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserById(c context.Context, userId string) (*User, error)

	GetUsers(c context.Context, filter UserFilter) (*[]UserOut, mongopagination.PaginationData, error)
	UpdateUser(c context.Context, userID string, updatedUser *UserUpdate) (*User, error)
	ActivateUser(c context.Context, userID string) error
	DeleteUser(c context.Context, userID string) error
	IsUserActive(c context.Context, userID string) (bool, error)
	IsOwner(c context.Context) (bool, error)
	UpdateUserPassword(c context.Context, userID string, updatePassword *UpdatePassword) error
	UpdateProfilePicture(c context.Context, userID string, filename string) error
	PromoteUserToAdmin(c context.Context, userID string) error
	DemoteAdminToUser(c context.Context, userID string) error
}

type UserRepository interface {
	CreateUser(c context.Context, user *User) (*User, error)
	IsOwner(c context.Context) (bool, error)
	UpdateRefreshToken(c context.Context, userID string, refreshToken string) error
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserById(c context.Context, userId string) (*User, error)
	GetUsers(c context.Context, filter bson.M, userFilter UserFilter) (*[]User, mongopagination.PaginationData, error)
	UpdateUser(c context.Context, userID string, updatedUser *UserUpdate) (*User, error)
	ActivateUser(c context.Context, userID string) (*User, error)
	DeleteUser(c context.Context, userID string) error
	IsUserActive(c context.Context, userID string) (bool, error)
	RevokeRefreshToken(c context.Context, userID, refreshToken string) error
	UpdateProfilePicture(c context.Context, userID string, filename string) error

	ResetUserPassword(c context.Context, userID string, resetPassword *ResetPasswordRequest) error
	UpdateUserPassword(c context.Context, userID string, updatePassword *UpdatePassword) error

	PromoteUserToAdmin(c context.Context, userID string) error
	DemoteAdminToUser(c context.Context, userID string) error
}
