package blogusecase

import "blogs/domain"

// GetBlogs implements domain.BlogUsecase.
func (b *BlogUsecase) GetBlogs(sortBy string, page int, limit int, reverse bool) ([]*domain.Blog, error) {
	var blogs []*domain.Blog
	var err error

	if sortBy == "popularity" {
		blogs, err = b.BlogRepo.GetBlogsByPopularity(page, limit, reverse)
		if err != nil {
			return nil, err
		}

	} else {
		blogs, err = b.BlogRepo.GetBlogsByRecent(page, limit, reverse)
		if err != nil {
			return nil, err
		}
	}

	return blogs, nil
}
