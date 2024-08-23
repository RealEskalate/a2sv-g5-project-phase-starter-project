package usecase_test

import (
	"Blog_Starter/domain/mocks"
	"Blog_Starter/domain"
	"Blog_Starter/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCaseSuite struct {
	suite.Suite
	blogUseCase domain.BlogUseCase
	blogRepo    *mocks.BlogRepository
	userRepo    *mocks.UserRepository
}

func (suite *BlogUseCaseSuite) SetupTest() {
	suite.blogRepo = new(mocks.BlogRepository)
	suite.userRepo = new(mocks.UserRepository)
	suite.blogUseCase = usecase.NewBlogUseCase(suite.blogRepo, suite.userRepo, time.Second*5)
}

func (suite *BlogUseCaseSuite) TestCreateBlog_Success() {
	ctx := context.TODO()
	userID := primitive.NewObjectID()

	blogCreate := &domain.BlogCreate{
		UserID:  userID.Hex(),
		Title:   "Test Blog",
		Content: "This is a test blog content.",
		Tags:    []string{"test", "blog"},
	}

	user := &domain.User{
		UserID:      userID,
		Username: "testuser",
	}

	suite.userRepo.On("GetUserByID", mock.Anything, blogCreate.UserID).Return(user, nil)

	blog := &domain.Blog{
		UserID:  primitive.NewObjectID(),
		Title:   blogCreate.Title,
		Content: blogCreate.Content,
		Tags:    blogCreate.Tags,
		Author:  user.Username,
	}

	suite.blogRepo.On("CreateBlog", mock.Anything, mock.AnythingOfType("*domain.Blog")).Return(blog, nil)

	createdBlog, err := suite.blogUseCase.CreateBlog(ctx, blogCreate)

	suite.NoError(err)
	suite.NotNil(createdBlog)
	suite.Equal(blog.Title, createdBlog.Title)
	suite.userRepo.AssertExpectations(suite.T())
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestCreateBlog_UserNotFound() {
	ctx := context.TODO()
	userID := primitive.NewObjectID().Hex()

	blogCreate := &domain.BlogCreate{
		UserID:  userID,
		Title:   "Test Blog",
		Content: "This is a test blog content.",
		Tags:    []string{"test", "blog"},
	}

	suite.userRepo.On("GetUserByID", mock.Anything, blogCreate.UserID).Return(nil, errors.New("user not found"))

	createdBlog, err := suite.blogUseCase.CreateBlog(ctx, blogCreate)

	suite.Error(err)
	suite.Nil(createdBlog)
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestGetBlogByID_Success() {
	ctx := context.TODO()
	blogID := primitive.NewObjectID().Hex()

	expectedBlog := &domain.Blog{
		Title: "Test Blog",
	}

	suite.blogRepo.On("IncrementViewCount", mock.Anything, blogID).Return(expectedBlog, nil)

	blog, err := suite.blogUseCase.GetBlogByID(ctx, blogID)

	suite.NoError(err)
	suite.NotNil(blog)
	suite.Equal(expectedBlog.Title, blog.Title)
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestGetBlogByID_NotFound() {
	ctx := context.TODO()
	blogID := primitive.NewObjectID().Hex()

	suite.blogRepo.On("IncrementViewCount", mock.Anything, blogID).Return(nil, errors.New("blog not found"))

	blog, err := suite.blogUseCase.GetBlogByID(ctx, blogID)

	suite.Error(err)
	suite.Nil(blog)
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestUpdateBlog_Success() {
	ctx := context.TODO()
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	existingBlog := &domain.Blog{
		UserID: primitive.NewObjectID(),
	}

	updatedBlog := &domain.BlogUpdate{
		UserID:  userID,
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"updated", "tags"},
	}

	suite.blogRepo.On("GetBlogByID", mock.Anything, blogID).Return(existingBlog, nil)
	suite.blogRepo.On("UpdateBlog", mock.Anything, updatedBlog, blogID).Return(&domain.Blog{}, nil)

	blog, err := suite.blogUseCase.UpdateBlog(ctx, updatedBlog, blogID)

	suite.NoError(err)
	suite.NotNil(blog)
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestUpdateBlog_UserNotOwner() {
	ctx := context.TODO()
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	existingBlog := &domain.Blog{
		UserID: primitive.NewObjectID(), // Different user ID
	}

	updatedBlog := &domain.BlogUpdate{
		UserID:  userID,
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"updated", "tags"},
	}

	suite.blogRepo.On("GetBlogByID", mock.Anything, blogID).Return(existingBlog, nil)

	blog, err := suite.blogUseCase.UpdateBlog(ctx, updatedBlog, blogID)

	suite.Error(err)
	suite.Nil(blog)
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestDeleteBlog_Success() {
	ctx := context.TODO()
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	role := "admin"

	existingBlog := &domain.Blog{
		UserID: primitive.NewObjectID(),
	}

	suite.blogRepo.On("GetBlogByID", mock.Anything, blogID).Return(existingBlog, nil)
	suite.blogRepo.On("DeleteBlog", mock.Anything, blogID).Return(nil)

	err := suite.blogUseCase.DeleteBlog(ctx, blogID, userID, role)

	suite.NoError(err)
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestDeleteBlog_UserNotOwner() {
	ctx := context.TODO()
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	role := "user"

	existingBlog := &domain.Blog{
		UserID: primitive.NewObjectID(), // Different user ID
	}

	suite.blogRepo.On("GetBlogByID", mock.Anything, blogID).Return(existingBlog, nil)

	err := suite.blogUseCase.DeleteBlog(ctx, blogID, userID, role)

	suite.Error(err)
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestGetAllBlog_Success() {
	ctx := context.TODO()
	skip := int64(0)
	limit := int64(10)
	sortBy := "createtimestamp"

	expectedBlogs := []*domain.Blog{
		{Title: "Blog 1"},
		{Title: "Blog 2"},
	}

	suite.blogRepo.On("GetAllBlog", mock.Anything, skip, limit, sortBy).Return(expectedBlogs, &domain.PaginationMetadata{}, nil)

	blogs, pagination, err := suite.blogUseCase.GetAllBlog(ctx, skip, limit, sortBy)

	suite.NoError(err)
	suite.Equal(expectedBlogs, blogs)
	suite.NotNil(pagination)
	suite.blogRepo.AssertExpectations(suite.T())
}
func (suite *BlogUseCaseSuite) TestSearchBlogs_Success() {
	ctx := context.TODO()

	searchRequest := &domain.BlogSearchRequest{
		Title:  "Test",
		Author: "Author",
	}

	expectedBlogs := []*domain.Blog{
		{Title: "Blog 1"},
	}

	suite.blogRepo.On("SearchBlogs", mock.Anything, searchRequest).Return(expectedBlogs, nil)

	blogs, err := suite.blogUseCase.SearchBlogs(ctx, searchRequest)

	suite.NoError(err)
	suite.Equal(expectedBlogs, blogs)
	suite.blogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseSuite) TestFilterBlogs_Success() {
	ctx := context.TODO()

	filterRequest := &domain.BlogFilterRequest{
		Tags:          []string{"test"},
		LikeLowerRange:  10,
		ViewLowerRange:  100,
	}

	expectedBlogs := []*domain.Blog{
		{Title: "Blog 1"},
	}

	suite.blogRepo.On("FilterBlogs", mock.Anything, filterRequest).Return(expectedBlogs, nil)

	blogs, err := suite.blogUseCase.FilterBlogs(ctx, filterRequest)

	suite.NoError(err)
	suite.Equal(expectedBlogs, blogs)
	suite.blogRepo.AssertExpectations(suite.T())
}


func TestBlogUseCaseSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseSuite))
}
