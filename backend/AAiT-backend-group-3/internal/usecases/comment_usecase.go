package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/repositories/interfaces"
)

type CommentUsecaseInterface interface {
	CreateComment(comment *models.Comment, blogID string) (string, error)
	GetCommentByID(commentID string) (*models.Comment, error)
	EditComment(commentID string, newComment *models.Comment) error
	DeleteComment(blogID string, commentID string) error
	GetCommentsByIDList(commentIDs []string) ([]*models.Comment, error)
	GetCommentByAuthorID(authorID string) ([]*models.Comment, error)
}

type CommentUsecase struct {
	commentRepo repository_interface.CommentRepositoryInterface
	blogRepo    repository_interface.BlogRepositoryInterface
}
func NewCommentUsecase(commentRepo repository_interface.CommentRepositoryInterface) CommentUsecaseInterface {
	return &CommentUsecase{
		commentRepo: commentRepo,
	}
}
func (u *CommentUsecase) CreateComment(comment *models.Comment, blogID string) (string, error ){
	err := u.blogRepo.AddCommentToTheList(blogID, comment.ID.Hex())
	if err != nil {
		return "", err
	}
	return u.commentRepo.CreateComment(comment, blogID)
}

func (u *CommentUsecase) GetCommentByID(commentID string) (*models.Comment, error) {
	return u.commentRepo.GetCommentByID(commentID)
}

func (u *CommentUsecase) EditComment(commentID string, newComment *models.Comment) error {
	return u.commentRepo.EditComment(commentID, newComment)
}

func (u *CommentUsecase) DeleteComment(blogID string, commentID string) error {
	err := u.blogRepo.DeleteCommentFromTheList(blogID, commentID)
	if err != nil {
		return err
	}
	return u.commentRepo.DeleteComment(commentID)
}

func (u *CommentUsecase) GetCommentsByIDList(commentIDs []string) ([]*models.Comment, error) {
	return u.commentRepo.GetCommentsByIDList(commentIDs)
}

func (u *CommentUsecase) GetCommentByAuthorID(authorID string) ([]*models.Comment, error) {
	return u.commentRepo.GetCommentByAuthorID(authorID)
}
