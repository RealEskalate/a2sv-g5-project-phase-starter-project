package blog_usecase

import (
	"blog-api/domain"
	"context"
)

func (bu *BlogUsecase) FilterBlog(ctx context.Context, filter domain.FilterRequest) ([]*domain.Blog, error) {
	// ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	// defer cancel()

	blogs, err := bu.blogRepo.FilterBlog(ctx, filter)
	if err != nil {
		return []*domain.Blog{}, err
	}
	return blogs, nil
}
