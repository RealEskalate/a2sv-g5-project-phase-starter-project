package usecase

import (
	"github.com/RealEskalate/blogpost/domain"
)

type CommentUseCase struct {
	CommentRepo domain.Comment_Usecase_interface
}

func (BC *CommentUseCase) GetComments(post_id string) ([]domain.Comment, error) {
	return BC.CommentRepo.GetComments(post_id)
}
func (BC *CommentUseCase) CreateComment(post_id string, user_id string) error {
	return BC.CommentRepo.CreateComment(post_id, user_id)
}
func (BC *CommentUseCase) DeleteComment(comment_id string) error {
	return BC.CommentRepo.DeleteComment(comment_id)
}
func (BC *CommentUseCase) UpdateComment(comment_id string) error {
	return BC.CommentRepo.UpdateComment(comment_id)
}
