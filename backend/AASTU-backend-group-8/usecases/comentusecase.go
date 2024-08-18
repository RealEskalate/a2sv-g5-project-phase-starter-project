package usecases

import (
	"meleket/domain"
	"meleket/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentUsecase implements the CommentUsecaseInterface
type CommentUsecase struct {
	commentRepo repository.CommentRepositoryInterface
}

func NewCommentUsecase(cr repository.CommentRepositoryInterface) *CommentUsecase {
	return &CommentUsecase{
		commentRepo: cr,
	}
}

// AddComment adds a new comment to a blog post
func (uc *CommentUsecase) AddComment(comment *domain.Comment) error {
	return uc.commentRepo.AddComment(comment)
}

// GetCommentsByBlogID retrieves all comments for a specific blog post
func (uc *CommentUsecase) GetCommentsByBlogID(blogID primitive.ObjectID) ([]domain.Comment, error) {
	return uc.commentRepo.GetCommentsByBlogID(blogID)
}

// UpdateComment updates an existing comment
func (uc *CommentUsecase) UpdateComment(commentID primitive.ObjectID, content string) error {
	return uc.commentRepo.UpdateComment(commentID, content)
}

// DeleteComment deletes a comment by its ID
func (uc *CommentUsecase) DeleteComment(commentID primitive.ObjectID) error {
	return uc.commentRepo.DeleteComment(commentID)
}
