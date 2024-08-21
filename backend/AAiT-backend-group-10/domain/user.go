package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" bson:"_id,omitempty"`
	FullName     string    `json:"fullname" bson:"fullname" binding:"required"`
	Email        string    `json:"email" bson:"email" binding:"required,email"`
	Bio          string    `json:"bio" bson:"bio"`
	ImageURL     string    `json:"imageUrl" bson:"imageUrl"`
	IsAdmin      bool      `json:"isAdmin" bson:"isAdmin"`
	AccessToken  string    `json:"accessToken" bson:"accessToken"`
	RefreshToken string    `json:"refreshToken" bson:"refreshToken"`
	ResetToken   string    `json:"resetToken" bson:"resetToken"`
	ResetCode    int64     `json:"resetCode" bson:"resetCode"`
	GoogleSignIn bool      `json:"-" bson:"googlesignin"`
	Password     string    `json:"password" bson:"password" binding:"required"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
