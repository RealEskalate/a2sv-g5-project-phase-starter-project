package domain

import (
	"github.com/google/uuid"
)

const (
	CollectionLike = "likes"
)

type Like struct {
	ID     uuid.UUID `json:"id,omitempty" bson:"id,omitempty"`
	IsLike bool      `json:"is_like" bson:"is_like"`
	UserID uuid.UUID `json:"user_id" bson:"user_id"`
	BlogID uuid.UUID `json:"blog_id" bson:"blog_id"`
}
