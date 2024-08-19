package usecase

import (
	"blogs/domain"
	"time"
)

// FilterBlog implements domain.BlogUsecase.
func (b *BlogUsecase) FilterBlog(tags []string, dateFrom time.Time, dateTo time.Time) ([]*domain.Blog, error) {
	blogs, err := b.BlogRepo.FilterBlog(tags, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	return blogs, nil
	
}