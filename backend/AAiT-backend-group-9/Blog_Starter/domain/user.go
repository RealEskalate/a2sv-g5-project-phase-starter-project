package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	UserID         primitive.ObjectID `json:"user_id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	Email          string             `json:"email" bson:"email"`
	Password       string             `json:"password" bson:"password"`
	Name           string             `json:"name" bson:"name"`
	Bio            string             `json:"bio" bson:"bio"`
	Role           string             `json:"role" bson:"role"`
	ContactInfo    ContactInfo        `json:"contact_info" bson:"contact_info"` //it requires contact_info in the user profile update requirement given
	IsActivated    bool               `json:"is_verified" bson:"is_verified"`   //useful for email verification
	AccessToken    string             `json:"accessToken"`
	RefreshToken   string             `json:"refreshToken"`
	CreatedAt      time.Time          `json:"created_at" bson:"createtimestamp"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updatetimestamp"`
	ProfilePicture string             `json:"profile_picture" bson:"profile_picture"`
}

type ContactInfo struct {
	Phone   string `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`
}

type UserResponse struct {
	UserID         primitive.ObjectID `json:"user_id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	Email          string             `json:"email" bson:"email"`
	Name           string             `json:"name" bson:"name"`
	Bio            string             `json:"bio" bson:"bio"`
	ContactInfo    ContactInfo        `json:"contact_info" bson:"contact_info"`
	Role           string             `json:"role" bson:"role"`
	IsActivated    bool               `json:"is_verified" bson:"is_verified"` //useful for email verification
	ProfilePicture string             `json:"profile_picture" bson:"profile_picture"`
}

type UserUpdate struct {
	Username    string      `json:"username" bson:"username"`
	Name        string      `json:"name" bson:"name"`
	Bio         string      `json:"bio" bson:"bio"`
	ContactInfo ContactInfo `json:"contact_info" bson:"contact_info"`
}

type AuthenticatedUser struct {
	UserID string
	Email  string
	Role   string
}

type UserRepository interface {
	CreateUser(c context.Context, user *User) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByID(c context.Context, userID string) (*User, error)
	GetAllUser(c context.Context) ([]*User, error)
	UpdateProfile(c context.Context, user *UserUpdate, userID string) (*User, error)
	UpdateProfilePicture(c context.Context, profilePicPath string, userID string) (*User, error)
	UpdatePassword(c context.Context, password, userID string) (*User, error)
	UpdateRole(c context.Context, role, userID string) (*User, error)
	UpdateSignup(c context.Context, user *User) error
	UpdateToken(c context.Context, accessToken, refreshToken, userID string) (*User, error)
	DeleteUser(c context.Context, userID string) error
}

type UserUsecase interface {
	GetUserByEmail(c context.Context, email string) (*UserResponse, error)
	GetUserByID(c context.Context, userID string) (*UserResponse, error)

	GetAllUser(c context.Context) ([]*UserResponse, error)              //superAdmin privilage
	DeleteUser(c context.Context, userID string, password string) error //superAdmin,Admin,User privilage
	PromoteUser(c context.Context, userID string) error                 //superAdmin privilage
	DemoteUser(c context.Context, userID string) error                  //superAdmin privilage
	UpdateUser(c context.Context, user *UserUpdate, userID string) (*UserResponse, error)
	UpdateProfilePicture(c context.Context, profilePicPath string, userID string) (*UserResponse, error)
}
