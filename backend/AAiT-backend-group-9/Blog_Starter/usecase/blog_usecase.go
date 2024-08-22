package usecase

import (
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCase struct {
	blogRepo  domain.BlogRepository
	userRepo  domain.UserRepository
	cacheServ utils.Cache
}

func NewBlogUseCase(blogRepo domain.BlogRepository, userRepo domain.UserRepository, cacheServ utils.Cache) domain.BlogUseCase {
	return &BlogUseCase{
		blogRepo:  blogRepo,
		userRepo:  userRepo,
		cacheServ: cacheServ,
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

func (uc *BlogUseCase) GetBlogByID(c context.Context, blogID string) (*domain.Blog, error) {
	// check if the blog exist in the cache else go for repository and save it in the cache
	blog, found := uc.cacheServ.Get(blogID)

	if found {
		return blog.(*domain.Blog), nil
	}

	updatedBlog, err := uc.blogRepo.IncrementViewCount(c, blogID)
	if err != nil {
		return nil, err
	}
	uc.cacheServ.Set(blogID, updatedBlog, 0)

	return updatedBlog, err
}

func (uc *BlogUseCase) GetAllBlog(c context.Context, skip int64, limit int64, sortBy string) ([]*domain.Blog, *domain.PaginationMetadata, error) {

	// if not make it default to createtimestamp
	if sortBy != "createtimestamp" && sortBy != "average_rating" && sortBy != "total_rating" && sortBy != "view_count" && sortBy != "like_count" && sortBy != "comment_count" {
		sortBy = "createtimestamp"
	}

	cacheKey := "allblogs" + sortBy + string(skip) + string(limit)
	paginationKey := "pagination" + sortBy + string(skip) + string(limit)
	blogs, found := uc.cacheServ.Get(cacheKey)
	paginations, foundPagination := uc.cacheServ.Get(paginationKey)

	// if both found get the data from the cache
	if found && foundPagination {
		return blogs.([]*domain.Blog), paginations.(*domain.PaginationMetadata), nil
	}

	allBlog, pagination, err := uc.blogRepo.GetAllBlog(c, skip, limit, sortBy)
	if err != nil {
		return nil, nil, err
	}
	uc.cacheServ.Set(cacheKey, allBlog, 0)
	uc.cacheServ.Set(paginationKey, pagination, 0)
	return allBlog, pagination, nil
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
	newUpdatedBlog, err := uc.blogRepo.UpdateBlog(c, updatedBlog, blogID)
	if err != nil {
		return newUpdatedBlog, err
	}

	// delete the blog from the cache
	// set the blog
	uc.cacheServ.Delete(blogID)
	uc.cacheServ.Set(blogID, newUpdatedBlog, 0)
	return newUpdatedBlog, err
}

func (uc *BlogUseCase) DeleteBlog(c context.Context, blogID string, userId string) error {
	// implementation
	// check if the user is the owner of the blog
	// delete the blog from the cache
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

	uc.cacheServ.Delete(blogID)

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
