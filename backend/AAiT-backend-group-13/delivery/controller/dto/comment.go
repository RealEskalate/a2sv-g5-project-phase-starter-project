package dto

import "github.com/google/uuid"

type CommentDto struct {
	Content string    `json:"content" binding:"required"`
	UserId  uuid.UUID `json:"userId" binding:"required"`
}
