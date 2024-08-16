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
	IsAdmin        bool               `json:"is_admin" bson:"is_admin"`
	IsActive       bool               `json:"is_active" bson:"is_active"`
	Address        string             `json:"address" bson:"address"`
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
	Email string `json:"email" bson:"email"`
}

type ResetPasswordInput struct {
	NewPassword string `json:"password" bson:"password"`
}
