package usecase

import (
	"Blog_Starter/domain"
	"context"
	"errors"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCase struct {
	blogRepo domain.BlogRepository
	userRepo domain.UserRepository
	timeout time.Duration
}

func NewBlogUseCase(blogRepo domain.BlogRepository, userRepo domain.UserRepository, timeout time.Duration) domain.BlogUseCase {
	return &BlogUseCase{
		blogRepo: blogRepo,
		userRepo: userRepo,
		timeout : timeout,
	}
}

func (uc *BlogUseCase) CreateBlog(c context.Context, blog *domain.BlogCreate) (*domain.Blog, error) {
	// implementation
	//system checks if content length constraints.
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()

	if len(blog.Content) < 10 {
		return nil, errors.New("content length should be greater than 10")
	}

	// check the user existence
	user, err := uc.userRepo.GetUserByID(ctx, blog.UserID)
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
	return uc.blogRepo.CreateBlog(ctx, blogModel)
}

func (uc *BlogUseCase) GetBlogByID(c context.Context, blogID string) (*domain.Blog,error) {
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()
	updatedBlog, err := uc.blogRepo.IncrementViewCount(ctx, blogID)
	return updatedBlog, err
}

func (uc *BlogUseCase) GetAllBlog(c context.Context, skip int64, limit int64, sortBy string) ([]*domain.Blog, *domain.PaginationMetadata, error) {

	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()
	// if not make it default to createtimestamp
	if sortBy != "createtimestamp" && sortBy != "average_rating" && sortBy != "total_rating" && sortBy != "view_count" && sortBy != "like_count" && sortBy != "comment_count" {
		sortBy = "createtimestamp"
	}

	return uc.blogRepo.GetAllBlog(ctx, skip, limit, sortBy)
}

func (uc *BlogUseCase) UpdateBlog(c context.Context, blog *domain.BlogUpdate, blogID string) (*domain.Blog, error) {
	// implementation i want only to change the title, content and tags
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()
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

	return uc.blogRepo.UpdateBlog(ctx, updatedBlog, blogID)
}

func (uc *BlogUseCase) DeleteBlog(c context.Context, blogID string, userId string, role string) error {
	// implementation
	// check if the user is the owner of the blog
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()
	blogUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	existedBlog, err := uc.blogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	if existedBlog.UserID != blogUserId  && role != "admin" && role != "superAdmin"{
		return errors.New("user is not the owner of the blog")
	}

	return uc.blogRepo.DeleteBlog(ctx, blogID)
}

// SearchBlogs implements domain.BlogUseCase.

func (uc *BlogUseCase) SearchBlogs(c context.Context, searchRequest *domain.BlogSearchRequest) ([]*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()
	filteredBlog, err := uc.blogRepo.SearchBlogs(ctx, searchRequest)
	return filteredBlog, err
}

func (uc *BlogUseCase) FilterBlogs(c context.Context, blogRequest *domain.BlogFilterRequest) ([]*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()
	if blogRequest.Date == nil && blogRequest.Tags == nil && blogRequest.LikeLowerRange == 0 && blogRequest.ViewLowerRange == 0 {
		return nil, errors.New("invalid request format")
	}
	filteredBlog, err := uc.blogRepo.FilterBlogs(ctx, blogRequest)
	return filteredBlog, err
}
