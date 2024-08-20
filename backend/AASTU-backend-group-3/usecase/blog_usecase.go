package usecase

import (
	"context"
	"errors"
	"group3-blogApi/domain"
	"log"
)

type BlogUsecaseImpl struct {
	blogRepo domain.BlogRepository
}

func NewBlogUsecase(blogRepo domain.BlogRepository) domain.BlogUsecase {
	return &BlogUsecaseImpl{
		blogRepo: blogRepo,
	}
}

func (uc *BlogUsecaseImpl) CreateBlog(ctx context.Context, blog domain.Blog) (string, error) {
	id, err := uc.blogRepo.CreateBlog(ctx, blog)
	if err != nil {
		log.Printf("Error creating blog: %v", err)
		return "", errors.New("failed          to create blog")
	}
	return id, nil
}

func (uc *BlogUsecaseImpl) GetBlogByID(ctx context.Context, id string) (*domain.Blog, error) {
	blog, err := uc.blogRepo.GetBlogByID(ctx, id)
	if err != nil {
		log.Printf("Error retrieving blog by ID: %v", err)
		return nil, errors.New("failed to retrieve blog")
	}
	// if blog == nil {
	// 	return nil, errors.New("blog not found")
	// }
	return blog, nil
}

func (uc *BlogUsecaseImpl) UpdateBlog(ctx context.Context, blog domain.Blog) error {
	err := uc.blogRepo.UpdateBlog(ctx, blog)
	if err != nil {
		log.Printf("Error updating blog: %v", err)
		return errors.New("failed to update blog")
	}
	return nil
}

func (uc *BlogUsecaseImpl) DeleteBlog(ctx context.Context, id string) error {
	err := uc.blogRepo.DeleteBlog(ctx, id)
	if err != nil {
		log.Printf("Error deleting blog: %v", err)
		return errors.New("failed to delete blog")
	}
	return nil
}

func (uc *BlogUsecaseImpl) GetBlogs(ctx context.Context, offset int64, limit int64, sortBy string) ([]domain.Blog, error) {
    blogs, err := uc.blogRepo.GetBlogs(ctx, offset, limit, sortBy)
    if err != nil {
        return nil, err
    }
    return blogs, nil
}


func (uc *BlogUsecaseImpl) SearchBlogs(ctx context.Context, query string, filters map[string]interface{}) ([]domain.Blog, error) {
	blogs, err := uc.blogRepo.SearchBlogs(ctx, query, filters)
	if err != nil {
		log.Printf("Error searching blogs: %v", err)
		return nil, errors.New("failed to search blogs")
	}
	return blogs, nil
}

func (uc *BlogUsecaseImpl) FilterBlogs(ctx context.Context, filters map[string]interface{}, sortBy string) ([]domain.Blog, error) {
	blogs, err := uc.blogRepo.FilterBlogs(ctx, filters, sortBy)
	if err != nil {
		log.Printf("Error filtering blogs: %v", err)
		return nil, errors.New("failed to filter blogs")
	}
	return blogs, nil
}

func (uc *BlogUsecaseImpl) TrackPopularity(ctx context.Context, blogID string, action string) error {
	err := uc.blogRepo.TrackPopularity(ctx, blogID, action)
	if err != nil {
		log.Printf("Error tracking blog popularity: %v", err)
		return errors.New("failed to track blog popularity")
	}
	return nil
}
