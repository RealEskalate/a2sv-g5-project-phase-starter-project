package controllers_test

import (
	"blog_g2/deliveries/controllers"
	"blog_g2/domain"
	"blog_g2/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CommentControllerTestSuite struct {
	suite.Suite
	controller   *controllers.CommentController
	mockUsecase  *mocks.CommentUsecase
	mockContext  *gin.Context
	mockResponse *httptest.ResponseRecorder
}

func (suite *CommentControllerTestSuite) SetupTest() {
	// Set up the mocks and the controller
	suite.mockUsecase = new(mocks.CommentUsecase)
	suite.controller = controllers.NewCommentController(suite.mockUsecase)
	suite.mockResponse = httptest.NewRecorder()
	suite.mockContext, _ = gin.CreateTestContext(suite.mockResponse)
}

// Test CreateComment
func (suite *CommentControllerTestSuite) TestCreateComment_Success() {
	blogID := "123"
	userID := "456"
	commentJSON := `{"comment": "This is a test comment"}`

	suite.mockContext.Params = gin.Params{gin.Param{Key: "blog_id", Value: blogID}}
	suite.mockContext.Request = httptest.NewRequest(http.MethodPost, "/comments", strings.NewReader(commentJSON))
	suite.mockContext.Request.Header.Set("Content-Type", "application/json")
	suite.mockContext.Set("userid", userID)

	suite.mockUsecase.On("CreateComment", mock.Anything, blogID, userID, mock.AnythingOfType("domain.Comment")).Return(nil)

	suite.controller.CreateComment(suite.mockContext)

	assert.Equal(suite.T(), http.StatusCreated, suite.mockResponse.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

// func (suite *CommentControllerTestSuite) TestCreateComment_BadRequest() {
// 	commentJSON := `{"invalid_json": "missing comment field"}`

// 	suite.mockContext.Request = httptest.NewRequest(http.MethodPost, "/comments", strings.NewReader(commentJSON))
// 	suite.mockContext.Request.Header.Set("Content-Type", "application/json")
// 	suite.controller.CreateComment(suite.mockContext)

// 	assert.Equal(suite.T(), http.StatusBadRequest, suite.mockResponse.Code)
// }

// Test GetComment
func (suite *CommentControllerTestSuite) TestGetComment_Success() {
	blogID := "123"
	expectedComments := []domain.Comment{
		{Comment: "This is a test comment 1"},
		{Comment: "This is a test comment 2"},
	}

	suite.mockContext.Params = gin.Params{gin.Param{Key: "blog_id", Value: blogID}}
	suite.mockUsecase.On("GetComments", mock.Anything, blogID).Return(expectedComments, nil)

	suite.controller.GetComment(suite.mockContext)

	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

// Test UpdateComment
func (suite *CommentControllerTestSuite) TestUpdateComment_Success() {
	commentID := "123"
	commentJSON := `{"comment": "Updated comment"}`

	suite.mockContext.Params = gin.Params{gin.Param{Key: "id", Value: commentID}}
	suite.mockContext.Request = httptest.NewRequest(http.MethodPut, "/comments/"+commentID, strings.NewReader(commentJSON))
	suite.mockContext.Request.Header.Set("Content-Type", "application/json")

	suite.mockUsecase.On("UpdateComment", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	suite.controller.UpdateComment(suite.mockContext)

	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

// Test DeleteComment
func (suite *CommentControllerTestSuite) TestDeleteComment_Success() {
	commentID := "123"

	suite.mockContext.Params = gin.Params{gin.Param{Key: "id", Value: commentID}}

	suite.mockUsecase.On("DeleteComment", mock.Anything, mock.Anything).Return(nil)

	suite.controller.DeleteComment(suite.mockContext)

	assert.Equal(suite.T(), http.StatusOK, suite.mockResponse.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

// Running the suite
func TestCommentControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CommentControllerTestSuite))
}
