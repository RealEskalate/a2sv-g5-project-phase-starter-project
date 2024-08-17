package usecase

import (
	"blogs/domain"
	"time"
)

// AddComment implements domain.BlogUsecase.
func (b *BlogUsecase) AddComment(comment *domain.Comment) error {
	panic("unimplemented")
}

// AddLike implements domain.BlogUsecase.
func (b *BlogUsecase) AddLike(like *domain.Like) error {
	panic("unimplemented")
}

// AddView implements domain.BlogUsecase.
func (b *BlogUsecase) AddView(view *domain.View) error {
	panic("unimplemented")
}

// DeleteBlogByID implements domain.BlogUsecase.
func (b *BlogUsecase) DeleteBlogByID(id string) error {
	panic("unimplemented")
}

// FilterBlog implements domain.BlogUsecase.
func (b *BlogUsecase) FilterBlog(tags []string, dateFrom time.Time, dateTo time.Time) ([]*domain.Blog, error) {
	panic("unimplemented")
}

// GetBlog implements domain.BlogUsecase.
func (b *BlogUsecase) GetBlog(page int, size int) ([]*domain.Blog, error) {
	panic("unimplemented")
}

// InsertBlog implements domain.BlogUsecase.
func (b *BlogUsecase) InsertBlog(blog *domain.Blog) error {
	err := b.BlogRepo.InsertBlog(blog)
	if err != nil {
		return err
	}

	return nil
}

// SearchBlog implements domain.BlogUsecase.
func (b *BlogUsecase) SearchBlog(title string, author string, tags []string) ([]*domain.Blog, error) {
	panic("unimplemented")
}

// UpdateBlogByID implements domain.BlogUsecase.
func (b *BlogUsecase) UpdateBlogByID(id string, blog *domain.Blog) error {
	panic("unimplemented")
}
