package models

import (
	"time"
)

const (
	Like    = "like"
	Dislike = "dislike"

)

type Blog struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Title     string    `bson:"title" json:"title" validate:"required"`
	Content   string    `bson:"content" json:"content" validate:"required"`
	AuthorID  string    `bson:"author_id" json:"author_id"`
	Tags      []string  `bson:"tags" json:"tags"`
	Slug      string    `bson:"slug" json:"slug"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// Comment struct represents a comment on a blog post
type Comment struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	BlogID    string    `bson:"blog_id" json:"blog_id" validate:"required"`
	UserID    string    `bson:"author_id" json:"author_id"`
	Content   string    `bson:"content" json:"content" validate:"required"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// Popularity struct represents popularity metrics for a blog post
type Popularity struct {
	BlogID       string `bson:"blog_id" json:"blog_id" validate:"required"`
	ViewCount    int    `bson:"view_count" json:"view_count"`
	LikeCount    int    `bson:"like_count" json:"like_count"`
	DislikeCount int    `bson:"dislike_count" json:"dislike_count"`
}

type PopularityAction struct {
	BlogID string `bson:"blog_id" json:"blog_id" validate:"required"`
	Action string `bson:"action" json:"action" validate:"required"`
	UserID string `bson:"user_id" json:"user_id" validate:"required"`
}
