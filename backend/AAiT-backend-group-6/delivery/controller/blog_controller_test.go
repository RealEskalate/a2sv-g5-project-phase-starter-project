package controller

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type BlogControllerSuite struct {
	suite.Suite
	blogController *BlogController
	blogUsecase    *mocks.BlogUseCase
	redisClient    *mocks.Client
}

// SetupTest is executed before each test in the suite
func (suite *BlogControllerSuite) SetupTest() {
	suite.blogUsecase = new(mocks.BlogUseCase)
	suite.redisClient = new(mocks.Client)
	suite.blogController = NewBlogController(suite.blogUsecase, suite.redisClient)
}

func (suite *BlogControllerSuite) TestCreateBlog() {
    // Mock blog data
    blog := domain.Blog{
        Title:   "Test Blog",
        Content: "This is a test blog",
    }

    // Expected response after creating the blog
    createdBlog := domain.Blog{
        ID:      primitive.NewObjectID(),
        Title:   blog.Title,
        Content: blog.Content,
    }

    // Mock the use case method
    suite.blogUsecase.On("CreateBlog", mock.Anything, &blog).Return(createdBlog, nil)

    // Set up the Gin context with a JSON body
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Simulate a POST request with JSON data
    jsonBody, _ := json.Marshal(blog)
    c.Request, _ = http.NewRequest("POST", "/blogs", bytes.NewBuffer(jsonBody))
    c.Request.Header.Set("Content-Type", "application/json")

    // Call the CreateBlog handler
    suite.blogController.CreateBlog(c)

    // Convert the expected data to JSON for comparison
    expectedResponse := gin.H{
        "message": "Blog Created Successfully",
        "Blog":    createdBlog,
    }
    expectedJSON, _ := json.Marshal(expectedResponse)

    // Assert the status code and response body
    suite.Equal(http.StatusCreated, w.Code)
    suite.JSONEq(string(expectedJSON), w.Body.String())

    // Assert that the expected methods were called
    suite.blogUsecase.AssertExpectations(suite.T())
}

func TestBlogControllerSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerSuite))
}
