package domain

import (
	"time"
	"github.com/google/uuid"
)

type Blog struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	AuthorID     uuid.UUID `json:"author_id"`
	Tags         []string  `json:"tags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ViewCount    int       `json:"view_count"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	Comments     []Comment `json:"comments"`
}
