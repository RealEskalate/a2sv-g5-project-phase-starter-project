package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog represents the blog entity.
type Blog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	AuthorID  primitive.ObjectID `bson:"author_id" json:"author_id"`
	Tags      []string           `bson:"tags" json:"tags"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Views     int                `bson:"views" json:"views"`
	Likes     int                `bson:"likes" json:"likes"`
	Comments  int                `bson:"comments" json:"comments"`
}

// BlogCreationRequest is used when creating a new blog post.
type BlogCreationRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags"`
}

// BlogUpdateRequest is used when updating an existing blog post.
type BlogUpdateRequest struct {
	Title   string   `json:"title,omitempty"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

// BlogResponse represents the response returned for a blog post.
type BlogResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	AuthorID  primitive.ObjectID `json:"author_id"`
	Tags      []string           `json:"tags"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Views     int                `json:"views"`
	Likes     int                `json:"likes"`
	Comments  int                `json:"comments"`
}

// BlogFilters are the filters that can be applied when searching or filtering blog posts.
type BlogFilters struct {
	Tags       []string `json:"tags,omitempty"`
	Date       string   `json:"date,omitempty"`
	Popularity string   `json:"popularity,omitempty"`
}

// PopularityAction represents an action to track popularity metrics like views, likes, or comments.
type PopularityAction struct {
	Metric string `json:"metric"`
	Action string `json:"action"`
}
