package blog_repository

import (
	"blog-api/domain/blog"
	"context"
)

func (r *BlogRepository) CreateBlog(ctx context.Context, blog *blog.Blog) error {
	_, err := r.collection.InsertOne(context.TODO(), blog)
	return err
}
