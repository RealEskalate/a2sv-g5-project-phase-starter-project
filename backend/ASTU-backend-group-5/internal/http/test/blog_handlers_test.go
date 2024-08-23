package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"blogApp/internal/domain"
	"blogApp/internal/http/handlers/blog"
	"blogApp/mocks/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogHandlerTestSuite struct {
	suite.Suite
	handler   *blog.BlogHandler
	useCase   *mocks.BlogUseCase
	router    *gin.Engine
	ctx       context.Context
}

func (suite *BlogHandlerTestSuite) SetupTest() {
	suite.useCase = new(mocks.BlogUseCase)
	suite.handler = blog.NewBlogHandler(suite.useCase)
	suite.router = gin.Default()
	suite.ctx = context.TODO()
}

func (suite *BlogHandlerTestSuite) TestCreateBlogHandler_Success() {
	// Setup
	authorID := "abcdef1234567890abcdef12"
	authorObjectID, _ := primitive.ObjectIDFromHex(authorID)
	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: []interface{}{"Test content"},
		Tags:    []domain.BlogTag{},
		Author:  authorObjectID,
	}
	
	suite.useCase.On("CreateBlog", mock.Anything, blog, authorID).Return(nil)

	// Create request
	blogJSON, _ := json.Marshal(blog)
	req, _ := http.NewRequest("POST", "/blog", bytes.NewBuffer(blogJSON))
	req.Header.Set("Content-Type", "application/json")

	// Setup the router
	suite.router.POST("/blog", suite.handler.CreateBlogHandler)

	// Perform the request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assertions
	suite.Equal(http.StatusCreated, w.Code)
	suite.useCase.AssertCalled(suite.T(), "CreateBlog", mock.Anything, mock.Anything, authorID)
}

func (suite *BlogHandlerTestSuite) TestCreateBlogHandler_Failure_BindingError() {
	// Setup an invalid JSON
	invalidJSON := `{"title": "Test Blog", "content":`

	req, _ := http.NewRequest("POST", "/blog", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	// Setup the router
	suite.router.POST("/blog", suite.handler.CreateBlogHandler)

	// Perform the request
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assertions
	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *BlogHandlerTestSuite) TestCreateBlogHandler_Failure_UseCaseError() {
	authorID := "abcdef1234567890abcdef12"
	authorObjectID, _ := primitive.ObjectIDFromHex(authorID)
	blog := &domain.Blog{
		Title:   "Test Blog",
		Content: []interface{}{"Test content"},
		Tags:    []domain.BlogTag{},
		Author:  authorObjectID,
	}

	expectedError := fmt.Errorf("some error")
	suite.useCase.On("CreateBlog", mock.Anything, blog, authorID).Return(expectedError)

	
	blogJSON, _ := json.Marshal(blog)
	req, _ := http.NewRequest("POST", "/blog", bytes.NewBuffer(blogJSON))
	req.Header.Set("Content-Type", "application/json")

	
	suite.router.POST("/blog", suite.handler.CreateBlogHandler)

	
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.useCase.AssertCalled(suite.T(), "CreateBlog", mock.Anything, blog, authorID)
}



func TestBlogHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogHandlerTestSuite))
}
