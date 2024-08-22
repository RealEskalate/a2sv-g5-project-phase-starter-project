package blog_usecase

import (
	"blog-api/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (bu *BlogUsecase) CreateComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	comment.ID = primitive.NewObjectID()
	comment.CreatedAt = time.Now()

	err := bu.commentRepo.CreateComment(ctx, *comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
