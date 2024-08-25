package controller

import (
	"Blog_Starter/domain"
	"Blog_Starter/domain/mocks"
	"Blog_Starter/utils"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define the test suite
type BlogControllerTestSuite struct {
	suite.Suite
	controller  *BlogController
	mockUseCase *mocks.BlogUseCase
	router      *gin.Engine
}

// Setup the test suite
func (suite *BlogControllerTestSuite) SetupTest() {
	suite.mockUseCase = new(mocks.BlogUseCase)
	suite.controller = NewBlogController(suite.mockUseCase)
	monkey.Patch(utils.CheckUser, func(c *gin.Context) (*domain.AuthenticatedUser, error) {
		return &domain.AuthenticatedUser{UserID: "mockedUserID"}, nil
	})
	suite.router = gin.Default()

	// Define routes
	suite.router.POST("/blogs", suite.controller.CreateBlog)
	suite.router.GET("/blogs/:id", suite.controller.GetBlogByID)
	suite.router.GET("/blogs", suite.controller.GetAllBlog)
	suite.router.PUT("/blogs/:id", suite.controller.UpdateBlog)
	suite.router.DELETE("/blogs/:id", suite.controller.DeleteBlog)
	suite.router.POST("/blogs/filter", suite.controller.FilterBlog)
	suite.router.GET("/blogs/search", suite.controller.SearchBlog)
}

// Test CreateBlog
func (suite *BlogControllerTestSuite) TestCreateBlog_Success() {
	blogCreate := &domain.BlogCreate{
		UserID:  "testUserID",
		Title:   "Test Title",
		Content: "Test Content",
		Tags:    []string{"tag1", "tag2"},
	}

	blog := &domain.Blog{
		BlogID: primitive.NewObjectID(),
		Title:  blogCreate.Title,
		Content: blogCreate.Content,
		Tags:    blogCreate.Tags,
		Author:  "Test Author",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockUseCase.On("CreateBlog", mock.Anything, blogCreate).Return(blog, nil)

	// Create request
	body, _ := json.Marshal(blogCreate)
	req, _ := http.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	// Assertions
	suite.Equal(http.StatusCreated, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestCreateBlog_InvalidContentLength() {
	blogCreate := &domain.BlogCreate{
		UserID:  "testUserID",
		Title:   "Test Title",
		Content: "Short",
		Tags:    []string{"tag1", "tag2"},
	}

	suite.mockUseCase.On("CreateBlog", mock.Anything, blogCreate).Return(nil, errors.New("content length should be greater than 10"))

	body, _ := json.Marshal(blogCreate)
	req, _ := http.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}


func (suite *BlogControllerTestSuite) TestGetBlogByID_Success() {
	blogID := primitive.NewObjectID()
	blog := &domain.Blog{
		BlogID:    blogID,
		Title:     "Test Blog",
		Content:   "Test Content",
		Author:    "Test Author",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockUseCase.On("GetBlogByID", mock.Anything, blogID.Hex()).Return(blog, nil)

	req, _ := http.NewRequest(http.MethodGet, "/blogs/" + blogID.Hex(), nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGetBlogByID_BlogNotFound() {
	blogID := primitive.NewObjectID().Hex()

	suite.mockUseCase.On("GetBlogByID", mock.Anything, blogID).Return(nil, errors.New("blog not found"))

	req, _ := http.NewRequest(http.MethodGet, "/blogs/"+blogID, nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGetBlogByID_InvalidBlogID() {
	invalidBlogID := "invalidID"

	suite.mockUseCase.On("GetBlogByID", mock.Anything, invalidBlogID).Return(nil, errors.New("invalid blog id"))

	req, _ := http.NewRequest(http.MethodGet, "/blogs/"+invalidBlogID, nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

// Test GetAllBlog
func (suite *BlogControllerTestSuite) TestGetAllBlog_Success() {
	blogs := []*domain.Blog{
		{
			BlogID:    primitive.NewObjectID(),
			Title:     "Test Blog 1",
			Content:   "Test Content 1",
			Author:    "Author 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			BlogID:    primitive.NewObjectID(),
			Title:     "Test Blog 2",
			Content:   "Test Content 2",
			Author:    "Author 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	paginationMetadata := &domain.PaginationMetadata{
		TotalRecords: 2,
		TotalPages:   1,
		CurrPage:     1,
	}

	suite.mockUseCase.On("GetAllBlog", mock.Anything, int64(0), int64(10), "").Return(blogs, paginationMetadata, nil)

	req, _ := http.NewRequest(http.MethodGet, "/blogs?skip=0&limit=10", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGetAllBlog_NoBlogs() {
	
	suite.mockUseCase.On("GetAllBlog", mock.Anything, int64(0), int64(10), "").Return(&domain.Blog{}, &domain.PaginationMetadata{}, errors.New("no blogs found"))

	req, _ := http.NewRequest(http.MethodGet, "/blogs?skip=0&limit=10", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

// Test UpdateBlog
func (suite *BlogControllerTestSuite) TestUpdateBlog_Success() {
	blogID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID()
	blogUpdate := &domain.BlogUpdate{
		UserID : "mockedUserID",
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"tag1", "tag2"},
	}

	updatedBlog := &domain.Blog{
		
		BlogID:    primitive.NewObjectID(),
		UserID:    userID,
		Title:     blogUpdate.Title,
		Content:   blogUpdate.Content,
		Tags:      blogUpdate.Tags,
		Author:    "Test Author",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockUseCase.On("UpdateBlog", mock.Anything, blogUpdate, blogID).Return(updatedBlog, nil)

	body, _ := json.Marshal(blogUpdate)
	req, _ := http.NewRequest(http.MethodPut, "/blogs/"+blogID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestUpdateBlog_BlogNotFound() {
	blogID := primitive.NewObjectID().Hex()
	blogUpdate := &domain.BlogUpdate{
		UserID : "mockedUserID",
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"tag1", "tag2"},
	}

	suite.mockUseCase.On("UpdateBlog", mock.Anything, blogUpdate, blogID).Return(nil, errors.New("blog not found"))

	body, _ := json.Marshal(blogUpdate)
	req, _ := http.NewRequest(http.MethodPut, "/blogs/"+blogID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

// Test DeleteBlog
func (suite *BlogControllerTestSuite) TestDeleteBlog_Success() {
	blogID := primitive.NewObjectID().Hex()

	suite.mockUseCase.On("DeleteBlog", mock.Anything, blogID, "mockedUserID","").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/blogs/"+blogID, nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNoContent, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestDeleteBlog_BlogNotFound() {
	blogID := primitive.NewObjectID().Hex()

	suite.mockUseCase.On("DeleteBlog", mock.Anything, blogID, "mockedUserID", "").Return(errors.New("blog not found"))

	req, _ := http.NewRequest(http.MethodDelete, "/blogs/"+blogID, nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

// Test FilterBlog
func (suite *BlogControllerTestSuite) TestFilterBlog_Success() {
	filter := &domain.BlogFilterRequest{
		Tags: []string{"tag1", "tag2"},
	}

	blogs := []*domain.Blog{
		{
			BlogID:    primitive.NewObjectID(),
			Title:     "Test Blog 1",
			Content:   "Test Content 1",
			Tags:      []string{"tag1", "tag2"},
			Author:    "Author 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			BlogID:    primitive.NewObjectID(),
			Title:     "Test Blog 2",
			Content:   "Test Content 2",
			Tags:      []string{"tag1", "tag2"},
			Author:    "Author 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	suite.mockUseCase.On("FilterBlogs", mock.Anything, filter).Return(blogs, nil)

	body, _ := json.Marshal(filter)
	req, _ := http.NewRequest(http.MethodPost, "/blogs/filter", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestFilterBlog_NoBlogsFound() {
	filter := &domain.BlogFilterRequest{
		Tags: []string{"nonexistentTag"},
	}

	suite.mockUseCase.On("FilterBlogs", mock.Anything, filter).Return(nil, errors.New("no matches found"))

	body, _ := json.Marshal(filter)
	req, _ := http.NewRequest(http.MethodPost, "/blogs/filter", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

// Test SearchBlog
func (suite *BlogControllerTestSuite) TestSearchBlog_Success() {
	searchRequest := &domain.BlogSearchRequest{
		Author: "Kidus Melaku",
		Title : "This is me",
	}
	blogs := []*domain.Blog{
		{
			BlogID:    primitive.NewObjectID(),
			Title:     "Test Blog 1",
			Content:   "Test Content 1",
			Tags:      []string{"tag1", "tag2"},
			Author:    "Author 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			BlogID:    primitive.NewObjectID(),
			Title:     "Test Blog 2",
			Content:   "Test Content 2",
			Tags:      []string{"tag1", "tag2"},
			Author:    "Author 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	suite.mockUseCase.On("SearchBlogs", mock.Anything, searchRequest).Return(blogs, nil)

	req, _ := http.NewRequest(http.MethodGet, "/blogs/search?author=Kidus%20Melaku&title=This%20is%20me" , nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestSearchBlog_NoBlogsFound() {
	searchRequest := &domain.BlogSearchRequest{
		Author: "Kidus Melaku",
		Title : "This is me",
	}

	suite.mockUseCase.On("SearchBlogs", mock.Anything, searchRequest).Return(nil, mongo.ErrNoDocuments)

	req, _ := http.NewRequest(http.MethodGet, "/blogs/search?author=Kidus%20Melaku&title=This%20is%20me", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

// Run the test suite
func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}
