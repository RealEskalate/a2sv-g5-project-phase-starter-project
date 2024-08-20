package usecase

import (
	"Blog_Starter/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCase struct {
	blogRepo domain.BlogRepository
}

// FilterBlogs implements domain.BlogUseCase.
func NewBlogUseCase(blogRepo domain.BlogRepository) domain.BlogUseCase {
	return &BlogUseCase{
		blogRepo: blogRepo,
	}
}

func (uc *BlogUseCase) CreateBlog(c context.Context, blog *domain.BlogCreate) (*domain.Blog, error) {
	// implementation
	objectID, _ := primitive.ObjectIDFromHex(blog.UserID)
	blogModel := &domain.Blog{
		UserID:	 objectID,
		Title:   blog.Title,
		Content: blog.Content,
		Tags:    blog.Tags,
	}
	return uc.blogRepo.CreateBlog(c, blogModel)
}

func (uc *BlogUseCase) GetBlogByID(c context.Context, blogID string) (error) {
	return uc.blogRepo.IncrementViewCount(c, blogID)
}

func (uc *BlogUseCase) GetAllBlog(c context.Context) ([]*domain.Blog, error) {
	// implementation
	return uc.blogRepo.GetAllBlog(c)
}

func (uc *BlogUseCase) UpdateBlog(c context.Context, blog *domain.BlogUpdate, blogID string) (*domain.Blog, error) {
	// implementation i want only to change the title, content and tags
	_, err := uc.blogRepo.GetBlogByID(c, blogID)
	if err != nil {
		return nil, errors.New("blog not found")
	}

	updatedBlog := &domain.BlogUpdate{
		Title:   blog.Title,
		Content: blog.Content,
		Tags:    blog.Tags,
	}

	return uc.blogRepo.UpdateBlog(c, updatedBlog, blogID)
}

func (uc *BlogUseCase) DeleteBlog(c context.Context, blogID string) error {
	// check whether the blog exists or not
	_, err := uc.blogRepo.GetBlogByID(c, blogID)
	if err != nil {
		return errors.New("blog not found")
	}
	// implementation
	return uc.blogRepo.DeleteBlog(c, blogID)
}

// SearchBlogs implements domain.BlogUseCase.

func (uc *BlogUseCase) SearchBlogs(ctx context.Context, searchRequest *domain.BlogSearchRequest) ([]*domain.Blog, error) {
	filteredBlog, err := uc.blogRepo.SearchBlogs(ctx, searchRequest)
	return filteredBlog, err
}

func (uc *BlogUseCase) FilterBlogs(ctx context.Context, blogRequest *domain.BlogFilterRequest) ([]*domain.Blog, error) {
	if blogRequest.Date == nil && blogRequest.Tags == nil && blogRequest.LikeLowerRange == 0 && blogRequest.ViewLowerRange == 0 {
		return nil, errors.New("invalid request format")
	}
	filteredBlog, err := uc.blogRepo.FilterBlogs(ctx, blogRequest)
	return filteredBlog, err
}
