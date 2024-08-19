package dtos

import (
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogResponse struct represents the response when a single blog is retrieved, including comments and popularity
type BlogResponse struct {
	Blog       models.Blog       `json:"blog"`
	Comments   []models.Comment  `json:"comments"`
	Popularity models.Popularity `json:"popularity"`
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
	BlogID  string `json:"blog_id" validate:"required"`
	UserID  string `json:"user_id"`
	Action  string `json:"action" validate:"required"` // like, dislike, view, comment
	Comment string `json:"comment"`
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
	Blogs []models.Blog `json:"blogs"`
}

// CustomeError struct represents a custom error message
type CustomeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
