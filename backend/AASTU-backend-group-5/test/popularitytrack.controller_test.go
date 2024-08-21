package test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/blogpost/delivery/controller"
	"github.com/RealEskalate/blogpost/domain"
	"github.com/RealEskalate/blogpost/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetPopularBlogs(t *testing.T) {
	mockUsecase := new(mocks.BlogPopularityUsecase)
	popularityController := controller.NewPopularityController(mockUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("valid sort_by and sort_order", func(t *testing.T) {
		router := gin.Default()
		router.GET("/blogs/popular", popularityController.GetPopularBlogs())

		expectedBlogs := []domain.Blog{
			{
				ID:           primitive.NewObjectID(),
				Title:        "Sample Blog 1",
				Content:      "This is a sample blog content",
				LikeCount:    100,
				CommentCount: 50,
				DisLikeCount: 5,
			},
			{
				ID:           primitive.NewObjectID(),
				Title:        "Sample Blog 2",
				Content:      "This is another blog content",
				LikeCount:    75,
				CommentCount: 30,
				DisLikeCount: 2,
			},
		}

		mockUsecase.On("GetSortedPopularBlogs", []domain.SortBy{domain.SortByLikeCount}, []domain.SortOrder{domain.SortOrderDescending}).
			Return(expectedBlogs, nil)

		req, _ := http.NewRequest(http.MethodGet, "/blogs/popular?sort_by=likes&sort_order=-1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		mockUsecase.AssertExpectations(t)

	})

	t.Run("invalid sort_by parameter", func(t *testing.T) {
		router := gin.Default()
		router.GET("/blogs/popular", popularityController.GetPopularBlogs())

		req, _ := http.NewRequest(http.MethodGet, "/blogs/popular?sort_by=invalid_sort&sort_order=-1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		mockUsecase.AssertNotCalled(t, "GetSortedPopularBlogs")
	})

	t.Run("invalid sort_order parameter", func(t *testing.T) {
		router := gin.Default()
		router.GET("/blogs/popular", popularityController.GetPopularBlogs())

		req, _ := http.NewRequest(http.MethodGet, "/blogs/popular?sort_by=likes&sort_order=invalid", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		mockUsecase.AssertNotCalled(t, "GetSortedPopularBlogs")
	})

	t.Run("usecase returns error", func(t *testing.T) {
		router := gin.Default()
		router.GET("/blogs/popular", popularityController.GetPopularBlogs())

		mockUsecase.On("GetSortedPopularBlogs", []domain.SortBy{domain.SortByLikeCount}, []domain.SortOrder{domain.SortOrderDescending}).
			Return(nil, errors.New("internal server error"))

		req, _ := http.NewRequest(http.MethodGet, "/blogs/popular?sort_by=likes&sort_order=-1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("empty sort_by and sort_order", func(t *testing.T) {
		router := gin.Default()
		router.GET("/blogs/popular", popularityController.GetPopularBlogs())

		expectedBlogs := []domain.Blog{}

		mockUsecase.On("GetSortedPopularBlogs", mock.Anything, mock.Anything).
			Return(expectedBlogs, nil)

		req, _ := http.NewRequest(http.MethodGet, "/blogs/popular", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Empty(t, resp.Body.String())
		mockUsecase.AssertExpectations(t)
	})
}
