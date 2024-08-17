package domain

import (
	"time"
	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id"`
	Author    Author    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}