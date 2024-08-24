package commentcontroller

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
)

type CommentDto struct {
	Content string `json:"content" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
}

type CommentResponse struct {
	ID      uuid.UUID `json:"_id"`
	Content string    `json:"content" binding:"required"`
	UserId  uuid.UUID `json:"userId" binding:"required"`
	BlogId  uuid.UUID `json:"blogId" binding:"required"`
}

func FromComment(comment *models.Comment) *CommentResponse {
	return &CommentResponse{
		ID:      comment.ID(),
		Content: comment.Content(),
		UserId:  comment.UserID(),
		BlogId:  comment.BlogID(),
	}
}
