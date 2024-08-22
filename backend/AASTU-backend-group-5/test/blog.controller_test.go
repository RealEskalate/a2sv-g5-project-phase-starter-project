package test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type BlogControllerTestSuite struct {
	suite.Suite
	mockBlogUsecase *mocks.Blog_Usecase_interface
	blogController  *controller.BlogController
}

func (suite *BlogControllerTestSuite) SetupTest() {
	suite.mockBlogUsecase = mocks.NewBlog_Usecase_interface(suite.T())
	suite.blogController = &controller.BlogController{
		BlogUsecase: suite.mockBlogUsecase,
	}
}

func (suite *BlogControllerTestSuite) TestGetAllBlogs_Success() {
	expectedBlogs := []domain.Blog{{Title: "Test Blog"}}
	suite.mockBlogUsecase.On("GetBlogs", 10, 1).Return(expectedBlogs, nil)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/blogs", nil)

	suite.blogController.GetAllBlogs()(c)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("All blog posts retrieved successfully!", response["message"])
	suite.NotNil(response["blogs"])

	suite.mockBlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGetAllBlogs_Error() {
	suite.mockBlogUsecase.On("GetBlogs", 10, 1).Return(nil, errors.New("database error"))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/blogs", nil)

	suite.blogController.GetAllBlogs()(c)

	suite.Equal(http.StatusNotFound, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Contains(response["error"], "Failed to retrieve blog posts")

	suite.mockBlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGetAllBlogs_CustomPagination() {
	expectedBlogs := []domain.Blog{{Title: "Test Blog"}}
	suite.mockBlogUsecase.On("GetBlogs", 5, 2).Return(expectedBlogs, nil)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/blogs?limit=5&page_number=2", nil)

	suite.blogController.GetAllBlogs()(c)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("All blog posts retrieved successfully!", response["message"])
	suite.NotNil(response["blogs"])

	suite.mockBlogUsecase.AssertExpectations(suite.T())
}

func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}
