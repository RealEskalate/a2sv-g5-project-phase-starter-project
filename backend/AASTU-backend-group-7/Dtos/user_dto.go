package Dtos

import (
	"time"
)

type RegisterUserDto struct {
	Email          string    `json:"email" validate:"required,email"`
	Password       string    `json:"password" validate:"required"`
	UserName       string    `json:"username" `
	Role           string    `json:"-",omitempty default:"user"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	EmailVerified  bool      `bson:"email_verified" default:"false"`
	Name           string    `json:name`
	CreatedAt      time.Time `json:"createdat"`
	UpdatedAt      time.Time `json:"updatedat"`
}

type LoginUserDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
