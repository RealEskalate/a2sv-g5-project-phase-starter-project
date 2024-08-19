package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentUsecase struct {
	commentRepo repository_interface.CommentRepositoryInterface
}
func NewCommentUsecase(commentRepo repository_interface.CommentRepositoryInterface) *CommentUsecase {
	return &CommentUsecase{
		commentRepo: commentRepo,
	}
}
func (u *CommentUsecase) CreateComment(comment *models.Comment) error {
	return u.commentRepo.CreateComment(comment)
}

func (u *CommentUsecase) GetCommentByID(commentID primitive.ObjectID) (*models.Comment, error) {
	return u.commentRepo.GetCommentByID(commentID)
}

func (u *CommentUsecase) EditComment(commentID primitive.ObjectID, newComment *models.Comment) error {
	return u.commentRepo.EditComment(commentID, newComment)
}

func (u *CommentUsecase) DeleteComment(commentID primitive.ObjectID) error {
	return u.commentRepo.DeleteComment(commentID)
}

func (u *CommentUsecase) GetCommentsByIDList(commentIDs []primitive.ObjectID) ([]*models.Comment, error) {
	return u.commentRepo.GetCommentsByIDList(commentIDs)
}

func (u *CommentUsecase) GetCommentByAuthorID(authorID primitive.ObjectID) ([]*models.Comment, error) {
	return u.commentRepo.GetCommentByAuthorID(authorID)
}
