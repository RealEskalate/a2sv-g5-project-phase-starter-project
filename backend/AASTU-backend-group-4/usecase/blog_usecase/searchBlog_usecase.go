package blog_usecase

import (
	"blog-api/domain"
	"context"
	"errors"
)

func (bu *BlogUsecase) SearchBlog(ctx context.Context, title string, author string) ([]*domain.Blog, error) {
	if title == "" && author == "" {
		return nil, errors.New("at least one search criterion (title or author) must be provided")
	}

	return bu.blogRepo.SearchBlog(ctx, title, author)
}
