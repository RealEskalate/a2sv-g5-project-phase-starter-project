package blog_usecase

import (
	"blog-api/domain"
	"errors"
	"time"

	"context"
)

func (u *BlogUsecase) SearchBlogs(ctx context.Context, author, title string, tags []string, dateFrom, dateTo string) ([]domain.Blog, error) {
	// ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	// defer cancel()

	filters := domain.BlogFilters{
		Author: author,
		Title:  title,
		Tags:   tags,
	}

	var err error
	if dateFrom != "" {
		filters.DateFrom, err = parseDate(dateFrom)
		if err != nil {
			return nil, errors.New("invalid dateFrom format")
		}
	}

	if dateTo != "" {
		filters.DateTo, err = parseDate(dateTo)
		if err != nil {
			return nil, errors.New("invalid dateTo format")
		}
	}

	if filters.DateFrom != nil && filters.DateTo == nil {
		now := time.Now()
		filters.DateTo = &now
	}

	if filters.DateTo != nil && filters.DateFrom == nil {
		epoch := time.Unix(0, 0)
		filters.DateFrom = &epoch
	}

	return u.blogRepo.SearchBlogs(ctx, filters)
}

func parseDate(dateStr string) (*time.Time, error) {
	layout := "2006/01/02 15:04"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
