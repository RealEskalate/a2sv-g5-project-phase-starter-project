package usecase

import (
	"Blog_Starter/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCase struct {
	blogRepo domain.BlogRepository
	userRepo domain.UserRepository
}

func NewBlogUseCase(blogRepo domain.BlogRepository, userRepo domain.UserRepository) domain.BlogUseCase {
	return &BlogUseCase{
		blogRepo: blogRepo,
		userRepo: userRepo,
	}
}

func (uc *BlogUseCase) CreateBlog(c context.Context, blog *domain.BlogCreate) (*domain.Blog, error) {
	// implementation
	//system checks if content length constraints.
	if len(blog.Content) < 10 {
		return nil, errors.New("content length should be greater than 10")
	}

	// check the user existance
	user, err := uc.userRepo.GetUserByID(c, blog.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	userID, err := primitive.ObjectIDFromHex(blog.UserID)
	if err != nil {
		return nil, err
	}

	blogModel := &domain.Blog{
		UserID:  userID,
		Title:   blog.Title,
		Content: blog.Content,
		Tags:    blog.Tags,
		Author:  user.Username,
	}
	return uc.blogRepo.CreateBlog(c, blogModel)
}

func (uc *BlogUseCase) GetBlogByID(c context.Context, blogID string) (*domain.Blog,error) {
	updatedBlog, err := uc.blogRepo.IncrementViewCount(c, blogID)
	return updatedBlog, err
}

func (uc *BlogUseCase) GetAllBlog(c context.Context, skip int64, limit int64, sortBy string) ([]*domain.Blog, *domain.PaginationMetadata, error) {

	
	// if not make it default to createtimestamp
	if sortBy != "createtimestamp" && sortBy != "average_rating" && sortBy != "total_rating" && sortBy != "view_count" && sortBy != "like_count" && sortBy != "comment_count" {
		sortBy = "createtimestamp"
	}

	return uc.blogRepo.GetAllBlog(c, skip, limit, sortBy)
}

func (uc *BlogUseCase) UpdateBlog(c context.Context, blog *domain.BlogUpdate, blogID string) (*domain.Blog, error) {
	// implementation i want only to change the title, content and tags
	existedBlog, err := uc.blogRepo.GetBlogByID(c, blogID)
	if err != nil {
		return nil, err
	}

	// check if the user is the owner of the blog
	blogUserId, err := primitive.ObjectIDFromHex(blog.UserID)
	if err != nil {
		return nil, err
	}

	if existedBlog.UserID != blogUserId {
		return nil, errors.New("user is not the owner of the blog")
	}

	updatedBlog := &domain.BlogUpdate{
		Title:   blog.Title,
		Content: blog.Content,
		Tags:    blog.Tags,
	}

	return uc.blogRepo.UpdateBlog(c, updatedBlog, blogID)
}

func (uc *BlogUseCase) DeleteBlog(c context.Context, blogID string, userId string) error {
	// implementation
	// check if the user is the owner of the blog
	blogUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	existedBlog, err := uc.blogRepo.GetBlogByID(c, blogID)
	if err != nil {
		return err
	}

	if existedBlog.UserID != blogUserId {
		return errors.New("user is not the owner of the blog")
	}

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
