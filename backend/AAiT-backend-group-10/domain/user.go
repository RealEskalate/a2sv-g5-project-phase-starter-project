package domain

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id,omitempty" bson:"id,omitempty"`
	Email        string    `json:"email" bson:"email" validate:"required,email,unique"`
	Password     string    `json:"password" bson:"password" validate:"required,strong,min=8"`
	FullName     string    `json:"full_name" bson:"full_name" validate:"required"`
	Bio          string    `json:"bio,omitempty" bson:"bio,omitempty"`
	ImageURL     string    `json:"image_url,omitempty" bson:"image_url,omitempty"`
	IsAdmin      bool      `json:"is_admin" bson:"is_admin"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token"`
}
