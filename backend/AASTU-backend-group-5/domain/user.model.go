package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// actual user model
type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName       string             `json:"username" bson:"username"`
	Bio            string             `json:"bio" bson:"bio"`
	ProfilePicture []byte             `json:"profile_picture" bson:"profile_picture"`
	Email          string             `json:"email" bson:"email"`
	Is_Admin       bool               `json:"is_admin" bson:"is_admin"`
	Password       string             `json:"password" bson:"password"`
	RefreshToken   string             `json:"refresh_token" bson:"refresh_token"`
}

// user model that will be returned from the server
type ResponseUser struct {
	ID             string `json:"_id" bson:"_id"`
	UserName       string `json:"username" bson:"username"`
	Bio            string `json:"bio" bson:"bio"`
	ProfilePicture []byte `json:"profile_picture" bson:"profile_picture"`
	Email          string `json:"email" bson:"email"`
	Is_Admin       bool   `json:"is_admin" bson:"is_admin"`
}

// from actual user model to response model to be don in usecase
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
