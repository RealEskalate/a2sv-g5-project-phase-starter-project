package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	mongopagination "github.com/gobeam/mongo-go-pagination"
)

type commentUsecase struct {
	commentRepository domain.CommentRepository
	contextTimeout    time.Duration
}

func NewCommentUsecase(commentRepository domain.CommentRepository, timeout time.Duration) domain.CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
		contextTimeout:    timeout,
	}
}

func (cu *commentUsecase) GetComments(c context.Context, blogID string, limit int64, page int64) ([]domain.Comment, mongopagination.PaginationData, error) {
	return cu.commentRepository.GetComments(c, blogID, limit, page)
}
func (cu *commentUsecase) CreateComment(c context.Context, blogID string, comment *domain.Comment) (domain.Comment, error) {
	return cu.commentRepository.CreateComment(c, blogID, comment)
}
func (cu *commentUsecase) GetComment(c context.Context, blogID, commentID string) (domain.Comment, error) {
	return cu.commentRepository.GetComment(c, blogID, commentID)
}
func (cu *commentUsecase) UpdateComment(c context.Context, blogID, commentID string, updatedComment *domain.Comment) (domain.Comment, error) {
	return cu.commentRepository.UpdateComment(c, blogID, commentID, updatedComment)
}
func (cu *commentUsecase) DeleteComment(c context.Context, blogID, commentID string) error {
	return cu.commentRepository.DeleteComment(c, blogID, commentID)
}
