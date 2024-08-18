package comment

import (
	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
)

const (
	minContentLength = 5
	maxContentLength = 1000
)

type Comment struct {
	content string
	userId  uuid.UUID
	blogId  uuid.UUID
}

type Config struct {
	Content string
	UserId  uuid.UUID
	BlogId  uuid.UUID
}

func New(config Config) (*Comment, error) {
	if err := validateContent(config.Content); err != nil {
		return nil, err
	}

	return &Comment{
		content: config.Content,
		userId:  config.UserId,
		blogId:  config.BlogId,
	}, nil
}

func validateContent(content string) error {
	if len(content) < minContentLength {
		return er.ContentTooShort
	}
	if len(content) > maxContentLength {
		return er.ContentTooLong
	}

	return nil
}

func (c Comment) Content() string {
	return c.content
}

func (c Comment) UserId() uuid.UUID {
	return c.userId
}

func (c Comment) BlogId() uuid.UUID {
	return c.blogId
}
