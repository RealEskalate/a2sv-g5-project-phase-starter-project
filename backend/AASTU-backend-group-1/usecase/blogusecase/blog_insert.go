package blogusecase

import (
	"blogs/domain"
)

// InsertBlog implements domain.BlogUsecase.
func (b *BlogUsecase) InsertBlog(blog *domain.Blog) (*domain.Blog, error) {
	newblog, err := b.BlogRepo.InsertBlog(blog)

	if err != nil {
		return nil, err
	}

	return newblog, nil
}
