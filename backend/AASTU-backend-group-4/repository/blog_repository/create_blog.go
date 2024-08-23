package blog_repository

import (
	"blog-api/domain"
	"context"
)

func (r *BlogRepository) CreateBlog(ctx context.Context, blog *domain.Blog) error {
	_, err := r.collection.InsertOne(ctx, blog)
	return err
}
