package blog

import (

	//"blogApp/internal/ai"

	"blogApp/internal/domain"
	"blogApp/internal/repository"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (u *blogUseCase) CreateBlog(ctx context.Context, blog *domain.Blog, authorId string) error {
	if blog == nil {
		return errors.New("blog cannot be nil")
	}
	Author, err := primitive.ObjectIDFromHex(authorId)
	if err != nil {
		log.Printf("Error creating blog: %v", err)
		return fmt.Errorf("failed to create blog: %w", err)
	}
	blog.Author = Author
	blog.CreatedAt = primitive.NewDateTimeFromTime(time.Now())



	//is_valid, message, err := ai.ModerateBlog(blog.Content, blog.Title)
	// if err != nil {
	// 	log.Printf("Error moderating blog: %v", err)
	// 	return fmt.Errorf("failed to moderate blog: %w", err)
	// }
	// if !is_valid {
	// 	return errors.New(message)
	// }

	err = u.repo.CreateBlog(ctx, blog)
	if err != nil {
		log.Printf("Error creating blog: %v", err)
		return fmt.Errorf("failed to create blog: %w", err)
	}
	return nil
}

// AddComment adds a comment to a blog
func (u *blogUseCase) AddComment(ctx context.Context, comment *domain.Comment, userId string) error {
	if comment == nil {
		return errors.New("comment cannot be nil")
	}
	comment.UserID, _ = primitive.ObjectIDFromHex(userId)

	blogObjectID, err := primitive.ObjectIDFromHex(comment.BlogID.Hex())
	if err != nil {
		return fmt.Errorf("invalid blog ID: %w", err)
	}
	comment.BlogID = blogObjectID

	err = u.repo.AddComment(ctx, comment)

	if err != nil {
		log.Printf("Error adding comment to blog with ID %s: %v", comment.BlogID.Hex(), err)
		return fmt.Errorf("failed to add comment: %w", err)
	}
	return nil
}

// AddLike adds a like to a blog
func (u *blogUseCase) AddLike(ctx context.Context, like *domain.Like, userId string) error {
	if like == nil {
		return errors.New("like cannot be nil")
	}
	alreadyLiked, _ := u.repo.HasUserLikedBlog(ctx, userId, like.BlogID.Hex())
	if alreadyLiked {
		return nil
	}
	like.UserID, _ = primitive.ObjectIDFromHex(userId)
	err := u.repo.AddLike(ctx, like)
	if err != nil {
		log.Printf("Error adding like to blog with ID %s: %v", like.BlogID.Hex(), err)
		return fmt.Errorf("failed to add like: %w", err)
	}
	return nil
}

// AddView adds a view to a blog
func (u *blogUseCase) AddView(ctx context.Context, view *domain.View, userId string) error {
	if view == nil {
		return errors.New("view cannot be nil")
	}
	AlreadyViewed, _ := u.repo.HasUserViewedBlog(ctx, userId, view.BlogID.Hex())
	if AlreadyViewed {
		return nil
	}
	view.UserID, _ = primitive.ObjectIDFromHex(userId)
	err := u.repo.AddView(ctx, view)
	if err != nil {
		log.Printf("Error adding view to blog with ID %s: %v", view.BlogID.Hex(), err)
		return fmt.Errorf("failed to add view: %w", err)
	}
	return nil
}

// GetAllTags retrieves all blog tags
func (u *blogUseCase) GetAllTags(ctx context.Context) ([]*domain.BlogTag, error) {
	tags, err := u.repo.GetAllTags(ctx)
	if err != nil {
		log.Printf("Error retrieving blog tags: %v", err)
		return nil, fmt.Errorf("failed to retrieve blog tags: %w", err)
	}
	return tags, nil
}

// CreateTag creates a new blog tag
func (u *blogUseCase) CreateTag(ctx context.Context, tag *domain.BlogTag) error {
	if tag == nil {
		return errors.New("tag cannot be nil")
	}
	err := u.repo.CreateTag(ctx, tag)
	if err != nil {
		log.Printf("Error creating blog tag: %v", err)
		return fmt.Errorf("failed to create blog tag: %w", err)
	}
	return nil
}

// UpdateTag updates an existing blog tag
func (u *blogUseCase) UpdateTag(ctx context.Context, id string, tag *domain.BlogTag) error {
	if tag == nil {
		return errors.New("tag cannot be nil")
	}
	tag.ID, _ = primitive.ObjectIDFromHex(id)
	err := u.repo.UpdateTag(ctx, id, tag)
	if err != nil {
		log.Printf("Error updating blog tag with ID %s: %v", id, err)
		return fmt.Errorf("failed to update blog tag: %w", err)
	}
	return nil
}

// DeleteTag deletes a blog tag
func (u *blogUseCase) DeleteTag(ctx context.Context, id string) error {
	err := u.repo.DeleteTag(ctx, id)
	if err != nil {
		log.Printf("Error deleting blog tag with ID %s: %v", id, err)
		return fmt.Errorf("failed to delete blog tag: %w", err)
	}
	return nil
}

// GetTagByID retrieves a blog tag by ID
func (u *blogUseCase) GetTagByID(ctx context.Context, id string) (*domain.BlogTag, error) {
	tag, err := u.repo.GetTagByID(ctx, id)
	if err != nil {
		log.Printf("Error retrieving blog tag with ID %s: %v", id, err)
		return nil, fmt.Errorf("failed to retrieve blog tag: %w", err)
	}
	return tag, nil
}
