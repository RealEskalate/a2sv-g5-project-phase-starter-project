package domain

import (
	"github.com/google/uuid"
)

type Like struct {
	ID        uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`
	IsLike    *bool      `json:"is_like" bson:"is_like" binding:"required"`
	ReacterID uuid.UUID `json:"reacter_id" bson:"reacter_id"`
	BlogID    uuid.UUID `json:"blog_id" bson:"blog_id" binding:"required"`
}
