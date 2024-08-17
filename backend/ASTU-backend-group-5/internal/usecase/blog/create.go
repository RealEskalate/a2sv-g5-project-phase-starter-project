package blog

import (
	"blogApp/internal/domain"
	"blogApp/internal/repository"
	"context"
	"errors"
	"fmt"
	"log"
)

// blogUseCase implements the BlogUseCase interface
type blogUseCase struct {
	repo repository.BlogRepository
}

// NewBlogUseCase creates a new instance of BlogUseCase
func NewBlogUseCase(repo repository.BlogRepository) BlogUseCase {
	return &blogUseCase{
		repo: repo,
	}
}

// CreateBlog creates a new blog
func (u *blogUseCase) CreateBlog(ctx context.Context, blog *domain.Blog) error {
	if blog == nil {
		return errors.New("blog cannot be nil")
	}
	// Add more input validation if necessary
	err := u.repo.CreateBlog(ctx, blog)
	if err != nil {
		log.Printf("Error creating blog: %v", err)
		return fmt.Errorf("failed to create blog: %w", err)
	}
	return nil
}

// AddComment adds a comment to a blog
func (u *blogUseCase) AddComment(ctx context.Context, comment *domain.Comment) error {
	if comment == nil {
		return errors.New("comment cannot be nil")
	}
	err := u.repo.AddComment(ctx, comment)
	if err != nil {
		log.Printf("Error adding comment to blog with ID %s: %v", comment.BlogID.Hex(), err)
		return fmt.Errorf("failed to add comment: %w", err)
	}
	return nil
}

// AddLike adds a like to a blog
func (u *blogUseCase) AddLike(ctx context.Context, like *domain.Like) error {
	if like == nil {
		return errors.New("like cannot be nil")
	}
	err := u.repo.AddLike(ctx, like)
	if err != nil {
		log.Printf("Error adding like to blog with ID %s: %v", like.BlogID.Hex(), err)
		return fmt.Errorf("failed to add like: %w", err)
	}
	return nil
}

// AddView adds a view to a blog
func (u *blogUseCase) AddView(ctx context.Context, view *domain.View) error {
	if view == nil {
		return errors.New("view cannot be nil")
	}
	err := u.repo.AddView(ctx, view)
	if err != nil {
		log.Printf("Error adding view to blog with ID %s: %v", view.BlogID.Hex(), err)
		return fmt.Errorf("failed to add view: %w", err)
	}
	return nil
}
