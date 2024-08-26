package usecase_test

import (
	"blogApp/internal/domain"
	"blogApp/mocks/repository"
	"blogApp/internal/usecase/blog"
	"github.com/stretchr/testify/mock"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUseCaseTestSuite struct {
	suite.Suite
	blogUseCase blog.BlogUseCase // Using the BlogUseCase interface
	mockUserRepo *mocks.UserRepository
	mockBlogRepo *mocks.BlogRepository
	ctx          context.Context
}

func (suite *BlogUseCaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.mockBlogRepo = new(mocks.BlogRepository)
	suite.blogUseCase = blog.NewBlogUseCase(suite.mockBlogRepo, suite.mockUserRepo) // Using the existing constructor
	suite.ctx = context.Background()
}

func (suite *BlogUseCaseTestSuite) TestCreateBlogSuccess() {
	authorId := primitive.NewObjectID().Hex()
	blog := &domain.Blog{
		Title:   "Sample Blog Title",
		Content: "Sample Blog Content",
	}

	user := &domain.User{
		ID: primitive.NewObjectID(),
		Profile: domain.UserProfile{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	suite.mockUserRepo.On("FindUserById", suite.ctx, authorId).Return(user, nil)
	suite.mockBlogRepo.On("CreateBlog", suite.ctx, blog).Return(nil)

	err := suite.blogUseCase.CreateBlog(suite.ctx, blog, authorId)

	suite.NoError(err)
	suite.Equal(user.Profile.FirstName+" "+user.Profile.LastName, blog.AuthorName)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestCreateBlogUserNotFound() {
	authorId := primitive.NewObjectID().Hex()
	blog := &domain.Blog{
		Title:   "Sample Blog Title",
		Content: "Sample Blog Content",
	}

	suite.mockUserRepo.On("FindUserById", suite.ctx, authorId).Return(nil, nil)

	err := suite.blogUseCase.CreateBlog(suite.ctx, blog, authorId)

	suite.Error(err)
	suite.EqualError(err, "user not found")
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockBlogRepo.AssertNotCalled(suite.T(), "CreateBlog", suite.ctx, blog)
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogSuccess() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	userRole := "author"
	p_ID :=  primitive.NewObjectID()
	A_ID, _ := primitive.ObjectIDFromHex(userID)


	blog := &domain.GetSingleBlogDTO{
		ID:     p_ID,
		Author: A_ID,
	}

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(blog, nil)
	suite.mockBlogRepo.On("DeleteBlog", suite.ctx, blogID).Return(nil)

	err := suite.blogUseCase.DeleteBlog(suite.ctx, blogID, userID, userRole)

	suite.NoError(err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogNotFound() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	userRole := "author"

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(nil, nil)

	err := suite.blogUseCase.DeleteBlog(suite.ctx, blogID, userID, userRole)

	suite.Error(err)
	suite.EqualError(err, "blog not found")
	suite.mockBlogRepo.AssertExpectations(suite.T())
	suite.mockBlogRepo.AssertNotCalled(suite.T(), "DeleteBlog", suite.ctx, blogID)
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogUnauthorized() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	differentUserID := primitive.NewObjectID().Hex() // Simulate a different user
	userRole := "author"
	A_ID, _ := primitive.ObjectIDFromHex(differentUserID) 

	blog := &domain.GetSingleBlogDTO{
		ID:     primitive.NewObjectID(),
		Author: A_ID,
	}

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(blog, nil)

	err := suite.blogUseCase.DeleteBlog(suite.ctx, blogID, userID, userRole)

	suite.Error(err)
	suite.EqualError(err, "you are not authorized to delete this blog")
	suite.mockBlogRepo.AssertExpectations(suite.T())
	suite.mockBlogRepo.AssertNotCalled(suite.T(), "DeleteBlog", suite.ctx, blogID)
}

func (suite *BlogUseCaseTestSuite) TestDeleteBlogFailureToDelete() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	userRole := "author"
	A_ID, _ := primitive.ObjectIDFromHex(userID) 

	blog := &domain.GetSingleBlogDTO{
		ID:     primitive.NewObjectID(),
		Author: A_ID,
	}

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(blog, nil)
	suite.mockBlogRepo.On("DeleteBlog", suite.ctx, blogID).Return(errors.New("failed to delete blog"))

	err := suite.blogUseCase.DeleteBlog(suite.ctx, blogID, userID, userRole)

	suite.Error(err)
	suite.EqualError(err, "failed to delete blog: failed to delete blog")
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetBlogByIDSuccess() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	expectedBlog := &domain.GetSingleBlogDTO{
		ID:    primitive.NewObjectID(),
		Title: "Sample Blog",
	}

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(expectedBlog, nil)
	suite.mockBlogRepo.On("HasUserViewedBlog", suite.ctx, userID, blogID).Return(false, nil)
	suite.mockBlogRepo.On("AddView", suite.ctx, mock.Anything).Return(nil)
	suite.mockBlogRepo.On("HasUserLikedBlog", suite.ctx, userID, blogID).Return(true, nil)

	result, err := suite.blogUseCase.GetBlogByID(suite.ctx, blogID, userID)

	suite.NoError(err)
	suite.Equal(expectedBlog, result)
	suite.True(result.Liked)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetBlogByIDUserAlreadyViewed() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	expectedBlog := &domain.GetSingleBlogDTO{
		ID:    primitive.NewObjectID(),
		Title: "Sample Blog",
	}

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(expectedBlog, nil)
	suite.mockBlogRepo.On("HasUserViewedBlog", suite.ctx, userID, blogID).Return(true, nil)
	suite.mockBlogRepo.On("HasUserLikedBlog", suite.ctx, userID, blogID).Return(false, nil)

	result, err := suite.blogUseCase.GetBlogByID(suite.ctx, blogID, userID)

	suite.NoError(err)
	suite.Equal(expectedBlog, result)
	suite.False(result.Liked)
	suite.mockBlogRepo.AssertExpectations(suite.T())
	suite.mockBlogRepo.AssertNotCalled(suite.T(), "AddView", suite.ctx, mock.Anything)
}

func (suite *BlogUseCaseTestSuite) TestGetBlogByIDErrorCheckingUserViewed() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	expectedBlog := &domain.GetSingleBlogDTO{
		ID:    primitive.NewObjectID(),
		Title: "Sample Blog",
	}

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(expectedBlog, nil)
	suite.mockBlogRepo.On("HasUserViewedBlog", suite.ctx, userID, blogID).Return(false, errors.New("failed to check view"))

	result, err := suite.blogUseCase.GetBlogByID(suite.ctx, blogID, userID)

	suite.Error(err)
	suite.Nil(result)
	suite.EqualError(err, "failed to check if user has viewed blog: failed to check view")
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetBlogByIDErrorAddingView() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	expectedBlog := &domain.GetSingleBlogDTO{
		ID:    primitive.NewObjectID(),
		Title: "Sample Blog",
	}

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(expectedBlog, nil)
	suite.mockBlogRepo.On("HasUserViewedBlog", suite.ctx, userID, blogID).Return(false, nil)
	suite.mockBlogRepo.On("AddView", suite.ctx, mock.Anything).Return(errors.New("failed to add view"))

	result, err := suite.blogUseCase.GetBlogByID(suite.ctx, blogID, userID)

	suite.Error(err)
	suite.Nil(result)
	suite.EqualError(err, "failed to create view: failed to add view")
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetBlogByIDErrorCheckingUserLiked() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	expectedBlog := &domain.GetSingleBlogDTO{
		ID:    primitive.NewObjectID(),
		Title: "Sample Blog",
	}

	suite.mockBlogRepo.On("GetBlogByID", suite.ctx, blogID).Return(expectedBlog, nil)
	suite.mockBlogRepo.On("HasUserViewedBlog", suite.ctx, userID, blogID).Return(false, nil)
	suite.mockBlogRepo.On("AddView", suite.ctx, mock.Anything).Return(nil)
	suite.mockBlogRepo.On("HasUserLikedBlog", suite.ctx, userID, blogID).Return(false, errors.New("failed to check like"))

	result, err := suite.blogUseCase.GetBlogByID(suite.ctx, blogID, userID)

	suite.Error(err)
	suite.Nil(result)
	suite.EqualError(err, "failed to check if user has liked blog: failed to check like")
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestUpdateBlogSuccess() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	A_ID, _ := primitive.ObjectIDFromHex(userID)

	updatedBlog := &domain.Blog{
		ID:      primitive.NewObjectID(),
		Title:   "Updated Title",
		Content: "Updated Content",
		Author:  A_ID,
	}

	suite.mockBlogRepo.On("UpdateBlog", suite.ctx, blogID, updatedBlog).Return(nil)

	err := suite.blogUseCase.UpdateBlog(suite.ctx, blogID, updatedBlog, userID)

	suite.NoError(err)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestUpdateBlogNilBlog() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	err := suite.blogUseCase.UpdateBlog(suite.ctx, blogID, nil, userID)

	suite.Error(err)
	suite.EqualError(err, "blog cannot be nil")
	suite.mockBlogRepo.AssertNotCalled(suite.T(), "UpdateBlog", suite.ctx, blogID, mock.Anything)
}

func (suite *BlogUseCaseTestSuite) TestUpdateBlogUnauthorized() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	updatedBlog := &domain.Blog{
		ID:      primitive.NewObjectID(),
		Title:   "Updated Title",
		Content: "Updated Content",
		Author:  primitive.NewObjectID(), // Different user ID
	}

	err := suite.blogUseCase.UpdateBlog(suite.ctx, blogID, updatedBlog, userID)

	suite.Error(err)
	suite.EqualError(err, "you are not authorized to update this blog")
	suite.mockBlogRepo.AssertNotCalled(suite.T(), "UpdateBlog", suite.ctx, blogID, updatedBlog)
}

func (suite *BlogUseCaseTestSuite) TestUpdateBlogFailure() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()
	A_ID, _ := primitive.ObjectIDFromHex(userID)


	updatedBlog := &domain.Blog{
		ID:      primitive.NewObjectID(),
		Title:   "Updated Title",
		Content: "Updated Content",
		Author:  A_ID,
	}

	suite.mockBlogRepo.On("UpdateBlog", suite.ctx, blogID, updatedBlog).Return(errors.New("update failed"))

	err := suite.blogUseCase.UpdateBlog(suite.ctx, blogID, updatedBlog, userID)

	suite.Error(err)
	suite.EqualError(err, "failed to update blog: update failed")
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetUserBlogsSuccess() {
	userID := "some-user-id"
	page := 1
	pageSize := 10

	expectedBlogs := []*domain.Blog{
		{
			ID:       primitive.NewObjectID(),
			Title:   "First Blog",
			Content: "Content of the first blog",
		},
		{
			ID:       primitive.NewObjectID(),
			Title:   "Second Blog",
			Content: "Content of the second blog",
		},
	}

	suite.mockBlogRepo.On("GetUserBlogs", suite.ctx, userID, page, pageSize).Return(expectedBlogs, nil)

	blogs, err := suite.blogUseCase.GetUserBlogs(suite.ctx, userID, page, pageSize)

	suite.NoError(err)
	suite.Equal(expectedBlogs, blogs)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetUserBlogsError() {
	userID := "some-user-id"
	page := 1
	pageSize := 10

	suite.mockBlogRepo.On("GetUserBlogs", suite.ctx, userID, page, pageSize).Return(nil, errors.New("failed to retrieve blogs"))

	blogs, err := suite.blogUseCase.GetUserBlogs(suite.ctx, userID, page, pageSize)

	suite.Error(err)
	suite.Nil(blogs)
	suite.EqualError(err, "failed to retrieve blogs: failed to retrieve blogs")
	suite.mockBlogRepo.AssertExpectations(suite.T())
}

func (suite *BlogUseCaseTestSuite) TestGetUserBlogsNoBlogs() {
	userID := "some-user-id"
	page := 1
	pageSize := 10

	suite.mockBlogRepo.On("GetUserBlogs", suite.ctx, userID, page, pageSize).Return([]*domain.Blog{}, nil)

	blogs, err := suite.blogUseCase.GetUserBlogs(suite.ctx, userID, page, pageSize)

	suite.NoError(err)
	suite.Empty(blogs)
	suite.mockBlogRepo.AssertExpectations(suite.T())
}


func TestBlogUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUseCaseTestSuite))
}

