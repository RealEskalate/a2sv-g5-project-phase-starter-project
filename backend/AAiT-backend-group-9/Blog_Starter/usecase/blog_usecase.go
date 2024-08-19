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

func (uc *BlogUseCase) GetBlogByID(c context.Context, blogID string) (*domain.Blog, error) {
	// implementation
	return uc.blogRepo.GetBlogByID(c, blogID)
}

func (uc *BlogUseCase) GetAllBlog(c context.Context, skip int64, limit int64, sortBy string) ([]*domain.Blog, *domain.PaginationMetadata, error) {
	// chech if th query sortBy is valid which means it must be one of the following
	/* 	CreatedAt time.Time          `json:"createtimestamp" bson:"createtimestamp"`

	AverageRating float64 `json:"average_rating" bson:"average_rating"` // do some math and modify when rating is called
	TotalRating   int     `json:"total_rating" bson:"total_rating"`     // add when rating is called
	ViewCount     int     `json:"view_count" bson:"view_count"`         //add when getbyid is used and blog is viewed
	LikeCount     int     `json:"like_count" bson:"like_count"`         // add when like is called and substarct when it is unliked
	CommentCount
	*/
	// if not make it default to createtimestamp
	if sortBy != "createtimestamp" || sortBy != "average_rating" || sortBy != "total_rating" || sortBy != "view_count" || sortBy != "like_count" || sortBy != "comment_count" {
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

func (uc *BlogUseCase) SearchBlogs(c context.Context, title string, author string) ([]*domain.Blog, error) {
	panic("unimplemented")
}
