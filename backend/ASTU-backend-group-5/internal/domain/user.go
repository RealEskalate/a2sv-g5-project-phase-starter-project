package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProfile struct {
	ProfileUrl string `bson:"profile_url" json:"profile_url" form:"profile_pic"`
	FirstName  string `bson:"first_name" json:"first_name" form:"first_name"`
	LastName   string `bson:"last_name" json:"last_name" form:"last_name"`
	Gender     string `bson:"gender" json:"gender" form:"gender"`
	Bio        string `bson:"bio" json:"bio" form:"bio"`
	Profession string `bson:"profession" json:"profession" form:"profession"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB ObjectID
	UserName string             `bson:"username" json:"username" form:"username"`
	Email    string             `bson:"email" json:"email" form:"email"`
	Password string             `bson:"password" json:"password" form:"password"`
	Profile  UserProfile        `bson:"profile" json:"profile" form:"profile"`
	Role     string             `bson:"role" json:"role" form:"role"`
	Created  primitive.DateTime `bson:"created" json:"created" form:"created"`
	Updated  primitive.DateTime `bson:"updated" json:"updated" form:"updated"`
	Verified bool               `bson:"verified" json:"verified" form:"verified"`
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
}

type GetUserDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB ObjectID
	UserName string             `bson:"username" json:"username" form:"username"`
	Email    string             `bson:"email" json:"email" form:"email"`
	Profile  UserProfile        `bson:"profile" json:"profile" form:"profile"`
	Role     string             `bson:"role" json:"role" form:"role"`
	Created  primitive.DateTime `bson:"created" json:"created" form:"created"`
	Updated  primitive.DateTime `bson:"updated" json:"updated" form:"updated"`
	Verified bool               `bson:"verified" json:"verified" form:"verified"`
}

type CreateUserDTO struct {
	UserName string
	Email    string
	Password string
}

type UpdateUserDTO struct {
	UserName string      `bson:"username" json:"username" form:"username"`
	Profile  UserProfile `bson:"profile" json:"profile" form:"profile"`
	Email	string      `bson:"email" json:"email" form:"email"`
}


type UpdateUserPasswordDTO struct {
	Password string `bson:"password" json:"password" form:"password"`
}

type LoginUserDTO struct {
	Email    string
	Password string
}
