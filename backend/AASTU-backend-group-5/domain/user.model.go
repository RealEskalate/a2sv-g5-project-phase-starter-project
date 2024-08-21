package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// actual user model
type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName       string             `json:"username" bson:"username"`
	Bio            string             `json:"bio,omitempty" bson:"bio,omitempty"`
	ProfilePicture Media              `json:"profile_picture,omitempty" bson:"profile_picture,omitempty"`
	Email          string             `json:"email" bson:"email"`
	Is_Admin       bool               `json:"is_admin" bson:"is_admin"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	IsVerified     bool               `json:"is_verified" bson:"is_verified"`
	OAuthProvider  string             `json:"oauth_provider,omitempty" bson:"oauth_provider,omitempty"`
	OAuthID        string             `json:"oauth_id,omitempty" bson:"oauth_id,omitempty"`
}

// user model that will be returned from the server
type ResponseUser struct {
	ID             string `json:"_id" bson:"_id"`
	UserName       string `json:"username" bson:"username"`
	Bio            string `json:"bio,omitempty" bson:"bio,omitempty"`
	ProfilePicture Media  `json:"profile_picture,omitempty" bson:"profile_picture,omitempty"`
	Email          string `json:"email" bson:"email"`
	Is_Admin       bool   `json:"is_admin" bson:"is_admin"`
}

type UpdateUser struct {
	UserName       string `json:"username" bson:"username"`
	Bio            string `json:"bio,omitempty" bson:"bio,omitempty"`
	IsVerified     bool   `json:"is_verified" bson:"is_verified"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
}

type LogINUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type OAuthLoginUser struct {
	Provider string `json:"provider" bson:"provider"`
	Token    string `json:"token" bson:"token"`
}

type RegisterUser struct {
	UserName string `json:"username" bson:"username"`
	Bio      string `json:"bio,omitempty" bson:"bio,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type UpdatePassword struct {
	Password        string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirm_password" bson:"confirm_password"`
}
 
type VerifyEmail struct {
	Email string `json:"email" bson:"email"`
}

// from actual user model to response model to be done in usecase
func CreateResponseUser(user User) ResponseUser {
	return ResponseUser{
		ID:             user.ID.Hex(),
		UserName:       user.UserName,
		Bio:            user.Bio,
		ProfilePicture: user.ProfilePicture,
		Email:          user.Email,
		Is_Admin:       user.Is_Admin,
	}
}
