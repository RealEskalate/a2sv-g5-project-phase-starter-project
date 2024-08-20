package usecase

import (
	"blog_g2/domain"
	"context"

	"time"
)

type CommentUsecase struct {
	CommentRepo    domain.CommentRepository
	contextTimeout time.Duration
}

func NewCommentUsecase(Commentrepo domain.CommentRepository, timeout time.Duration) domain.CommentUsecase {
	return &CommentUsecase{
		CommentRepo:    Commentrepo,
		contextTimeout: timeout,
	}

}

func (cuse *CommentUsecase) CreateComment(c context.Context, blogID string, user_id string, comment domain.Comment) error {
	_, cancel := context.WithTimeout(c, cuse.contextTimeout)
	defer cancel()
	return cuse.CommentRepo.CreateComment(blogID, user_id, comment)
}

func (cuse *CommentUsecase) GetComments(c context.Context, blogID string) ([]domain.Comment, error) {
	_, cancel := context.WithTimeout(c, cuse.contextTimeout)
	defer cancel()
	return cuse.CommentRepo.GetComments(blogID)
}

func (cuse *CommentUsecase) UpdateComment(c context.Context, comment_id string, comment domain.Comment) error {
	_, cancel := context.WithTimeout(c, cuse.contextTimeout)
	defer cancel()
	return cuse.CommentRepo.UpdateComment(comment_id, comment)
}

func (cuse *CommentUsecase) DeleteComment(c context.Context, comment_id string) error {
	_, cancel := context.WithTimeout(c, cuse.contextTimeout)
	defer cancel()
	return cuse.CommentRepo.DeleteComment(comment_id)
}
