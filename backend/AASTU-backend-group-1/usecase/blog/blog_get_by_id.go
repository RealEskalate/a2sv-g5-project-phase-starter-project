package usecase

import "blogs/domain"

func (b *BlogUsecase) GetBlogByID(id string) (*domain.Blog, error) {

	blog, err := b.BlogRepo.GetBlogByID(id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}