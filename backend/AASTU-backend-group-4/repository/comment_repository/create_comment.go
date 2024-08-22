package comment_repository

import (
	"blog-api/domain"
	"context"
)

func (cr *CommentRepository) CreateBlog(ctx context.Context, comment *domain.Comment) error {
	_, err := cr.collection.InsertOne(ctx, comment)
	return err
}
