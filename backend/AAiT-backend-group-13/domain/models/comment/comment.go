package comment

import (
	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
)

const (
	minContentLength = 5
	maxContentLength = 500
)

// Comment represents the Comment with private fields.
type Comment struct {
	content string
	userId  uuid.UUID
	blogId  uuid.UUID
}

// Config holds parameters for creating and Mapping a Comment.
type Config struct {
	Content string
	UserId  uuid.UUID
	BlogId  uuid.UUID
}

// New creates or maps a Comment with the provided configuration.
func New(config Config) (*Comment, error) {
	if err := validateContent(config.Content); err != nil {
		return nil, err
	}

	//returns Comment with specified fields
	return &Comment{
		content: config.Content,
		userId:  config.UserId,
		blogId:  config.BlogId,
	}, nil
}

// validateContent checks if length of content is allowed length if not returns error

func validateContent(content string) error {
	if len(content) < minContentLength {
		return er.ContentTooShort
	}
	if len(content) > maxContentLength {
		return er.ContentTooLong
	}

	return nil
}

// Content returns comments's content
func (c Comment) Content() string {
	return c.content
}

// UserId returns comments's owner Id
func (c Comment) UserId() uuid.UUID {
	return c.userId
}

// BlogId returns BlogId of the blog that the comment belongs to
func (c Comment) BlogId() uuid.UUID {
	return c.blogId
}
