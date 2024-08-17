package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_Name    string             `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	Last_Name     string             `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Username      *string            `json:"user_name" bson:"user_name" validate:"min=5"`
	Email         *string            `json:"email" bson:"email" validate:"email"`
	Password      string             `json:"password" bson:"password" validate:"required,min=6"`
	Phone         *string            `json:"phone" bson:"phone"`
	Bio           *string            `json:"bio" bson:"bio"`
	ProfileImage  *string            `json:"profile_image" bson:"profile_image"`
	User_Role     string             `json:"user_role" bson:"user_role" validate:"omitempty,eq=Admin|eq=USER"`
	Access_Token  string             `json:"access_token" bson:"access_token"`
	Refresh_Token string             `json:"refresh_token" bson:"refresh_token"`
	Verified      bool               `json:"verified" bson:"verified"`
	Created_At    time.Time          `json:"created_at" bson:"created_at"`
	Updated_At    time.Time          `json:"updated_at" bson:"updated_at"`
}
