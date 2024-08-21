package controllers_test

import (
	"blog_g2/deliveries/controllers"
	"blog_g2/domain"
	"blog_g2/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogControllerSuite struct {
	suite.Suite
	BlogController     *controllers.BlogController
	MockBlogUsecase    *mocks.BlogUsecase
	MockLikeUsecase    *mocks.LikeUsecase
	MockCommentUsecase *mocks.CommentUsecase
	MockDislikeUsecase *mocks.DisLikeUsecase
	MockAIService      *mocks.AIService
}

func (suite *BlogControllerSuite) SetupTest() {
	suite.MockBlogUsecase = new(mocks.BlogUsecase)
	suite.MockLikeUsecase = new(mocks.LikeUsecase)
	suite.MockCommentUsecase = new(mocks.CommentUsecase)
	suite.MockDislikeUsecase = new(mocks.DisLikeUsecase)
	suite.MockAIService = new(mocks.AIService)
	suite.BlogController = controllers.NewBlogController(
		suite.MockBlogUsecase,
		suite.MockLikeUsecase,
		suite.MockCommentUsecase,
		suite.MockDislikeUsecase,
		suite.MockAIService,
	)
	gin.SetMode(gin.TestMode)
}

func (suite *BlogControllerSuite) TestCreateBlog() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	requestBody := `{"title": "Test Blog", "content": "Test Content", "tags": ["tag1", "tag2"]}`
	c.Request, _ = http.NewRequest(http.MethodPost, "/blogs", strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.MockBlogUsecase.On("CreateBlog", mock.Anything, mock.AnythingOfType("*domain.Blog")).Return(nil)

	suite.BlogController.CreateBlog(c)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.MockBlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestRetrieveBlog() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	c.Request, _ = http.NewRequest(http.MethodGet, "/blogs?page=1", nil)

	mockBlogs := []domain.Blog{
		{Title: "Blog 1", Content: "Content 1", Date: time.Now()},
		{Title: "Blog 2", Content: "Content 2", Date: time.Now()},
	}

	suite.MockBlogUsecase.On("RetrieveBlog", mock.Anything, 1).Return(mockBlogs, nil)

	suite.BlogController.RetrieveBlog(c)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.MockBlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestUpdateBlog() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	getID := "12345"
	c.Params = gin.Params{{Key: "id", Value: getID}}

	requestBody := `{"title": "Updated Blog", "content": "Updated Content"}`
	c.Request, _ = http.NewRequest(http.MethodPut, "/blogs/"+getID, strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.MockBlogUsecase.On("UpdateBlog", mock.Anything, mock.AnythingOfType("domain.Blog"), getID).Return(nil)

	suite.BlogController.UpdateBlog(c)

	assert.Equal(suite.T(), http.StatusAccepted, recorder.Code)
	suite.MockBlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestDeleteBlog() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	getID := "12345"
	c.Params = gin.Params{{Key: "id", Value: getID}}

	c.Request, _ = http.NewRequest(http.MethodDelete, "/blogs/"+getID, nil)

	suite.MockBlogUsecase.On("DeleteBlog", mock.Anything, getID).Return(nil)

	suite.BlogController.DeleteBlog(c)

	assert.Equal(suite.T(), http.StatusAccepted, recorder.Code)
	suite.MockBlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestSearchBlog() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	c.Request, _ = http.NewRequest(http.MethodGet, "/blogs/search?name=Test&user=User1", nil)

	mockBlogs := []domain.Blog{
		{Title: "Test Blog", Content: "Test Content", Date: time.Now()},
	}

	suite.MockBlogUsecase.On("SearchBlog", mock.Anything, "Test", "User1").Return(mockBlogs, nil)

	suite.BlogController.SearchBlog(c)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.MockBlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestFilterBlog() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	tags := []string{"tag1", "tag2"}
	date := "2024-01-01"

	convDate, _ := time.Parse("2006-01-02", date)

	c.Request, _ = http.NewRequest(http.MethodGet, "/blogs/filter?tags[]=tag1&tags[]=tag2&date="+date, nil)

	mockBlogs := []domain.Blog{
		{Title: "Filtered Blog", Content: "Filtered Content", Date: convDate},
	}

	suite.MockBlogUsecase.On("FilterBlog", mock.Anything, tags, convDate).Return(mockBlogs, nil)

	suite.BlogController.FilterBlog(c)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.MockBlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestGeneratePost() {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	requestBody := `{"title": "Generated Title", "content": "Generated Content"}`
	c.Request, _ = http.NewRequest(http.MethodPost, "/blogs/generate", strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	mockPost := domain.PostRequest{Title: "Generated Title", Content: "Generated Content"}

	suite.MockAIService.On("GeneratePost", mockPost.Title, mockPost.Content).Return(mockPost, nil)

	suite.BlogController.GeneratePost(c)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.MockAIService.AssertExpectations(suite.T())
}

func TestBlogControllerSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerSuite))
}
