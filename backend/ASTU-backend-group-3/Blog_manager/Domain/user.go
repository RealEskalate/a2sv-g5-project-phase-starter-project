package Domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id             primitive.ObjectID `json:"id" bson:"id"`
	Name           string             `json:"name" bson:"name"`
	Username       string             `json:"username" bson:"username"`
	Password       string             `json:"password" bson:"password"`
	Email          string             `json:"email" bson:"email"`
	PostsIDs       []string           `json:"posts_id" bson:"posts_id"`
	ProfilePicture string             `json:"profile_picture" bson:"profile_picture"`
	Bio            string             `json:"bio" bson:"bio"`
	Gender         string             `json:"gender" bson:"gender"`
	Role           string             `json:"role" bson:"role"`
	IsActive       bool               `json:"is_active" bson:"is_active"`
	Address        string             `json:"address" bson:"address"`
	IsOauth        bool               `json:"isoauth" bson:"isoauth"`
}

type RegisterInput struct {
	Name           string `json:"name" bson:"name"`
	Username       string `json:"username" bson:"username"`
	Password       string `json:"password" bson:"password"`
	Email          string `json:"email" bson:"email"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
	Bio            string `json:"bio" bson:"bio"`
	Gender         string `json:"gender" bson:"gender"`
	Address        string `json:"address" bson:"address"`
	IsOauth        bool   `json:"isoauth" bson:"isoauth"`
}

type LoginInput struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UpdateUserInput struct {
	Username       string `json:"username" bson:"username"`
	Password       string `json:"password" bson:"password"`
	ProfilePicture string `json:"profile_picture" bson:"profile_picture"`
	Bio            string `json:"bio" bson:"bio"`
	Address        string `json:"address" bson:"address"`
}

type ForgetPasswordInput struct {
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
}

type ResetPasswordInput struct {
	Username    string `json:"username" bson:"username"`
	NewPassword string `json:"password" bson:"password"`
}

type ChangePasswordInput struct {
	NewPassword string `json:"password" bson:"password"`
}
