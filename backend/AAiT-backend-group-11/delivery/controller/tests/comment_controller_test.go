package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend-starter-project/delivery/controller"
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockCommentService is a mock implementation of the CommentService interface
type MockCommentService struct {
	mock.Mock
}

func (m *MockCommentService) AddComment(comment *entities.Comment) (*entities.Comment, error) {
	args := m.Called(comment)
	return args.Get(0).(*entities.Comment), args.Error(1)
}

func (m *MockCommentService) DeleteComment(commentId string) error {
	args := m.Called(commentId)
	return args.Error(0)
}

func (m *MockCommentService) GetCommentsByBlogPostId(blogPostId string) ([]entities.Comment, error) {
	args := m.Called(blogPostId)
	return args.Get(0).([]entities.Comment), args.Error(1)
}

func (m *MockCommentService) UpdateComment(comment *entities.Comment) (*entities.Comment, error) {
	args := m.Called(comment)
	return args.Get(0).(*entities.Comment), args.Error(1)
}

// CommentControllerSuite is the test suite for the CommentController
type CommentControllerSuite struct {
	suite.Suite
	mockService *MockCommentService
	cc          *controller.CommentController
	router      *gin.Engine
}

func (suite *CommentControllerSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.mockService = new(MockCommentService)
	suite.cc = controller.NewCommentController(suite.mockService)
	suite.router = gin.Default()
}

func (suite *CommentControllerSuite) SetupTest() {
	// Common setup tasks before each test can go here
}

func (suite *CommentControllerSuite) TestAddComment() {
	suite.router.POST("/comments/:blogId", suite.cc.AddComment)

	comment := entities.Comment{
		ID:         primitive.NewObjectID(),
		Content:    "Test comment",
		BlogPostID: primitive.NewObjectID(),
		AuthorID:   primitive.NewObjectID(),
	}

	suite.mockService.On("AddComment", mock.AnythingOfType("*entities.Comment")).Return(&comment, nil)

	body, _ := json.Marshal(comment)
	req, _ := http.NewRequest(http.MethodPost, "/comments/"+comment.BlogPostID.Hex(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("userId", comment.AuthorID.Hex())

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.True(response.Success)
	suite.Equal("Comment added successfully", response.Message)
	suite.NotNil(response.Data)
}

func (suite *CommentControllerSuite) TestDeleteComment() {
	suite.router.DELETE("/comments/:id", suite.cc.DeleteComment)

	commentID := primitive.NewObjectID().Hex()

	suite.mockService.On("DeleteComment", commentID).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/comments/"+commentID, nil)

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.True(response.Success)
	suite.Equal("Comment deleted successfully", response.Message)
}

func (suite *CommentControllerSuite) TestGetCommentsByBlogPostId() {
	suite.router.GET("/comments/:blogId", suite.cc.GetCommentsByBlogPostId)

	blogPostID := primitive.NewObjectID().Hex()
	comments := []entities.Comment{
		{
			ID:         primitive.NewObjectID(),
			Content:    "Test comment",
			BlogPostID: primitive.NewObjectID(),
			AuthorID:   primitive.NewObjectID(),
		},
	}

	suite.mockService.On("GetCommentsByBlogPostId", blogPostID).Return(comments, nil)

	req, _ := http.NewRequest(http.MethodGet, "/comments/"+blogPostID, nil)

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.True(response.Success)
	suite.Equal("Comments retrieved successfully", response.Message)
	suite.NotNil(response.Data)
}

func (suite *CommentControllerSuite) TestUpdateComment() {
	suite.router.PUT("/comments", suite.cc.UpdateComment)

	comment := entities.Comment{
		ID:         primitive.NewObjectID(),
		Content:    "Test comment",
		BlogPostID: primitive.NewObjectID(),
		AuthorID:   primitive.NewObjectID(),
	}

	suite.mockService.On("UpdateComment", mock.AnythingOfType("*entities.Comment")).Return(&comment, nil)

	body, _ := json.Marshal(comment)
	req, _ := http.NewRequest(http.MethodPut, "/comments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("userId", comment.AuthorID.Hex())

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.True(response.Success)
	suite.Equal("Comment updated successfully", response.Message)
	suite.NotNil(response.Data)
}

func TestCommentControllerSuite(t *testing.T) {
	suite.Run(t, new(CommentControllerSuite))
}
