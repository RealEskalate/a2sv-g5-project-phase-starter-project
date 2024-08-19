package blogusecase

import "blogs/domain"

// SearchBlog implements domain.BlogUsecase.
func (b *BlogUsecase) SearchBlog(title string, author string, tags []string) ([]*domain.Blog, error) {
	blogs, err := b.BlogRepo.SearchBlog(title, author, tags)
	if err != nil {
		return nil, err
	}

	return blogs, nil
	
}
