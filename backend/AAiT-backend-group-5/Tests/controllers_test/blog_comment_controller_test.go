package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Test Suite
type BlogCommentControllerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.BlogCommentUsecase
	controller  interfaces.BlogCommentController
	router      *gin.Engine
}

func (suite *BlogCommentControllerTestSuite) SetupSuite() {
	suite.mockUsecase = new(mocks.BlogCommentUsecase)
	suite.controller = controllers.NewBlogCommentController(suite.mockUsecase)
	suite.router = gin.Default()

	// Define the routes
	suite.router.POST("/comment/:blogID", suite.controller.AddCommentController)
	suite.router.GET("/comments/:blogID", suite.controller.GetCommentsController)
	suite.router.GET("/comment/:commentID", suite.controller.GetCommentController)
	suite.router.PUT("/comment/:commentID", suite.controller.UpdateCommentController)
	suite.router.DELETE("/comment/:commentID", suite.controller.DeleteCommentController)
}

func (suite *BlogCommentControllerTestSuite) TearDownSuite() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *BlogCommentControllerTestSuite) TestAddCommentController_Success() {
	blogID := "blog123"
	userID := "user123"

	commentRequest := dtos.CommentCreateRequest{
		Content: "This is a test comment",
	}

	suite.mockUsecase.On("AddComment", mock.Anything, mock.Anything).Return(nil).Once()

	requestBody, _ := json.Marshal(commentRequest)
	request, _ := http.NewRequest(http.MethodPost, "/comment/"+blogID, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = request
	ctx.Set("id", userID)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: blogID})

	suite.controller.AddCommentController(ctx)

	suite.Equal(http.StatusCreated, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Comment added successfully")
}

func (suite *BlogCommentControllerTestSuite) TestGetCommentsController_Success() {
	blogID := "blog123"

	expectedComments := []models.Comment{
		{Content: "Comment 1", BlogID: blogID, UserID: "user1"},
		{Content: "Comment 2", BlogID: blogID, UserID: "user2"},
	}

	suite.mockUsecase.On("GetComments", mock.Anything, mock.Anything).Return(expectedComments, nil).Once()

	request, _ := http.NewRequest(http.MethodGet, "/comments/"+blogID, nil)
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Comment 1")
	suite.Contains(responseWriter.Body.String(), "Comment 2")
}

func (suite *BlogCommentControllerTestSuite) TestGetCommentController_Success() {
	commentID := "comment123"

	expectedComment := models.Comment{
		Content: "This is a test comment",
		BlogID:  "blog123",
		UserID:  "user123",
	}

	suite.mockUsecase.On("GetComment", mock.Anything, commentID).Return(&expectedComment, nil).Once()

	request, _ := http.NewRequest(http.MethodGet, "/comment/"+commentID, nil)
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "This is a test comment")
}

func (suite *BlogCommentControllerTestSuite) TestUpdateCommentController_Success() {
	commentID := "comment123"
	userID := "user123"

	commentUpdateRequest := dtos.CommentUpdateRequest{
		Content: "Updated comment content",
	}

	suite.mockUsecase.On("UpdateComment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	requestBody, _ := json.Marshal(commentUpdateRequest)
	request, _ := http.NewRequest(http.MethodPut, "/comment/"+commentID, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = request
	ctx.Set("id", userID)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: commentID})

	suite.controller.UpdateCommentController(ctx)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Comment updated successfully")
}

func (suite *BlogCommentControllerTestSuite) TestDeleteCommentController_Success() {
	commentID := "comment123"
	userID := "user123"

	suite.mockUsecase.On("DeleteComment", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	request, _ := http.NewRequest(http.MethodDelete, "/comment/"+commentID, nil)
	responseWriter := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = request
	ctx.Set("id", userID)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: commentID})

	suite.controller.DeleteCommentController(ctx)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Comment deleted successfully")
}

func TestBlogCommentControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogCommentControllerTestSuite))
}
