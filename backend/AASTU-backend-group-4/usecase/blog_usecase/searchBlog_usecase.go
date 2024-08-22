package blog_usecase

import (
	"blog-api/domain"

	"context"
)

func (bu *BlogUsecase) SearchBlog(ctx context.Context, filter map[string]string) ([]*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	return bu.blogRepo.SearchBlog(ctx, filter)
}
