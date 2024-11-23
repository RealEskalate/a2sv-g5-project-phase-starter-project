package comment_repository

import (
	"blog-api/domain"
	"context"
)

func (cr *CommentRepository) CreateComment(ctx context.Context, comment domain.Comment) error {
	_, err := cr.collection.InsertOne(ctx, comment)
	return err
}
