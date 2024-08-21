package dto

import (
	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

type CommentDto struct {
	ID        uuid.UUID `json:"id,omitempty"`
	BlogID    uuid.UUID `json:"blog_id"`
	UserID    uuid.UUID `json:"user_id"`
	Commenter string    `json:"commenter"`
	Comment   string    `json:"comment"`
}

func NewCommentDto(comment domain.Comment, commenter string) *CommentDto {
	return &CommentDto{
		ID:        comment.ID,
		BlogID:    comment.BlogID,
		UserID:    comment.CommenterID,
		Commenter: commenter,
		Comment:   comment.Comment,
	}
}
