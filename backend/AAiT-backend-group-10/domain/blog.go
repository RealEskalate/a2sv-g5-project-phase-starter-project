package domain

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID        uuid.UUID   `json:"id" bson:"_id"`
	Title     string      `json:"title" bson:"title" binding:"required"`
	Content   string      `json:"content" bson:"content" binding:"required"`
	Author    uuid.UUID   `json:"author" bson:"author" binding:"required"`
	Tags      []string    `json:"tags" bson:"tags" binding:"required"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`		
	UpdatedAt time.Time  `json:"updatedAt" bson:"updatedAt"`
	ViewCount int        `json:"viewCount" bson:"viewCount"`
}