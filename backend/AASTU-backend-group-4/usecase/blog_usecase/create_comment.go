package blog_usecase

import (
	"blog-api/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) CreateComment(ctx context.Context, comment *domain.CommentRequest) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	var newComment domain.Comment

	newComment.ID = primitive.NewObjectID()
	newComment.BlogID = comment.BlogID
	newComment.UserID = comment.UserID
	newComment.Content = comment.Content
	newComment.CreatedAt = time.Now()

	err := bu.commentRepo.CreateComment(ctx, newComment)
	if err != nil {
		return nil, err
	}

	return &newComment, nil
}
