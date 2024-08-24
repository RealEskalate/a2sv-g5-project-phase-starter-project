package commentrepo

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
)

// CommentDto represents the data structure for a comment in the database.
type CommentDto struct {
	ID      uuid.UUID `bson:"_id"`
	Content string    `bson:"content"`
	BlogId  uuid.UUID `bson:"blogId"`
	UserId  uuid.UUID `bson:"userId"`
}

func FromComment(c *models.Comment) *CommentDto {
	return &CommentDto{
		ID:      c.ID(),
		Content: c.Content(),
		BlogId:  c.BlogID(),
		UserId:  c.UserID(),
	}
}
