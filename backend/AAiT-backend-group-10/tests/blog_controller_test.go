package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/domain"
	"aait.backend.g10/tests/mocks"
	"aait.backend.g10/usecases/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogControllerSuite struct {
	suite.Suite
	blogController *controllers.BlogController
	mockUseCase    *mocks.IBlogUseCase
}

func (suite *BlogControllerSuite) SetupTest() {
	suite.mockUseCase = new(mocks.IBlogUseCase)
	suite.blogController = controllers.NewBlogController(suite.mockUseCase)
}

func (suite *BlogControllerSuite) TestCreateBlog_Success() {
	// Create mock data
	blogDto := dto.BlogDto{
		Title:   "Test Title",
		Content: "Test Content",
		Tags:    []string{"tag1", "tag2"},
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("CreateBlog", mock.AnythingOfType("*domain.Blog")).Return(&blogDto, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/blogs", bytes.NewBufferString(`{
		"title": "Test Title",
		"content": "Test Content",
		"tags": ["tag1", "tag2"]
	}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Simulate setting the author ID in context
	c.Set("id", uuid.New().String())

	// Call the actual function
	suite.blogController.CreateBlog(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestCreateBlog_InvalidJson() {
	// Simulate the HTTP request with invalid JSON
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/blogs", bytes.NewBufferString(`{}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the actual function
	suite.blogController.CreateBlog(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "errors")
	suite.mockUseCase.AssertNotCalled(suite.T(), "CreateBlog")
}

func (suite *BlogControllerSuite) TestCreateBlog_InvalidId() {
	// Simulate the HTTP request with invalid ID
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/blogs", bytes.NewBufferString(`{
		"title": "Test Title",
		"content": "Test Content",
		"tags": ["tag1", "tag2"]
	}`))
	c.Request.Header.Set("Content-Type", "application/json")

	c.Set("id", "invalid-id")

	// Call the actual function
	suite.blogController.CreateBlog(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid ID")
	suite.mockUseCase.AssertNotCalled(suite.T(), "CreateBlog")
}

func (suite *BlogControllerSuite) TestGetAllBlogs_Success() {
	// Create mock data
	blogs := []*dto.BlogDto{
		{
			Title:   "Test Title",
			Content: "Test Content",
			Tags:    []string{"tag1", "tag2"},
		},
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("GetAllBlogs").Return(blogs, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the actual function
	suite.blogController.GetAllBlogs(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestGetBlogByID_Success() {
	// Create mock data
	blog := &dto.BlogDto{
		Title:   "Test Title",
		Content: "Test Content",
		Tags:    []string{"tag1", "tag2"},
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("GetBlogByID", mock.AnythingOfType("uuid.UUID")).Return(blog, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: uuid.New().String()})

	// Call the actual function
	suite.blogController.GetBlogByID(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestGetBlogByID_InvalidId() {
	// Simulate the HTTP request with invalid ID
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "invalid-id"})

	// Call the actual function
	suite.blogController.GetBlogByID(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid ID")
	suite.mockUseCase.AssertNotCalled(suite.T(), "GetBlogByID")
}

func (suite *BlogControllerSuite) TestGetBlogByID_BlogNotFound() {
	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("GetBlogByID", mock.AnythingOfType("uuid.UUID")).Return(nil, domain.ErrBlogNotFound)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: uuid.New().String()})

	// Call the actual function
	suite.blogController.GetBlogByID(c)

	// Assertions
	assert.Equal(suite.T(), domain.ErrBlogNotFound.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), domain.ErrBlogNotFound.Error())
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestUpdateBlog_Success() {
	// Create mock data
	blogID := uuid.New()
	requesterID := uuid.New()
	updatedBlog := domain.Blog{
		ID:      blogID,
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"updated-tag"},
		Author:  requesterID,
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("UpdateBlog", &updatedBlog).Return(nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/blogs/"+blogID.String(), bytes.NewBufferString(`{
		"title": "Updated Title",
		"content": "Updated Content",
		"tags": ["updated-tag"]
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: blogID.String()}}
	c.Set("id", requesterID.String())

	// Call the actual function
	suite.blogController.UpdateBlog(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Blog updated successfully")
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestUpdateBlog_InvalidID() {
	// Simulate the HTTP request with an invalid UUID
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/blogs/invalid-id", bytes.NewBufferString(`{}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: "invalid-id"}}

	// Call the actual function
	suite.blogController.UpdateBlog(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid ID")
	suite.mockUseCase.AssertNotCalled(suite.T(), "UpdateBlog")
}

func (suite *BlogControllerSuite) TestUpdateBlog_InvalidJson() {
	// Simulate the HTTP request with invalid JSON
	blogID := uuid.New()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/blogs/"+blogID.String(), bytes.NewBufferString(`{}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: blogID.String()}}
	c.Set("id", uuid.New().String())

	// Call the actual function
	suite.blogController.UpdateBlog(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "errors")
	suite.mockUseCase.AssertNotCalled(suite.T(), "UpdateBlog")
}

func (suite *BlogControllerSuite) TestUpdateBlog_BlogNotFound() {
	// Create mock data
	blogID := uuid.New()
	requesterID := uuid.New()
	updatedBlog := domain.Blog{
		ID:      blogID,
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"updated-tag"},
		Author:  requesterID,
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("UpdateBlog", &updatedBlog).Return(domain.ErrBlogNotFound)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/blogs/"+blogID.String(), bytes.NewBufferString(`{
		"title": "Updated Title",
		"content": "Updated Content",
		"tags": ["updated-tag"]
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: blogID.String()}}
	c.Set("id", requesterID.String())

	// Call the actual function
	suite.blogController.UpdateBlog(c)

	// Assertions
	assert.Equal(suite.T(), domain.ErrBlogNotFound.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), domain.ErrBlogNotFound.Error())
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestUpdateBlog_Unauthorized() {
	// Create mock data
	blogID := uuid.New()
	requesterID := uuid.New()
	updatedBlog := domain.Blog{
		ID:      blogID,
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"updated-tag"},
		Author:  requesterID,
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("UpdateBlog", &updatedBlog).Return(domain.ErrUnAuthorized)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/blogs/"+blogID.String(), bytes.NewBufferString(`{
		"title": "Updated Title",
		"content": "Updated Content",
		"tags": ["updated-tag"]
	}`))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: blogID.String()}}
	c.Set("id", requesterID.String())

	// Call the actual function
	suite.blogController.UpdateBlog(c)

	// Assertions
	assert.Equal(suite.T(), domain.ErrUnAuthorized.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), domain.ErrUnAuthorized.Error())
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestDeleteBlog_Success() {
	// Create mock data
	blogID := uuid.New()
	requesterID := uuid.New()

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("DeleteBlog", blogID, requesterID, false).Return(nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: blogID.String()}}
	c.Set("id", requesterID.String())
	c.Set("is_admin", false)

	// Call the actual function
	suite.blogController.DeleteBlog(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Blog deleted successfully")
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestDeleteBlog_InvalidID() {
	// Simulate the HTTP request with an invalid UUID
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "invalid-id"}}

	// Call the actual function
	suite.blogController.DeleteBlog(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid ID")
	suite.mockUseCase.AssertNotCalled(suite.T(), "DeleteBlog")
}

func (suite *BlogControllerSuite) TestDeleteBlog_Unauthorized() {
	// Create mock data
	blogID := uuid.New()
	requesterID := uuid.New()

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("DeleteBlog", blogID, requesterID, false).Return(domain.ErrUnAuthorized)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: blogID.String()}}
	c.Set("id", requesterID.String())
	c.Set("is_admin", false)

	// Call the actual function
	suite.blogController.DeleteBlog(c)

	// Assertions
	assert.Equal(suite.T(), domain.ErrUnAuthorized.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), domain.ErrUnAuthorized.Error())
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestDeleteBlog_BlogNotFound() {
	// Create mock data
	blogID := uuid.New()
	requesterID := uuid.New()

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("DeleteBlog", blogID, requesterID, false).Return(domain.ErrBlogNotFound)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: blogID.String()}}
	c.Set("id", requesterID.String())
	c.Set("is_admin", false)

	// Call the actual function
	suite.blogController.DeleteBlog(c)

	// Assertions
	assert.Equal(suite.T(), domain.ErrBlogNotFound.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), domain.ErrBlogNotFound.Error())
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestAddView_Success() {
	// Create mock data
	blogID := uuid.New()

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("AddView", blogID).Return(nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: blogID.String()}}

	// Call the actual function
	suite.blogController.AddView(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "View added successfully")
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestAddView_InvalidID() {
	// Simulate the HTTP request with an invalid UUID
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "invalid-id"}}

	// Call the actual function
	suite.blogController.AddView(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid ID")
	suite.mockUseCase.AssertNotCalled(suite.T(), "AddView")
}

func (suite *BlogControllerSuite) TestSearchBlogs_Success() {
	// Create mock data
	blogs := []dto.BlogDto{
		{
			Title:   "Test Title 1",
			Content: "Test Content 1",
			Tags:    []string{"tag1", "tag2"},
		},
		{
			Title:   "Test Title 2",
			Content: "Test Content 2",
			Tags:    []string{"tag3", "tag4"},
		},
	}
	totalPages := 1
	totalCount := 2

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("SearchBlogs", mock.AnythingOfType("domain.BlogFilter")).
		Return(blogs, totalPages, totalCount, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/blogs/search?title=Test", nil)

	// Call the actual function
	suite.blogController.SearchBlogs(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Test Title 1")
	assert.Contains(suite.T(), w.Body.String(), "Test Title 2")
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestSearchBlogs_InvalidParams() {
	blogs := []dto.BlogDto{
		{
			Title:   "Test Title 1",
			Content: "Test Content 1",
			Tags:    []string{"tag1", "tag2"},
		},
		{
			Title:   "Test Title 2",
			Content: "Test Content 2",
			Tags:    []string{"tag3", "tag4"},
		},
	}
	totalPages := 1
	totalCount := 2
	suite.mockUseCase.On("SearchBlogs", mock.AnythingOfType("domain.BlogFilter")).
		Return(blogs, totalPages, totalCount, nil)

	// Simulate the HTTP request with invalid pagination parameters
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/blogs/search?page=invalid&limit=invalid", nil)

	// Call the actual function
	suite.blogController.SearchBlogs(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	// Here we expect the default behavior as the controller handles invalid pagination gracefully
	assert.Contains(suite.T(), w.Body.String(), "blogs")
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestGenerateBlogContent_Success() {
	// Create mock data
	request := domain.BlogContentRequest{
		Topic: "Test Title",
		Keywords: []string{
			"keyword1", "keyword2",
		},
	}
	expectedContent := domain.BlogContentResponse{
		SuggestedContent: "Generated content based on the provided title and keywords.",
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("GenerateBlogContent", request).Return(&expectedContent, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(request)
	c.Request = httptest.NewRequest("POST", "/blogs/generate-content", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the actual function
	suite.blogController.GenerateBlogContent(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Generated content based on the provided title and keywords.")
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestGenerateBlogContent_InvalidJson() {
	// Simulate the HTTP request with an invalid JSON (missing required fields)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/blogs/generate-content", bytes.NewBuffer([]byte(`{}`)))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the actual function
	suite.blogController.GenerateBlogContent(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "error")
	suite.mockUseCase.AssertNotCalled(suite.T(), "GenerateBlogContent")
}

func (suite *BlogControllerSuite) TestGenerateBlogContent_Failure() {
	// Create mock data
	request := domain.BlogContentRequest{
		Topic:    "Test Title",
		Keywords: []string{"keyword1", "keyword2"},
	}
	expectedError := &domain.CustomError{
		Message:    "Failed to generate content",
		StatusCode: http.StatusInternalServerError,
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("GenerateBlogContent", request).Return(&domain.BlogContentResponse{}, expectedError)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(request)
	c.Request = httptest.NewRequest("POST", "/blogs/generate-content", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the actual function
	suite.blogController.GenerateBlogContent(c)

	// Assertions
	assert.Equal(suite.T(), expectedError.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), expectedError.Message)
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestSuggestImprovements_Success() {
	// Create mock data
	request := domain.SuggestionRequest{
		Content: "This is a test blog content.",
	}
	expectedSuggestions := domain.SuggestionResponse{
		Suggestions: "Consider adding more examples. Improve the introduction.",
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("SuggestImprovements", request.Content).Return(&expectedSuggestions, nil)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(request)
	c.Request = httptest.NewRequest("POST", "/blogs/suggest-improvements", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the actual function
	suite.blogController.SuggestImprovements(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Consider adding more examples.")
	assert.Contains(suite.T(), w.Body.String(), "Improve the introduction.")
	suite.mockUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestSuggestImprovements_InvalidJson() {
	// Simulate the HTTP request with an invalid JSON (missing required fields)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/blogs/suggest-improvements", bytes.NewBuffer([]byte(`{}`)))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the actual function
	suite.blogController.SuggestImprovements(c)

	// Assertions
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "error")
	suite.mockUseCase.AssertNotCalled(suite.T(), "SuggestImprovements")
}

func (suite *BlogControllerSuite) TestSuggestImprovements_Failure() {
	// Create mock data
	request := domain.SuggestionRequest{
		Content: "This is a test blog content.",
	}
	expectedError := &domain.CustomError{
		Message:    "Failed to generate suggestions",
		StatusCode: http.StatusInternalServerError,
	}

	// Mock the expected behavior of the usecase
	suite.mockUseCase.On("SuggestImprovements", request.Content).Return(&domain.SuggestionResponse{}, expectedError)

	// Simulate the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body, _ := json.Marshal(request)
	c.Request = httptest.NewRequest("POST", "/blogs/suggest-improvements", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the actual function
	suite.blogController.SuggestImprovements(c)

	// Assertions
	assert.Equal(suite.T(), expectedError.StatusCode, w.Code)
	assert.Contains(suite.T(), w.Body.String(), expectedError.Message)
	suite.mockUseCase.AssertExpectations(suite.T())
}


func TestBlogControllerSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerSuite))
}
