package interfaces

import  "backend-starter-project/domain/entities"

type CommentRepository interface {
	AddComment(comment *entities.Comment) (*entities.Comment, error)
	DeleteComment(commentId string) error
	GetCommentsByBlogPostId(blogPostId string) ([]entities.Comment, error)
	UpdateComment(comment *entities.Comment) (*entities.Comment, error)
}

type CommentService interface {
	AddComment(comment *entities.Comment) (*entities.Comment, error)
	DeleteComment(commentId string) error
	GetCommentsByBlogPostId(blogPostId string) ([]entities.Comment, error)
	UpdateComment(comment *entities.Comment) (*entities.Comment, error)
}