package domain

import (
	"time"
)

type BlogPost struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" validate:"required,min=5,max=255"`
    Content     string    `json:"content" validate:"required"`
    AuthorID    uint      `json:"author_id"` // Foreign key to User model
    Tags        []string  `json:"tags" gorm:"type:text[]"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    //PublishedAt *time.Time `json:"published_at"`  Optional
}