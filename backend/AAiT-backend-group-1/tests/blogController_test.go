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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogControllerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.MockBlogUseCase
	controller  domain.BlogController
}

func (suite *BlogControllerTestSuite) SetupTest() {
	// Set the Gin mode to Test mode
	gin.SetMode(gin.TestMode)

	// Initialize the mock usecase and controller before each test
	suite.mockUsecase = new(mocks.MockBlogUseCase)
	suite.controller = controllers.NewBlogController(suite.mockUsecase)
}

func (suite *BlogControllerTestSuite) TestCreateBlogController() {
	// Define the expected request body
	expectedBlog := &domain.Blog{
		Title:          "Test Blog Title",
		Content:        "Test Blog Content",
		AuthorUsername: "testuser",
	}

	// Serialize the expectedBlog to JSON
	jsonValue, err := json.Marshal(expectedBlog)
	suite.Require().NoError(err)

	// Define the expected input and output
	suite.mockUsecase.On("CreateBlog", mock.AnythingOfType("*domain.Blog"), "12345").Return(nil)

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
	suite.controller.CreateBlog(c)

	// Assert the status code
	suite.Equal(http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)
	suite.Equal(expectedBlog.Title, response["data"].(map[string]interface{})["title"])
	suite.Equal(expectedBlog.Content, response["data"].(map[string]interface{})["content"])

	// Assert that the mock's expectations were met
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGetBlogController() {
	// Define the expected input and output
	expectedBlog := &domain.Blog{
		Title:          "Test Blog Title",
		Content:        "Test Blog Content",
		AuthorUsername: "testuser",
	}

	// Define the expected input and output
	suite.mockUsecase.On("GetBlog", "1", "12345").Return(expectedBlog, nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "12345")
	c.Set("username", "testuser")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Call the GetBlog method
	suite.controller.GetBlog(c)

	// Assert the status code
	suite.Equal(http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)
	suite.Equal(expectedBlog.Title, response["title"])
	suite.Equal(expectedBlog.Content, response["content"])

	// Assert that the mock's expectations were met
	suite.mockUsecase.AssertExpectations(suite.T())
}

// Define other test methods similarly...

func (suite *BlogControllerTestSuite) TestDislikeBlogController() {
	// Define the expected input and output
	suite.mockUsecase.On("DisLike", "1", "1").Return(nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set context values
	c.Set("user_id", "1")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Call the DislikeBlog method
	suite.controller.DislikeBlog(c)

	// Assert the status code
	suite.Equal(http.StatusOK, w.Code)

	// Assert the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.Require().NoError(err)

	// Check if response contains "message"
	msg, ok := response["message"].(string)
	suite.True(ok)
	suite.Equal("Blog Disliked successfully", msg)

	// Assert that the mock's expectations were met
	suite.mockUsecase.AssertExpectations(suite.T())
}

func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}
