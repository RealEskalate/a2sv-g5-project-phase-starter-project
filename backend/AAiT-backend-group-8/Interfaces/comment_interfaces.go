package Interfaces

import domain "AAiT-backend-group-8/Domain"

type ICommentUseCase interface {
	GetCommentByID(commentID string) (*domain.Comment, error)
	DecodeToken(tokenStr string) (string, string, error)
	DeleteCommentsOfBlog(blogID string) error
	UpdateComment(comment *domain.Comment, commentID string) (string, error)
	DeleteComment(commentID string) (string, error)
	GetComments(blogID string) ([]domain.Comment, error)
	CreateComment(comment *domain.Comment, blogID string) error
}
