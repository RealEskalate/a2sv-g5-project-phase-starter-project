package controller

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/domain/dto"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PopularityTrackingControllerTestSuite struct {
	suite.Suite
	controller          *controller.PopularityTrackingController
	mockService         *MockPopularityTrackingService
}

type MockPopularityTrackingService struct {
	mock.Mock
}

func (m *MockPopularityTrackingService) IncrementViewCount(blogPostId string) error {
	args := m.Called(blogPostId)
	return args.Error(0)
}

func (m *MockPopularityTrackingService) LikeBlogPost(blogPostId string, userId string) error {
	args := m.Called(blogPostId, userId)
	return args.Error(0)
}

func (m *MockPopularityTrackingService) DislikeBlogPost(blogPostId string, userId string) error {
	args := m.Called(blogPostId, userId)
	return args.Error(0)
}

func (m *MockPopularityTrackingService) GetPopularityMetrics(blogPostId string) (map[string]int, error) {
	args := m.Called(blogPostId)
	return args.Get(0).(map[string]int), args.Error(1)
}

func (suite *PopularityTrackingControllerTestSuite) SetupTest() {
	suite.mockService = new(MockPopularityTrackingService)
	suite.controller = controller.NewPopularityTrackingController(suite.mockService)
}

func (suite *PopularityTrackingControllerTestSuite) TestLikeBlogPostSuccess() {
	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Set("userId", "user123")

	suite.mockService.On("LikeBlogPost", "1", "user123").Return(nil)

	// Act
	suite.controller.LikeBlogPost(c)

	// Assert
	suite.Equal(http.StatusOK, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.True(response.Success)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *PopularityTrackingControllerTestSuite) TestLikeBlogPostUserNotFound() {
	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	// No userId is set in the context

	// Act
	suite.controller.LikeBlogPost(c)

	// Assert
	suite.Equal(http.StatusInternalServerError, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.False(response.Success)
	suite.Equal("User not found", response.Error)
}

func (suite *PopularityTrackingControllerTestSuite) TestLikeBlogPostError() {
	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Set("userId", "user123")

	suite.mockService.On("LikeBlogPost", "1", "user123").Return(errors.New("some error"))

	// Act
	suite.controller.LikeBlogPost(c)

	// Assert
	suite.Equal(http.StatusInternalServerError, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.False(response.Success)
	suite.Equal("some error", response.Error)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *PopularityTrackingControllerTestSuite) TestDislikeBlogPostSuccess() {
	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Set("userId", "user123")

	suite.mockService.On("DislikeBlogPost", "1", "user123").Return(nil)

	// Act
	suite.controller.DislikeBlogPost(c)

	// Assert
	suite.Equal(http.StatusOK, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.True(response.Success)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *PopularityTrackingControllerTestSuite) TestDislikeBlogPostUserNotFound() {
	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	// No userId is set in the context

	// Act
	suite.controller.DislikeBlogPost(c)

	// Assert
	suite.Equal(http.StatusInternalServerError, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.False(response.Success)
	suite.Equal("User not found", response.Error)
}

func (suite *PopularityTrackingControllerTestSuite) TestDislikeBlogPostError() {
	// Arrange
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Set("userId", "user123")

	suite.mockService.On("DislikeBlogPost", "1", "user123").Return(errors.New("some error"))

	// Act
	suite.controller.DislikeBlogPost(c)

	// Assert
	suite.Equal(http.StatusInternalServerError, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.False(response.Success)
	suite.Equal("some error", response.Error)
	suite.mockService.AssertExpectations(suite.T())
}

func TestPopularityTrackingControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PopularityTrackingControllerTestSuite))
}
