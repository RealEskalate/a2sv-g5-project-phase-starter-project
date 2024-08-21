package dtos

import (
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type BlogResponse struct {
	Blog       models.Blog       `json:"blog"`
	Comments   []models.Comment  `json:"comments"`
	Popularity models.Popularity `json:"popularity"`
}

type CreateBlogResponse struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type CreateBlogRequest struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags" `
}

type UpdateBlogRequest struct {
	Title   string   `json:"title,omitempty"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

type DeleteBlogRequest struct {
	BlogID   string `json:"blog_id" validate:"required"`
	AuthorID string `bson:"author_id" json:"author_id"`
}

type TrackPopularityRequest struct {
	BlogID  string `json:"blog_id" validate:"required"`
	UserID  string `json:"user_id"`
	Action  string `json:"action" validate:"required"` // like, dislike, view, comment
	Comment string `json:"comment"`
}

type FilterBlogRequest struct {
	Title        string   `json:"title,omitempty"`
	AuthorName   string   `json:"author_name,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Date         string   `json:"date,omitempty"`
	ViewCount    int      `bson:"view_count" json:"view_count"`
	LikeCount    int      `bson:"like_count" json:"like_count"`
	DislikeCount int      `bson:"dislike_count" json:"dislike_count"`
}

type CustomeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
