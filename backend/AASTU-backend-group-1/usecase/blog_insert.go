package usecase

import (
	"blogs/domain"
)

// InsertBlog implements domain.BlogUsecase.
func (b *BlogUsecase) InsertBlog(blog *domain.Blog) error {
	err := b.BlogRepo.InsertBlog(blog)
	if err != nil {
		return err
	}

	return nil
}
