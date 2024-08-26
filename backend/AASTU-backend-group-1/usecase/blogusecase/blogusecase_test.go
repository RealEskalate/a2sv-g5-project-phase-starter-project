package blogusecase_test

import (
	"blogs/config"
	"blogs/domain"
	"blogs/mocks"
	"blogs/usecase/blogusecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogUsecaseTestSuite struct {
	suite.Suite
	usecase *blogusecase.BlogUsecase
	repo    *mocks.BlogRepository
}

func (suite *BlogUsecaseTestSuite) SetupTest() {
	suite.repo = new(mocks.BlogRepository)
	suite.usecase = blogusecase.NewBlogUsecase(suite.repo)
}

func (suite *BlogUsecaseTestSuite) TestInsertBlog_Success() {
	testBlog := &domain.Blog{Title: "Test Blog", Author: "Test Author", Tags: []string{"tag1"}}
	expectedBlog := &domain.Blog{Title: "Test Blog", Author: "Test Author", Tags: []string{"tag1"}}

	// Setup mock expectations
	suite.repo.On("InsertBlog", testBlog).Return(expectedBlog, nil)

	// Call the method
	result, err := suite.usecase.InsertBlog(testBlog)

	// Assert results
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedBlog, result)

	// Verify mock expectations
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestInsertBlog_Failure() {	
	testBlog := &domain.Blog{Title: "Test Blog", Author: "Test Author", Tags: []string{"tag1"}}
	expectedError := assert.AnError // Use an actual error if assert.AnError is not available

	// Setup mock expectations
	suite.repo.On("InsertBlog", testBlog).Return(nil, expectedError)

	// Call the method
	result, err := suite.usecase.InsertBlog(testBlog)

	// Assert results
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), expectedError, err)

	// Verify mock expectations
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestAddComment_Success() {    
	comment := &domain.Comment{BlogID: primitive.NewObjectID(), Author: "Test Author", Content: "Test Comment"}
	
	// Setup mock expectations
	suite.repo.On("GetBlogByID", comment.BlogID.Hex()).Return(&domain.Blog{ID: comment.BlogID}, nil).Once()
	suite.repo.On("IncrmentBlogComments", comment.BlogID.Hex()).Return(nil).Once()
	suite.repo.On("AddComment", comment).Return(nil).Once()

	// Call the method
	err := suite.usecase.AddComment(comment)

	// Assert results
	assert.NoError(suite.T(), err)

	// Verify mock expectations
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestAddComment_BlogNotFound() {
	comment := &domain.Comment{BlogID: primitive.NewObjectID(), Author: "Test Author", Content: "Test Comment"}

	// Setup mock expectations
	suite.repo.On("GetBlogByID", comment.BlogID.Hex()).Return(nil, mongo.ErrNoDocuments).Once()

	// Call the method
	err := suite.usecase.AddComment(comment)

	// Assert results
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), config.ErrBlogNotFound, err)

	// Verify mock expectations
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestGetBlogComments_Success() {
	blogID := primitive.NewObjectID().Hex()
	comment := &domain.Comment{BlogID: primitive.NewObjectID(), Author: "Test Author", Content: "Test Comment"}

	// Setup mock expectations
	suite.repo.On("GetBlogByID", blogID).Return(&domain.Blog{ID: primitive.NewObjectID()}, nil).Once()
	suite.repo.On("GetBlogComments", blogID).Return([]*domain.Comment{comment}, nil).Once()

	// Call the method
	result, err := suite.usecase.GetBlogComments(blogID)

	// Assert results
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), []*domain.Comment{comment}, result)

	// Verify mock expectations
	suite.repo.AssertExpectations(suite.T())
}

func (suite *BlogUsecaseTestSuite) TestGetBlogComments_BlogNotFound() {
	blogID := primitive.NewObjectID().Hex()

	// Setup mock expectations
	suite.repo.On("GetBlogByID", blogID).Return(nil, mongo.ErrNoDocuments).Once()
	suite.repo.On("GetBlogComments", blogID).Return(nil, nil).Once()
	// Call the method
	result, err := suite.usecase.GetBlogComments(blogID)

	// Assert results
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), config.ErrBlogNotFound, err)

	// Verify mock expectations
	suite.repo.AssertExpectations(suite.T())
}

//Delete Comment
func (suite *BlogUsecaseTestSuite) TestDeleteComment_Success() {
	commentID := primitive.NewObjectID().Hex()
	BlogID := primitive.NewObjectID()

	claims := &domain.LoginClaims{
		Username: "Test Author",
		Role:     "user",
	}


	// Setup mock expectations
	suite.repo.On("DeleteComment", commentID).Return(nil).Once()
	suite.repo.On("GetCommentByID", commentID).Return(&domain.Comment{Author: "Test Author", BlogID: BlogID}, nil).Once()
	suite.repo.On("DecrementBlogComments", BlogID.Hex()).Return(nil).Once()

	// Call the method
	err := suite.usecase.DeleteComment(commentID,claims )


	// Assert results
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), nil, err)

	suite.repo.AssertExpectations(suite.T())
}



func (suite *BlogUsecaseTestSuite) TearDownTest() {
	suite.repo.AssertExpectations(suite.T())
}

func TestBlogUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogUsecaseTestSuite))
}



