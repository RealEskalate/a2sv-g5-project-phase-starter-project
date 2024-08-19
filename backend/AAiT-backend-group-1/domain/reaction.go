package domain

import (
	"time"
	"github.com/google/uuid"
)

type Reaction struct {
	ID        uuid.UUID `json:"id"`
	Author    Author    `json:"author"`
	ReactionType string `json:"reaction_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}