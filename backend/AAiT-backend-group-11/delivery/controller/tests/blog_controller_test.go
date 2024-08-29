package controller_test

import (
	"bytes"
	"encoding/json"
	"backend-starter-project/delivery/controller"
	"backend-starter-project/domain/dto"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogControllerTestSuite struct {
	suite.Suite
	mockBlogService *MockBlogService
	blogController  *controller.BlogController
}

type MockBlogService struct {
	mock.Mock
}

func (m *MockBlogService) CreateBlogPost(blogPost *dto.AddBlogRequest, userId string) (*dto.AddBlogResponse, error) {
	args := m.Called(blogPost, userId)
	return args.Get(0).(*dto.AddBlogResponse), args.Error(1)
}

func (m *MockBlogService) GetBlogPostById(blogPostId string, userId string) (*dto.GetBlogByIDResponse, error) {
	args := m.Called(blogPostId, userId)
	return args.Get(0).(*dto.GetBlogByIDResponse), args.Error(1)
}

func (m *MockBlogService) GetBlogPosts(page int, pageSize int, sortBy string) (*dto.GetBlogPostsResponse, int, error) {
	args := m.Called(page, pageSize, sortBy)
	return args.Get(0).(*dto.GetBlogPostsResponse), args.Int(1), args.Error(2)
}

func (m *MockBlogService) UpdateBlogPost(blogPost *dto.UpdateBlogRequest, userId string) (*dto.UpdateBlogResponse, error) {
	args := m.Called(blogPost, userId)
	return args.Get(0).(*dto.UpdateBlogResponse), args.Error(1)
}

func (m *MockBlogService) DeleteBlogPost(blogPostId string, userId string, role string) error {
	args := m.Called(blogPostId, userId, role)
	return args.Error(0)
}

func (m *MockBlogService) SearchBlogPosts(searchText string) (*dto.GetBlogPostsResponse, error) {
	args := m.Called(searchText)
	return args.Get(0).(*dto.GetBlogPostsResponse), args.Error(1)
}

func (m *MockBlogService) FilterBlogPosts(filterReq dto.FilterBlogPostsRequest) (*dto.GetBlogPostsResponse, error) {
	args := m.Called(filterReq)
	return args.Get(0).(*dto.GetBlogPostsResponse), args.Error(1)
}

func (suite *BlogControllerTestSuite) SetupTest() {
	suite.mockBlogService = new(MockBlogService)
	suite.blogController = controller.NewBlogController(suite.mockBlogService)
}

func (suite *BlogControllerTestSuite) TestCreateBlogPost_Success() {
	blogPost := dto.AddBlogRequest{
		Title:   "Test Title",
		Content: "Test Content",
		Tags:    []string{"test", "golang"},
		Username: "TestUser",
	}
	userId := "123"
	createdBlogPost := &dto.AddBlogResponse{
		ID:             "1",
		AutherID:       "123",
		AutherUserName: "TestUser",
		Title:          "Test Title",
		Content:        "Test Content",
		Tags:           []string{"test", "golang"},
	}

	suite.mockBlogService.On("CreateBlogPost", &blogPost, userId).Return(createdBlogPost, nil)

	// Create a request with the JSON body
	body, _ := json.Marshal(blogPost)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userId", userId)
	c.Set("username", "TestUser")
	c.Request = httptest.NewRequest("POST", "/blog", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.blogController.CreateBlogPost(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Blog post created successfully")
	suite.mockBlogService.AssertCalled(suite.T(), "CreateBlogPost", &blogPost, userId)
}

func (suite *BlogControllerTestSuite) TestCreateBlogPost_Failure() {
	blogPost := dto.AddBlogRequest{
		Title:    "Test Title",
		Content:  "Test Content",
		Tags:     []string{"test", "golang"},
		Username: "TestUser",
	}
	userId := "123"

	// Mock the service to return nil as the first argument and an error as the second
	suite.mockBlogService.On("CreateBlogPost", &blogPost, userId).Return((*dto.AddBlogResponse)(nil), errors.New("error creating blog post"))

	body, _ := json.Marshal(blogPost)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userId", userId)
	c.Set("username", "TestUser")
	c.Request = httptest.NewRequest("POST", "/blog", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.blogController.CreateBlogPost(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "Error while creating user")
	suite.mockBlogService.AssertCalled(suite.T(), "CreateBlogPost", &blogPost, userId)
}


func (suite *BlogControllerTestSuite) TestGetBlogPost_Success() {
	blogPostId := "1"
	userId := "123"
	blogPost := &dto.GetBlogByIDResponse{
		ID:             "1",
		AutherID:       "123",
		AutherUserName: "TestUser",
		Title:          "Test Title",
		Content:        "Test Content",
	}

	suite.mockBlogService.On("GetBlogPostById", blogPostId, userId).Return(blogPost, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userId", userId)
	c.Params = gin.Params{{Key: "id", Value: blogPostId}}

	suite.blogController.GetBlogPost(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Test Title")
	suite.mockBlogService.AssertCalled(suite.T(), "GetBlogPostById", blogPostId, userId)
}

func (suite *BlogControllerTestSuite) TestGetBlogPost_Failure() {
	blogPostId := "1"
	userId := "123"

	suite.mockBlogService.On("GetBlogPostById", blogPostId, userId).Return(&dto.GetBlogByIDResponse{}, errors.New("error getting blog post"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userId", userId)
	c.Params = gin.Params{{Key: "id", Value: blogPostId}}

	suite.blogController.GetBlogPost(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "Error getting blog post")
	suite.mockBlogService.AssertCalled(suite.T(), "GetBlogPostById", blogPostId, userId)
}

func (suite *BlogControllerTestSuite) TestGetBlogPosts_Success() {
	page := 1
	pageSize := 10
	sortBy := "createdAt"
	totalPosts := 50
	blogPosts := &dto.GetBlogPostsResponse{
		BlogPosts: []interface{}{
			dto.GetBlogByIDResponse{
				ID:             "1",
				AutherID:       "123",
				AutherUserName: "TestUser",
				Title:          "Test Title 1",
				Content:        "Test Content 1",
			},
			dto.GetBlogByIDResponse{
				ID:             "2",
				AutherID:       "124",
				AutherUserName: "TestUser2",
				Title:          "Test Title 2",
				Content:        "Test Content 2",
			},
		},
		Pagination: dto.Pagination{
			CurrentPage: page,
			PageSize:    pageSize,
			TotalPages:  (totalPosts + pageSize - 1) / pageSize,
			TotalPosts:  totalPosts,
		},
	}

	suite.mockBlogService.On("GetBlogPosts", page, pageSize, sortBy).Return(blogPosts, totalPosts, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/blogs?page=1&pageSize=10&sortBy=createdAt", nil)

	suite.blogController.GetBlogPosts(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Test Title 1")
	suite.Contains(w.Body.String(), "Test Title 2")
	suite.mockBlogService.AssertCalled(suite.T(), "GetBlogPosts", page, pageSize, sortBy)
}

func (suite *BlogControllerTestSuite) TestGetBlogPosts_Failure() {
	page := 1
	pageSize := 10
	sortBy := "createdAt"

	suite.mockBlogService.On("GetBlogPosts", page, pageSize, sortBy).Return(&dto.GetBlogPostsResponse{}, 0, errors.New("error getting blog posts"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/blogs?page=1&pageSize=10&sortBy=createdAt", nil)

	suite.blogController.GetBlogPosts(c)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "Error while getting blog posts")
	suite.mockBlogService.AssertCalled(suite.T(), "GetBlogPosts", page, pageSize, sortBy)
}

func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}
