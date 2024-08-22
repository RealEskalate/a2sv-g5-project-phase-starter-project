package test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentControllerTestSuite struct {
	suite.Suite
	mockCommentUsecase *mocks.Comment_Usecase_interface
	commentController  *controller.CommentController
}

func (suite *CommentControllerTestSuite) SetupTest() {
	suite.mockCommentUsecase = mocks.NewComment_Usecase_interface(suite.T())
	suite.commentController = &controller.CommentController{
		CommentUsecase: suite.mockCommentUsecase,
	}
}

// Test case for successfully deleting a comment
func (suite *CommentControllerTestSuite) TestDeleteComment_Success() {
	commentID := primitive.NewObjectID().Hex()
	suite.mockCommentUsecase.On("DeleteComment", commentID).Return(nil)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "comment_id", Value: commentID}}

	suite.commentController.DeleteComment()(c)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("Comment deleted successfully!", response["message"])

	suite.mockCommentUsecase.AssertExpectations(suite.T())
}

// Test case for invalid comment ID format
func (suite *CommentControllerTestSuite) TestDeleteComment_InvalidID() {
	invalidID := "invalid-id"

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "comment_id", Value: invalidID}}

	suite.commentController.DeleteComment()(c)

	suite.Equal(http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Contains(response["error"], "Invalid comment ID format. Please provide a valid comment ID.")

	suite.mockCommentUsecase.AssertNotCalled(suite.T(), "DeleteComment")
}

// Test case for error from the use case during comment deletion
func (suite *CommentControllerTestSuite) TestDeleteComment_UsecaseError() {
	commentID := primitive.NewObjectID().Hex()
	expectedError := errors.New("database error")
	suite.mockCommentUsecase.On("DeleteComment", commentID).Return(expectedError)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "comment_id", Value: commentID}}

	suite.commentController.DeleteComment()(c)

	suite.Equal(http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Contains(response["error"], "Failed to delete comment")
	suite.Contains(response["error"], expectedError.Error())

	suite.mockCommentUsecase.AssertExpectations(suite.T())
}

// Run the test suite
func TestCommentControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CommentControllerTestSuite))
}
