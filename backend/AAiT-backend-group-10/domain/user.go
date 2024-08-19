package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID    `json:"id" bson:"_id,omitempty"`
	FullName     string    `json:"fullname" bson:"fullname" binding:"required"`
	Email        string    `json:"email" bson:"email" binding:"required,email"`
	Bio          string    `json:"bio" bson:"bio"`
	ImageURL     string    `json:"imageUrl" bson:"imageUrl"`
	IsAdmin      bool      `json:"isAdmin" bson:"isAdmin"`
	RefreshToken string    `json:"refreshToken" bson:"refreshToken"`
	Password     string    `json:"password" bson:"password" binding:"required"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}