// Package models defines the domain models for the blogging application,
// including structures for managing blog posts, users, comments, and reactions.
package models

import (
	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
)

const (
	minCommentContentLength = 5   // Minimum allowed length for a comment's content
	maxCommentContentLength = 500 // Maximum allowed length for a comment's content
)

// Comment represents a comment on a blog post. It includes the content of the comment,
// the ID of the user who made the comment, and the ID of the blog post being commented on.
type Comment struct {
	id      uuid.UUID
	content string
	userID  uuid.UUID
	blogID  uuid.UUID
}

// CommentConfig holds the parameters required for creating or mapping a Comment.
type CommentConfig struct {
	Content string
	UserID  uuid.UUID
	BlogID  uuid.UUID
}

// NewComment creates a new Comment instance using the provided configuration.
// It validates the content according to predefined length constraints.
func NewComment(config CommentConfig) (*Comment, error) {
	if err := validateCommentContent(config.Content); err != nil {
		return nil, err
	}
	return &Comment{
		id:      uuid.New(),
		content: config.Content,
		userID:  config.UserID,
		blogID:  config.BlogID,
	}, nil
}

type MapCommentConfig struct {
	Id      uuid.UUID
	Content string
	UserID  uuid.UUID
	BlogID  uuid.UUID
}

func MapComment(config MapCommentConfig) *Comment {

	return &Comment{
		id:      config.Id,
		content: config.Content,
		userID:  config.UserID,
		blogID:  config.BlogID,
	}
}

// validateCommentContent checks if the length of the content is within the allowed limits.
// Returns an error if the content is too short or too long.
func validateCommentContent(content string) error {
	switch {
	case len(content) < minCommentContentLength:
		return er.ContentTooShort
	case len(content) > maxCommentContentLength:
		return er.ContentTooLong
	default:
		return nil
	}
}

// ID returns the unique identifier of the Comment.
func (c *Comment) ID() uuid.UUID { return c.id }

// Content returns the content of the Comment.
func (c *Comment) Content() string { return c.content }

// UserID returns the ID of the user who made the Comment.
func (c *Comment) UserID() uuid.UUID { return c.userID }

// BlogID returns the ID of the blog post being commented on.
func (c *Comment) BlogID() uuid.UUID { return c.blogID }
