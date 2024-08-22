package blogcontroller

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
)

type BlogDto struct {
	Title   string    `json:"title" binding:"required"`
	Content string    `json:"content" binding:"required"`
	Tags    []string  `json:"tags"`
	UserId  uuid.UUID `json:"auther_id" binding:"required"`
}

type BlogUpdateDto struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type BlogResponse struct {
	ID           uuid.UUID `json:"_id"`
	Title        string    `json:"title" binding:"required"`
	Content      string    `json:"content" binding:"required"`
	Tags         []string  `json:"tags"`
	UserId       uuid.UUID `json:"auther_id" binding:"required"`
	LikeCount    int       `json:"like_count"`
	DisLikeCount int       `json:"dislike_count"`
	CommentCount int       `json:"comment_count"`
}

func FromBlog(blog *models.Blog) *BlogResponse {
	return &BlogResponse{
		ID:           blog.ID(),
		Title:        blog.Title(),
		Content:      blog.Content(),
		Tags:         blog.Tags(),
		UserId:       blog.UserID(),
		LikeCount:    blog.LikeCount(),
		DisLikeCount: blog.DislikeCount(),
		CommentCount: blog.CommentCount(),
	}
}
