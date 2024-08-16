package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	FirstName  string             `json:"first_name" bson:"first_name" binding:"required,min=3,max=30"`
	LastName   string             `json:"last_name" bson:"last_name" binding:"max=30"`
	Email      string             `json:"email" bson:"email" binding:"required,email"`
	Bio        string             `json:"bio" bson:"bio"`
	ProfileImg string             `json:"profile_img" bson:"profile_img"`
	Password   string             `json:"password" bson:"password"`
	IsOwner    bool               `json:"is_owner" bson:"is_owner"`
	Role       string             `json:"role" bson:"role"` //may make only tobe admin or user
	Tokens     []string           `json:"tokens" bson:"tokens"`
}

// this structure defined to take data when user registers
type UserIn struct {
}
