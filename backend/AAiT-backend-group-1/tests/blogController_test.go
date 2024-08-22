package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/delivery/controllers"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBlogController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected request body
	expectedBlog := &domain.Blog{
		Title:         "Test Blog Title",
		Content:       "Test Blog Content",
		AuthorUsername: "testuser",
	}

	// Serialize the expectedBlog to JSON
	jsonValue, err := json.Marshal(expectedBlog)
	if err != nil {
		t.Fatalf("Failed to marshal expected blog: %v", err)
	}

	// Define the expected input and output
	mockUsecase.On("CreateBlog", mock.AnythingOfType("*domain.Blog"), "12345").Return(nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set JSON as the content type
	c.Request = httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")

	// Call the CreateBlog method
	controller.CreateBlog(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedBlog.Title, response["data"].(map[string]interface{})["title"])
	assert.Equal(t, expectedBlog.Content, response["data"].(map[string]interface{})["content"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


func TestGetBlogController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected input and output
	expectedBlog := &domain.Blog{
		Title:         "Test Blog Title",
		Content:       "Test Blog Content",
		AuthorUsername: "testuser",
	}

	// Define the expected input and output
	mockUsecase.On("GetBlog", "1", "12345").Return(expectedBlog, nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Call the GetBlog method
	controller.GetBlog(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedBlog.Title, response["title"])
	assert.Equal(t, expectedBlog.Content, response["content"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


func TestGetBlogsController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected input and output
	expectedBlogs := []domain.Blog{
		{
			Title:         "Test Blog Title",
			Content:       "Test Blog Content",
			AuthorUsername: "testuser",
		},
		{
			Title:         "Test Blog Title 2",
			Content:       "Test Blog Content 2",
			AuthorUsername: "testuser",
		},
	}

	// Define the expected input and output
	mockUsecase.On("GetBlogs", "1").Return(expectedBlogs, nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")
	c.Request = httptest.NewRequest(http.MethodGet, "/blogs?p=1", nil)

	// Call the GetBlogs method
	controller.GetBlogs(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedBlogs[0].Title, response[0]["title"])
	assert.Equal(t, expectedBlogs[0].Content, response[0]["content"])
	assert.Equal(t, expectedBlogs[1].Title, response[1]["title"])
	assert.Equal(t, expectedBlogs[1].Content, response[1]["content"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}



func TestUpdateBlogController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected request body
	expectedBlog := &domain.Blog{
		Title:         "Test Blog Title",
		Content:       "Test Blog Content",
		AuthorUsername: "testuser",
	}

	// Serialize the expectedBlog to JSON
	jsonValue, err := json.Marshal(expectedBlog)
	if err != nil {
		t.Fatalf("Failed to marshal expected blog: %v", err)
	}

	// Define the expected input and output
	mockUsecase.On("UpdateBlog", "1", mock.AnythingOfType("*domain.Blog"), "12345").Return(nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set JSON as the content type
	c.Request = httptest.NewRequest(http.MethodPut, "/blogs/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Call the UpdateBlog method
	controller.UpdateBlog(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedBlog.Title, response["data"].(map[string]interface{})["title"])
	assert.Equal(t, expectedBlog.Content, response["data"].(map[string]interface{})["content"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


func TestDeleteBlogController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected input and output
	mockUsecase.On("DeleteBlog", "1").Return(nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Call the DeleteBlog method
	controller.DeleteBlog(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Blog deleted successfully", response["message"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


func TestSearchBlogsByTitleController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected input and output
	expectedBlogs := []domain.Blog{
		{
			Title:         "Test Blog Title",
			Content:       "Test Blog Content",
			AuthorUsername: "testuser",
		},
		{
			Title:         "Test Blog Title 2",
			Content:       "Test Blog Content 2",
			AuthorUsername: "testuser",
		},
	}

	// Define the expected input and output
	mockUsecase.On("SearchBlogsByTitle", "Test", "1").Return(expectedBlogs, nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")
	c.Request = httptest.NewRequest(http.MethodGet, "/blogs/search/title?title=Test&p=1", nil)

	// Call the SearchBlogsByTitle method
	controller.SearchBlogsByTitle(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedBlogs[0].Title, response[0]["title"])
	assert.Equal(t, expectedBlogs[0].Content, response[0]["content"])
	assert.Equal(t, expectedBlogs[1].Title, response[1]["title"])
	assert.Equal(t, expectedBlogs[1].Content, response[1]["content"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


func TestSearchBlogsByAuthorController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected input and output
	expectedBlogs := []domain.Blog{
		{
			Title:         "Test Blog Title",
			Content:       "Test Blog Content",
			AuthorUsername: "testuser",
		},
		{
			Title:         "Test Blog Title 2",
			Content:       "Test Blog Content 2",
			AuthorUsername: "testuser",
		},
	}

	// Define the expected input and output
	mockUsecase.On("SearchBlogsByAuthor", "Test", "1").Return(expectedBlogs, nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")
	c.Request = httptest.NewRequest(http.MethodGet, "/blogs/search/author?author=Test&p=1", nil)

	// Call the SearchBlogsByAuthor method
	controller.SearchBlogsByAuthor(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedBlogs[0].Title, response[0]["title"])
	assert.Equal(t, expectedBlogs[0].Content, response[0]["content"])
	assert.Equal(t, expectedBlogs[1].Title, response[1]["title"])
	assert.Equal(t, expectedBlogs[1].Content, response[1]["content"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


func TestFilterBlogsController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected input and output
	expectedBlogs := []domain.Blog{
		{
			Title:         "Test Blog Title",
			Content:       "Test Blog Content",
			AuthorUsername: "testuser",
		},
		{
			Title:         "Test Blog Title 2",
			Content:       "Test Blog Content 2",
			AuthorUsername: "testuser",
		},
	}

	// Define the expected input and output
	mockUsecase.On("FilterBlogs", []string{"test"}, mock.AnythingOfType("time.Time"), true).Return(expectedBlogs, nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")
	c.Request = httptest.NewRequest(http.MethodGet, "/blogs/filter?tags=test&date=2021-01-01&popular=true", nil)

	// Call the FilterBlogs method
	controller.FilterBlogs(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedBlogs[0].Title, response[0]["title"])
	assert.Equal(t, expectedBlogs[0].Content, response[0]["content"])
	assert.Equal(t, expectedBlogs[1].Title, response[1]["title"])
	assert.Equal(t, expectedBlogs[1].Content, response[1]["content"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}

func TestLikeBlogController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected input and output
	mockUsecase.On("LikeBlog", "12345", "1").Return(nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "12345")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Call the LikeBlog method
	controller.LikeBlog(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Blog liked successfully", response["message"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


func TestDislikeBlogController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected input and output
	mockUsecase.On("DislikeBlog", "12345", "1").Return(nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "12345")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Call the DislikeBlog method
	controller.DislikeBlog(c)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Blog disliked successfully", response["message"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


func TestAddCommentController(t *testing.T) {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Create a mock instance of the BlogUseCase
	mockUsecase := new(mocks.MockBlogUseCase)

	// Create a new blogController with the mock usecase
	controller := controllers.NewBlogController(mockUsecase)

	// Define the expected request body
	expectedComment := &domain.Comment{
		Content: "Test Comment Content",
		AuthorUsername: "testuser",
	}

	// Serialize the expectedComment to JSON
	jsonValue, err := json.Marshal(expectedComment)
	if err != nil {
		t.Fatalf("Failed to marshal expected comment: %v", err)
	}

	// Define the expected input and output
	mockUsecase.On("AddComment", "1", mock.AnythingOfType("*domain.Comment")).Return(nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set JSON as the content type
	c.Request = httptest.NewRequest(http.MethodPost, "/blogs/1/comments", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Call the AddComment method
	controller.AddComment(c)

	// Assert the status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedComment.Content, response["data"].(map[string]interface{})["content"])

	// Assert that the mock's expectations were met
	mockUsecase.AssertExpectations(t)
}


