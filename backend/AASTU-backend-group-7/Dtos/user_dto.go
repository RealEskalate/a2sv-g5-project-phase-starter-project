package Dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterUserDto struct {
	Email          string               `json:"email" validate:"required,email"`
	Password       string               `json:"password" validate:"required"`
	UserName       string               `json:"username" `
	Role           string               `json:"-",omitempty default:"user"`
	ProfilePicture string               `json:"profile_picture"`
	Bio            string               `json:"bio"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
	Posts          []primitive.ObjectID `json:"posts"`
}

type LoginUserDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
