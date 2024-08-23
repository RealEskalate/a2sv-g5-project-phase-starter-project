package usecase_test

// import (
// 	"context"
// 	"testing"
// 	"time"
// 	"fmt"

// 	"github.com/stretchr/testify/suite"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"blogApp/internal/domain"
// 	"blogApp/internal/usecase/blog"
// 	"blogApp/mocks/repository" 
// 	"github.com/stretchr/testify/mock"
// )


// type BlogUseCaseTestSuite struct {
// 	suite.Suite
// 	useCase    blog.BlogUseCase
// 	repoMock   *mocks.BlogRepository
// 	ctx        context.Context
// }


// func (suite *BlogUseCaseTestSuite) SetupTest() {
// 	suite.repoMock = new(mocks.BlogRepository)
// 	suite.useCase = blog.NewBlogUseCase(suite.repoMock)
// 	suite.ctx = context.TODO()
// }


// func (suite *BlogUseCaseTestSuite) TestCreateBlog_Success() {
// 	authorID := "abcdef1234567890abcdef12" 
// 	blog := &domain.Blog{
// 		ID:      primitive.NewObjectID(),
// 		Title:   "Test Blog",
// 		Content: []interface{}{"Content"},
// 		Tags:    []domain.BlogTag{},
// 	}
// 	authorObjectID, _ := primitive.ObjectIDFromHex(authorID)
// 	blog.Author = authorObjectID
// 	blog.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

// 	suite.repoMock.On("CreateBlog", mock.Anything, blog).Return(nil)

// 	err := suite.useCase.CreateBlog(context.Background(), blog, authorID)

// 	suite.NoError(err)
// 	suite.repoMock.AssertCalled(suite.T(), "CreateBlog", mock.Anything, blog)
// }


// func (suite *BlogUseCaseTestSuite) TestCreateBlog_Failure_CreateBlog() {
	
// 	authorID := "abcdef1234567890abcdef12"
// 	blog := &domain.Blog{
// 		ID:      primitive.NewObjectID(),
// 		Title:   "Test Blog",
// 		Content: []interface{}{"Content"},
// 		Tags:    []domain.BlogTag{},
// 	}
// 	expectedError := fmt.Errorf("some repository error")
// 	suite.repoMock.On("CreateBlog", mock.Anything, mock.Anything).Return(expectedError)

	
// 	err := suite.useCase.CreateBlog(context.Background(), blog, authorID)

// 	suite.Error(err)
// 	suite.EqualError(err, "failed to create blog: some repository error")
// 	suite.repoMock.AssertCalled(suite.T(), "CreateBlog", mock.Anything, mock.Anything)
// }

// func (suite *BlogUseCaseTestSuite) TestDeleteBlog_Success() {

// 	blogID := "1234567890abcdef12345678" 
// 	userID := "abcdef1234567890abcdef12"
// 	blogObjectID, _ := primitive.ObjectIDFromHex(blogID)
// 	authorObjectID, _ := primitive.ObjectIDFromHex(userID)
// 	expectedBlog := &domain.Blog{
// 		ID:     blogObjectID,
// 		Author: authorObjectID,
// 	}

// 	suite.repoMock.On("GetBlogByID", mock.Anything, blogID).Return(expectedBlog, nil)
// 	suite.repoMock.On("DeleteBlog", mock.Anything, blogID).Return(nil)

// 	err := suite.useCase.DeleteBlog(context.Background(), blogID, userID)

// 	suite.NoError(err)
// 	suite.repoMock.AssertCalled(suite.T(), "GetBlogByID", mock.Anything, blogID)
// 	suite.repoMock.AssertCalled(suite.T(), "DeleteBlog", mock.Anything, blogID)
// }

// func (suite *BlogUseCaseTestSuite) TestDeleteBlog_Failure_GetBlogByID() {

// 	blogID := "1234567890abcdef12345678"
// 	userID := "abcdef1234567890abcdef12"
// 	expectedError := fmt.Errorf("some repository error")

// 	suite.repoMock.On("GetBlogByID", mock.Anything, blogID).Return(nil, expectedError)

// 	err := suite.useCase.DeleteBlog(context.Background(), blogID, userID)

// 	suite.Error(err)
// 	suite.EqualError(err, "failed to retrieve blog: some repository error")
// 	suite.repoMock.AssertCalled(suite.T(), "GetBlogByID", mock.Anything, blogID)
// 	suite.repoMock.AssertNotCalled(suite.T(), "DeleteBlog", mock.Anything, blogID)
// }

// func (suite *BlogUseCaseTestSuite) TestDeleteBlog_Failure_Unauthorized() {
// 	blogID := "1234567890abcdef12345678"
// 	userID := "abcdef1234567890abcdef12"
// 	blogObjectID, _ := primitive.ObjectIDFromHex(blogID)
// 	authorObjectID, _ := primitive.ObjectIDFromHex("different_user_id") // Different author
// 	expectedBlog := &domain.Blog{
// 		ID:     blogObjectID,
// 		Author: authorObjectID,
// 	}

// 	suite.repoMock.On("GetBlogByID", mock.Anything, blogID).Return(expectedBlog, nil)

// 	err := suite.useCase.DeleteBlog(context.Background(), blogID, userID)

// 	suite.Error(err)
// 	suite.EqualError(err, "you are not authorized to delete this blog")
// 	suite.repoMock.AssertCalled(suite.T(), "GetBlogByID", mock.Anything, blogID)
// 	suite.repoMock.AssertNotCalled(suite.T(), "DeleteBlog", mock.Anything, blogID)
// }

// func (suite *BlogUseCaseTestSuite) TestDeleteBlog_Failure_DeleteBlog() {

// 	blogID := "1234567890abcdef12345678"
// 	userID := "abcdef1234567890abcdef12"
// 	blogObjectID, _ := primitive.ObjectIDFromHex(blogID)
// 	authorObjectID, _ := primitive.ObjectIDFromHex(userID)
// 	expectedBlog := &domain.Blog{
// 		ID:     blogObjectID,
// 		Author: authorObjectID,
// 	}
// 	expectedError := fmt.Errorf("some repository error")

// 	suite.repoMock.On("GetBlogByID", mock.Anything, blogID).Return(expectedBlog, nil)
// 	suite.repoMock.On("DeleteBlog", mock.Anything, blogID).Return(expectedError)

// 	err := suite.useCase.DeleteBlog(context.Background(), blogID, userID)

// 	suite.Error(err)
// 	suite.EqualError(err, "failed to delete blog: some repository error")
// 	suite.repoMock.AssertCalled(suite.T(), "GetBlogByID", mock.Anything, blogID)
// 	suite.repoMock.AssertCalled(suite.T(), "DeleteBlog", mock.Anything, blogID)
// }

// func (suite *BlogUseCaseTestSuite) TestFilterBlogs_Success() {
// 	filter := domain.BlogFilter{
// 		Title: ptrToString("Test Blog"),
// 	}
// 	expectedBlogs := []*domain.Blog{
// 		{
// 			ID:    primitive.NewObjectID(),
// 			Title: "Test Blog",
// 		},
// 	}

// 	suite.repoMock.On("FilterBlogs", suite.ctx, filter).Return(expectedBlogs, nil)

// 	blogs, err := suite.useCase.FilterBlogs(suite.ctx, filter)

// 	suite.NoError(err)
// 	suite.Equal(expectedBlogs, blogs)
// 	suite.repoMock.AssertExpectations(suite.T())
// }


// func (suite *BlogUseCaseTestSuite) TestFilterBlogs_Error() {
// 	filter := domain.BlogFilter{
// 		Title: ptrToString("Nonexistent Blog"),
// 	}
// 	expectedError := fmt.Errorf("some repository error")

// 	suite.repoMock.On("FilterBlogs", suite.ctx, filter).Return(nil, expectedError)

// 	blogs, err := suite.useCase.FilterBlogs(suite.ctx, filter)

// 	suite.Error(err)
// 	suite.Nil(blogs)
// 	suite.EqualError(err, "failed to filter blogs: some repository error")
// 	suite.repoMock.AssertExpectations(suite.T())
// }


// func ptrToString(s string) *string {
// 	return &s
// }

// func TestBlogUseCaseTestSuite(t *testing.T) {
// 	suite.Run(t, new(BlogUseCaseTestSuite))
// }
