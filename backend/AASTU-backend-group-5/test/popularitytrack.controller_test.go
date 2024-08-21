package test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type PopularityControllerTestSuite struct {
	suite.Suite
	mockPopularityUsecase *mocks.BlogPopularityUsecase
	popularityController  *controller.PopularityController
}

func (suite *PopularityControllerTestSuite) SetupTest() {
	suite.mockPopularityUsecase = mocks.NewBlogPopularityUsecase(suite.T())
	suite.popularityController = controller.NewPopularityController(suite.mockPopularityUsecase)
}

func (suite *PopularityControllerTestSuite) TestGetPopularBlogs_Success() {
	expectedBlogs := []domain.Blog{{Title: "Popular Blog 1"}, {Title: "Popular Blog 2"}}
	sortBy := "likes"
	sortOrder := -1

	suite.mockPopularityUsecase.On("GetSortedPopularBlogs", sortBy, sortOrder).Return(expectedBlogs, nil)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/popular-blogs?sort_by="+sortBy+"&sort_order="+strconv.Itoa(sortOrder), nil)
	c.Request = req

	suite.popularityController.GetPopularBlogs(c)

	suite.Equal(http.StatusOK, w.Code)

	var response []domain.Blog
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal(expectedBlogs, response)

	suite.mockPopularityUsecase.AssertExpectations(suite.T())
}

func (suite *PopularityControllerTestSuite) TestGetPopularBlogs_Error() {
	sortBy := "likes"
	sortOrder := -1

	suite.mockPopularityUsecase.On("GetSortedPopularBlogs", sortBy, sortOrder).Return(nil, errors.New("database error"))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/popular-blogs?sort_by="+sortBy+"&sort_order="+strconv.Itoa(sortOrder), nil)
	c.Request = req

	suite.popularityController.GetPopularBlogs(c)

	suite.Equal(http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Contains(response["error"], "database error")

	suite.mockPopularityUsecase.AssertExpectations(suite.T())
}

func (suite *PopularityControllerTestSuite) TestGetPopularBlogs_InvalidSortOrder() {
	expectedBlogs := []domain.Blog{{Title: "Popular Blog 1"}, {Title: "Popular Blog 2"}}
	sortBy := "likes"
	sortOrderStr := "invalid"

	suite.mockPopularityUsecase.On("GetSortedPopularBlogs", sortBy, -1).Return(expectedBlogs, nil)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/popular-blogs?sort_by="+sortBy+"&sort_order="+sortOrderStr, nil)
	c.Request = req

	suite.popularityController.GetPopularBlogs(c)

	suite.Equal(http.StatusOK, w.Code)

	var response []domain.Blog
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal(expectedBlogs, response)

	suite.mockPopularityUsecase.AssertExpectations(suite.T())
}

func TestPopularityControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PopularityControllerTestSuite))
}
