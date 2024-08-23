package blogusecase

import "blogs/domain"

// GetBlogs implements domain.BlogUsecase.
func (b *BlogUsecase) GetBlogs(sortBy string, page int, limit int, reverse bool) ([]*domain.Blog, int, error) {
	var blogs []*domain.Blog
	var err error

	if sortBy == "popularity" {
		blogs, err = b.BlogRepo.GetBlogsByPopularity(page, limit, reverse)
		if err != nil {
			return nil, 0, err
		}

	} else {
		blogs, err = b.BlogRepo.GetBlogsByRecent(page, limit, reverse)
		if err != nil {
			return nil, 0, err
		}
	}

	total, err := b.BlogRepo.GetTotalBlogs()
	pages := (total-1)/limit + 1
	if err != nil {
		return nil, 0, err
	}

	return blogs, pages, nil
}
