package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Test Suite
type BlogControllerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.BlogUsecase
	controller  *controllers.BlogController
	router      *gin.Engine
}

func (suite *BlogControllerTestSuite) SetupSuite() {
	suite.mockUsecase = new(mocks.BlogUsecase)
	suite.controller = controllers.NewBlogController(suite.mockUsecase)
	suite.router = gin.Default()

	// Define the routes
	suite.router.POST("/blogs", suite.controller.CreateBlogController)
	suite.router.GET("/blogs/:id", suite.controller.GetBlogController)
	suite.router.GET("/blogs", suite.controller.GetBlogsController)
	suite.router.GET("/blogs/search", suite.controller.SearchBlogsController)
	suite.router.PUT("/blogs/:id", suite.controller.UpdateBlogController)
	suite.router.DELETE("/blogs/:id", suite.controller.DeleteBlogController)
	suite.router.POST("/blogs/:id/popularity", suite.controller.TrackPopularityController)
}

func (suite *BlogControllerTestSuite) TearDownSuite() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestCreateBlogController_Success() {
	authorID := "author123"

	// The request data that is sent in the HTTP POST request
	blogRequest := &dtos.CreateBlogRequest{
		Title:   "Test Title",
		Content: "Test Content",
		Tags:    []string{"tag1", "tag2"},
	}

	// The expected Blog model that the mock Usecase should receive
	expectedBlog := models.Blog{
		Title:    blogRequest.Title,
		Content:  blogRequest.Content,
		Tags:     blogRequest.Tags,
		AuthorID: authorID,
	}

	mockBlog := expectedBlog

	// The response we expect to be returned by the Usecase
	expectedResponse := dtos.BlogResponse{
		Blog: expectedBlog,
	}

	// Set up the mock to expect specific parameters and return the expected response
	suite.mockUsecase.On("CreateBlog", mock.AnythingOfType("*gin.Context"), &mockBlog).Return(&expectedResponse, nil).Once()

	// Marshal the request struct to JSON
	requestBody, _ := json.Marshal(expectedBlog)
	request, _ := http.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = request
	ctx.Set("id", authorID)

	suite.controller.CreateBlogController(ctx)

	// Assertions to check the response
	suite.Equal(http.StatusCreated, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Test Title")
	suite.Contains(responseWriter.Body.String(), "Test Content")
}

func (suite *BlogControllerTestSuite) TestCreateBlogController_BadRequest() {
	request, _ := http.NewRequest(http.MethodPost, "/blogs", nil)
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusBadRequest, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Invalid request payload")
}

func (suite *BlogControllerTestSuite) TestGetBlogController_Success() {
	blogID := "blog123"

	expectedBlog := &models.Blog{
		ID:      blogID,
		Title:   "Test Title",
		Content: "Test Content",
	}

	expectedResponse := dtos.BlogResponse{
		Blog: *expectedBlog,
	}

	suite.mockUsecase.On("GetBlog", mock.Anything, blogID).Return(&expectedResponse, nil).Once()

	request, _ := http.NewRequest(http.MethodGet, "/blogs/"+blogID, nil)
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Test Title")
	suite.Contains(responseWriter.Body.String(), "Test Content")
}

func (suite *BlogControllerTestSuite) TestGetBlogController_NotFound() {
	blogID := "nonexistent"
	expectedError := &models.ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "Blog not found",
	}

	suite.mockUsecase.On("GetBlog", mock.Anything, blogID).Return(nil, expectedError).Once()

	request, _ := http.NewRequest(http.MethodGet, "/blogs/"+blogID, nil)
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(expectedError.Code, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Blog not found")
}

func (suite *BlogControllerTestSuite) TestGetBlogsController_Success() {
	page := "1"
	expectedBlogs := []*dtos.BlogResponse{
		{Blog: models.Blog{ID: "blog1", Title: "Test Title 1", Content: "Test Content 1"}},
		{Blog: models.Blog{ID: "blog2", Title: "Test Title 2", Content: "Test Content 2"}},
	}

	suite.mockUsecase.On("GetBlogs", mock.Anything, 1).Return(expectedBlogs, nil).Once()

	request, _ := http.NewRequest(http.MethodGet, "/blogs?page="+page, nil)
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Test Title 1")
	suite.Contains(responseWriter.Body.String(), "Test Title 2")
}

func (suite *BlogControllerTestSuite) TestGetBlogsController_BadRequest() {
	request, _ := http.NewRequest(http.MethodGet, "/blogs?page=invalid", nil)
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusBadRequest, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Invalid page number")
}

func (suite *BlogControllerTestSuite) TestSearchBlogsController_Success() {
	expectedBlogs := []*dtos.BlogResponse{
		{Blog: models.Blog{ID: "blog1", Title: "Test_Title_1", Content: "Test Content 1"}},
		{Blog: models.Blog{ID: "blog2", Title: "Test_Title_2", Content: "Test Content 2"}},
	}

	suite.mockUsecase.On("SearchBlogs", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("dtos.FilterBlogRequest")).Return(expectedBlogs, nil).Once()

	request, _ := http.NewRequest(http.MethodGet, "/blogs/search?title=Test_Title_1", nil)
	responseWriter := httptest.NewRecorder()

	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Test_Title_1")
}

func (suite *BlogControllerTestSuite) TestUpdateBlogController_Success() {
	blogID := "blog123"

	blogRequest := dtos.UpdateBlogRequest{
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"tag1", "tag2"},
	}

	suite.mockUsecase.On("UpdateBlog", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	requestBody, _ := json.Marshal(blogRequest)
	request, _ := http.NewRequest(http.MethodPut, "/blogs/"+blogID, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()
	suite.router.ServeHTTP(responseWriter, request)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Blog updated successfully")
}

func (suite *BlogControllerTestSuite) TestDeleteBlogController_Success() {
	blogID := "blog123"
	authorID := "author123"

	deleteBlogReq := dtos.DeleteBlogRequest{
		BlogID:   blogID,
		AuthorID: authorID,
	}

	suite.mockUsecase.On("DeleteBlog", mock.Anything, deleteBlogReq).Return(nil).Once()
	request, _ := http.NewRequest(http.MethodDelete, "/blogs/"+blogID, nil)
	responseWriter := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = request
	ctx.Set("id", authorID)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: blogID})

	suite.controller.DeleteBlogController(ctx)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Blog deleted successfully")
}

func (suite *BlogControllerTestSuite) TestTrackPopularityController_Success() {
	blogID := "blog123"
	authorID := "author123"

	trackPopularityReq := dtos.TrackPopularityRequest{
		UserID: authorID,
		BlogID: blogID,
	}

	suite.mockUsecase.On("TrackPopularity", mock.Anything, trackPopularityReq).Return(nil).Once()

	requestBody, _ := json.Marshal(trackPopularityReq)
	request, _ := http.NewRequest(http.MethodPost, "/blogs/"+blogID+"/popularity", bytes.NewBuffer(requestBody))
	responseWriter := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = request
	ctx.Set("id", authorID)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: blogID})

	suite.controller.TrackPopularityController(ctx)

	suite.Equal(http.StatusOK, responseWriter.Code)
	suite.Contains(responseWriter.Body.String(), "Popularity tracked successfully")
}

func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}
