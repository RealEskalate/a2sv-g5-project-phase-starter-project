package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
)

type CommentUseCase struct {
	CommentRepo domain.Comment_Usecase_interface
}

func NewCommentUseCase(commentRepo domain.Comment_Usecase_interface) *CommentUseCase {
	return &CommentUseCase{
		CommentRepo: commentRepo,
	}
}

func (c *CommentUseCase) GetComments(post_id string) ([]domain.Comment, error) {
	return c.CommentRepo.GetComments(post_id)
}

func (c *CommentUseCase) CreateComment(post_id string, user_id string) error {
	return c.CommentRepo.CreateComment(post_id, user_id)
}

func (c *CommentUseCase) DeleteComment(comment_id string) error {
	return c.CommentRepo.DeleteComment(comment_id)
}

func (c *CommentUseCase) UpdateComment(comment_id string) error {
	return c.CommentRepo.UpdateComment(comment_id)
}
