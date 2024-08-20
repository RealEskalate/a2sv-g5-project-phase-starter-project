package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	UserName       string             `json:"username"`
	Email          string             `json:"email" validate:"required"`
	Password       string             `json:"password,omitempty" validate:"required"`
	Role           string             `json:"role"`
	ProfilePicture string             `json:"profile_picture"`
	Bio            string             `json:"bio"`
	EmailVerified  bool               `json:"email_verified"`
	Name           string             `json:name`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

// this could have been handled in a better way but i was too lazy to do it
type OmitedUser struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	UserName       string             `json:"username"`
	Email          string             `json:"email" validate:"required"`
	Password       string             `json:"-"`
	Role           string             `json:"role"`
	ProfilePicture string             `json:"profile_picture"`
	Bio            string             `json:"bio"`
	EmailVerified  bool               `json:"email_verified"`
	Name           string             `json:name`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
