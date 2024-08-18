package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog struct represents a blog post in the system
type Blog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title" validate:"required"`
	Content   string             `bson:"content" json:"content" validate:"required"`
	AuthorID  primitive.ObjectID `bson:"author_id" json:"author_id"`
	Tags      []string           `bson:"tags" json:"tags"`
	Slug      string             `bson:"slug" json:"slug"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// Comment struct represents a comment on a blog post
type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BlogID    primitive.ObjectID `bson:"blog_id" json:"blog_id" validate:"required"`
	UserID    primitive.ObjectID `bson:"author_id" json:"author_id"`
	Content   string             `bson:"content" json:"content" validate:"required"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// Popularity struct represents popularity metrics for a blog post
type Popularity struct {
	BlogID       primitive.ObjectID `bson:"blog_id" json:"blog_id" validate:"required"`
	ViewCount    int                `bson:"view_count" json:"view_count"`
	LikeCount    int                `bson:"like_count" json:"like_count"`
	DislikeCount int                `bson:"dislike_count" json:"dislike_count"`
}

// BlogResponse struct represents the response when a single blog is retrieved, including comments and popularity
type BlogResponse struct {
	Blog       Blog       `json:"blog"`
	Comments   []Comment  `json:"comments"`
	Popularity Popularity `json:"popularity"`
}

// CreateBlogResponse represents the response after creating a new blog post
type CreateBlogResponse struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

// UpdateBlogRequest represents the payload for updating an existing blog post
type UpdateBlogRequest struct {
	BlogID   string             `json:"blog_id" validate:"required"`
	AuthorID primitive.ObjectID `bson:"author_id" json:"author_id"`
	Title    string             `json:"title,omitempty"`
	Content  string             `json:"content,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

// UpdateBlogResponse represents the response after updating a blog post
type UpdateBlogResponse struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

// DeleteBlogRequest represents the payload for deleting a blog post
type DeleteBlogRequest struct {
	BlogID   string             `json:"blog_id" validate:"required"`
	AuthorID primitive.ObjectID `bson:"author_id" json:"author_id"`
}

// TrackPopularityRequest represents the payload for tracking blog post popularity
type TrackPopularityRequest struct {
	BlogID primitive.ObjectID `json:"blog_id" validate:"required"`
	Action string             `json:"action" validate:"required"` // like, dislike, view, comment
}

// TrackPopularityResponse represents the response after tracking blog post popularity
type TrackPopularityResponse struct {
	Message string `json:"message"`
}

// FilterBlogRequest represents the payload for filtering blog posts
type FilterBlogRequest struct {
	Title        string   `json:"title,omitempty"`
	AuthorName   string   `json:"author_name,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Date         string   `json:"date,omitempty"`
	ViewCount    int      `bson:"view_count" json:"view_count"`
	LikeCount    int      `bson:"like_count" json:"like_count"`
	DislikeCount int      `bson:"dislike_count" json:"dislike_count"`
}

// FilterBlogResponse represents the response for filtered blog posts
type FilterBlogResponse struct {
	Blogs []Blog `json:"blogs"`
}

type CustomeError struct {
	Code    int
	Message string
}
