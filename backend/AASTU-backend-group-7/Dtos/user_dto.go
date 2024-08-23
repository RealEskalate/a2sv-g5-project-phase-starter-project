package Dtos

import (
	"reflect"
	"time"

	"github.com/go-playground/validator"
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
	UserName string `json:"username",omitempty`
	Password string `json:"password" validate:"required"`
}

func CustomValidator(fl validator.FieldLevel) bool {
	// Get the struct field value
	v := fl.Field()

	// Check if both Email and UserName are not empty
	switch v.Type().Kind() {
	case reflect.Struct:
		email := v.FieldByName("Email").String()
		username := v.FieldByName("UserName").String()
		return email != "" || username != ""
	default:
		return false
	}
}
