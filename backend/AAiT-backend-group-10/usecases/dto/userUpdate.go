package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserUpdate struct {
	ID           uuid.UUID `json:"id" bson:"_id,omitempty"`
	FullName     string    `json:"fullname" bson:"fullname,omitempty"`
	Bio          string    `json:"bio" bson:"bio,omitempty"`
	ImageURL     string    `json:"imageUrl" bson:"imageUrl,omitempty"`
	Password     string    `json:"password" bson:"password,omitempty"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}