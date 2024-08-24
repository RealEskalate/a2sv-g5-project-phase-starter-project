package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/forms"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentUsecase struct {
	commentRepository entities.CommentRepository
	contextTimeout    time.Duration
}

func NewCommentUsecase(commentRepository entities.CommentRepository, timeout time.Duration) entities.CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
		contextTimeout:    timeout,
	}
}

func (cu *commentUsecase) GetComments(c context.Context, blogID string, limit int64, page int64) ([]entities.Comment, mongopagination.PaginationData, error) {
	return cu.commentRepository.GetComments(c, blogID, limit, page)
}
func (cu *commentUsecase) CreateComment(c context.Context, userID string, blogID string, comment *forms.CommentForm) (entities.Comment, error) {
	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return entities.Comment{}, err
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	newComment := entities.Comment{
		ID:        primitive.NewObjectID(),
		UserID:    userObjID,
		BlogID:    blogObjID,
		Content:   comment.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return cu.commentRepository.CreateComment(c, &newComment)
}
func (cu *commentUsecase) GetComment(c context.Context, commentID string) (entities.Comment, error) {
	return cu.commentRepository.GetComment(c, commentID)
}
func (cu *commentUsecase) UpdateComment(c context.Context, commentID string, updatedComment *entities.CommentUpdate) (entities.Comment, error) {
	updatedComment.UpdatedAt = time.Now()
	return cu.commentRepository.UpdateComment(c, commentID, updatedComment)
}
func (cu *commentUsecase) DeleteComment(c context.Context, commentID string) error {
	return cu.commentRepository.DeleteComment(c, commentID)
}
