// Package models defines the domain models for the blogging application,
// including structures for managing blog posts, users, comments, and reactions.
package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
)

const (
	minTitleLength   = 3
	maxTitleLength   = 100
	minContentLength = 50
	maxContentLength = 10000
)

// Blog represents a blog post with metadata such as title, content, tags, and counts for likes, dislikes, and comments.
type Blog struct {
	id           uuid.UUID
	title        string
	content      string
	tags         []string
	createdDate  time.Time
	updatedDate  time.Time
	userID       uuid.UUID
	likeCount    int
	dislikeCount int
	commentCount int
}

// BlogConfig holds the parameters required for creating or updating a Blog.
type BlogConfig struct {
	Title   string
	Content string
	Tags    []string
	UserID  uuid.UUID
}

// MapBlogConfig holds the parameters for mapping a Blog from the database.
type MapBlogConfig struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Title        string
	Content      string
	Tags         []string
	CreatedDate  time.Time
	UpdatedDate  time.Time
	LikeCount    int
	DislikeCount int
	CommentCount int
}

// NewBlog creates a new Blog instance with the provided configuration.
// It validates the title and content according to predefined length constraints.
func NewBlog(config BlogConfig) (*Blog, error) {
	if err := validateBlogTitle(config.Title); err != nil {
		return nil, err
	}
	if err := validateBlogContent(config.Content); err != nil {
		return nil, err
	}
	now := time.Now()
	return &Blog{
		id:          uuid.New(),
		title:       config.Title,
		content:     config.Content,
		tags:        config.Tags,
		createdDate: now,
		updatedDate: now,
		userID:      config.UserID,
	}, nil
}

// MapBlog maps a Blog instance from the database using the provided configuration.
func MapBlog(config MapBlogConfig) *Blog {
	return &Blog{
		id:           config.ID,
		title:        config.Title,
		content:      config.Content,
		tags:         config.Tags,
		createdDate:  config.CreatedDate,
		updatedDate:  config.UpdatedDate,
		userID:       config.UserID,
		likeCount:    config.LikeCount,
		dislikeCount: config.DislikeCount,
		commentCount: config.CommentCount,
	}
}

// ID returns the unique identifier of the Blog.
func (b *Blog) ID() uuid.UUID { return b.id }

// UserID returns the ID of the user who created the Blog.
func (b *Blog) UserID() uuid.UUID { return b.userID }

// Title returns the title of the Blog.
func (b *Blog) Title() string { return b.title }

// Content returns the content of the Blog.
func (b *Blog) Content() string { return b.content }

// Tags returns the tags associated with the Blog.
func (b *Blog) Tags() []string { return b.tags }

// CreatedDate returns the date when the Blog was created.
func (b *Blog) CreatedDate() time.Time { return b.createdDate }

// UpdatedDate returns the date when the Blog was last updated.
func (b *Blog) UpdatedDate() time.Time { return b.updatedDate }

// LikeCount returns the number of likes on the Blog.
func (b *Blog) LikeCount() int { return b.likeCount }

// DislikeCount returns the number of dislikes on the Blog.
func (b *Blog) DislikeCount() int { return b.dislikeCount }

// CommentCount returns the number of comments on the Blog.
func (b *Blog) CommentCount() int { return b.commentCount }

// UpdateTitle updates the title of the Blog after validating the new title.
func (b *Blog) UpdateTitle(title string) error {
	if err := validateBlogTitle(title); err != nil {
		return err
	}
	b.title = title
	return nil
}

// UpdateContent updates the content of the Blog after validating the new content.
func (b *Blog) UpdateContent(content string) error {
	if err := validateBlogContent(content); err != nil {
		return err
	}
	b.content = content
	return nil
}

// UpdateTags updates the tags associated with the Blog.
func (b *Blog) UpdateTags(tags []string) error {
	b.tags = tags
	return nil
}

// UpdateCommentCount increments or decrements the comment count based on the increment parameter.
func (b *Blog) UpdateCommentCount(increment bool) error {
	return updateCount(&b.commentCount, increment)
}

// UpdateLikeCount increments or decrements the like count based on the increment parameter.
func (b *Blog) UpdateLikeCount(increment bool) error {
	fmt.Println(b.likeCount, increment)
	return updateCount(&b.likeCount, increment)
}

// UpdateDislikeCount increments or decrements the dislike count based on the increment parameter.
func (b *Blog) UpdateDislikeCount(increment bool) error {
	return updateCount(&b.dislikeCount, increment)
}

// validateBlogTitle validates the blog's title according to predefined length constraints.
func validateBlogTitle(title string) error {
	if len(title) < minTitleLength {
		return er.TitleTooShort
	}
	if len(title) > maxTitleLength {
		return er.TitleTooLong
	}
	return nil
}

// validateBlogContent validates the blog's content according to predefined length constraints.
func validateBlogContent(content string) error {
	if len(content) < minContentLength {
		return er.ContentTooShort
	}
	if len(content) > maxContentLength {
		return er.ContentTooLong
	}
	return nil
}

// updateCount increments or decrements the given count based on the increment parameter.
// Returns an error if decrementing would result in a negative count.
func updateCount(count *int, increment bool) error {
	if *count == 0 && !increment {
		return er.NewValidation("cannot decrement count below zero")
	}
	if increment {
		*count++
	} else {
		*count--
	}
	return nil
}
