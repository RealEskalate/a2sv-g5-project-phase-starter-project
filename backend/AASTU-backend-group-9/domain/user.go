package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Role string

const (
	CollectionUser = "users"
	AdminRole Role = "ADMIN"
	UserRole Role = "USER"
)
type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	First_Name      string        `bson:"first_name" json:"first_name"`
	Last_Name       string        `bson:"last_name" json:"last_name" `
	Username        string 				`bson:"username" json:"username"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password" `
	Role            Role                `bson:"role" json:"role"`
	Bio             string        `bson:"bio" json:"bio"`
	Profile_Picture string        `bson:"profile_picture" json:"profile_picture"`
	Contact_Info    []ContactInfo `bson:"contact_info" json:"contact_info"`
}

type Privilage struct {
	Username  		string `json:"username"`
	Email           string `json:"email"`
	Role            Role `json:"role"`
}




