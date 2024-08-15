package domain

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProfile struct {
	ProfileUrl string `bson:"profile_url" json:"profile_url"`
	FirstName  string `bson:"first_name" json:"first_name"`
	LastName   string `bson:"last_name" json:"last_name"`
	Gender     string `bson:"gender" json:"gender"`
	Bio        string `bson:"bio" json:"bio"`
	Profession string `bson:"profession" json:"profession"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` 
	UserName string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Profile  UserProfile        `bson:"profile" json:"profile"`
	Role	 string 			`bson:"role" json:"role"`
}

