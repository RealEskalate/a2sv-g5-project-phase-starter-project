package controller_test

import (
	"blog/config"
	"blog/delivery/controller"
	"blog/domain"
	"blog/domain/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogControllerSuite struct {
	suite.Suite
	router          *gin.Engine
	BlogUsecase    *mocks.BlogUsecase
	BlogController *controller.BlogController
}

func (suite *BlogControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.BlogUsecase = new(mocks.BlogUsecase)
	env := &config.Env{}
	suite.BlogController = &controller.BlogController{
		BlogUsecase: suite.BlogUsecase,
		Env:          env,
	}
	suite.router = gin.Default()
	suite.router.POST("/blogs/", suite.BlogController.CreateBlog)
	suite.router.POST("/blogs/:id", suite.BlogController.GetBlogByID)
	suite.router.POST("/blogs/:id", suite.BlogController.UpdateBlog)
	suite.router.POST("/blogs/:id", suite.BlogController.DeleteBlog)
}
func (suite *BlogControllerSuite) TearDownTest() {
	suite.BlogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerSuite) TestCreateBlog() {
	suite.Run("create_blog_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		blog := domain.Blog{
			Title:    "title",
			Content:  "content",
		}
		resp := domain.Blog{}
		suite.BlogUsecase.On("CreateBlog", mock.Anything, &blog).Return(&resp, nil).Once()
		payload, _ := json.Marshal(blog)
		req, _ := http.NewRequest(http.MethodPost, "/blogs/", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("create_blog_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		blog := domain.Blog{
			Title:    "title",
			Content:  "content",
		}
		suite.BlogUsecase.On("CreateBlog", mock.Anything, &blog).Return(nil, errors.New("error")).Once()
		payload, _ := json.Marshal(blog)
		req, _ := http.NewRequest(http.MethodPost, "/blogs/", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func (suite *BlogControllerSuite) TestGetBlogByID() {
	suite.Run("get_blog_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		blog := domain.Blog{
			Title:    "title",
			Content:  "content",
		}
		resp := domain.Blog{}
		suite.BlogUsecase.On("GetBlogByID", mock.Anything, blog.ID).Return(&resp, nil).Once()
		payload, _ := json.Marshal(blog)
		req, _ := http.NewRequest(http.MethodPost, "/blogs/:id", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("get_blog_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		blog := domain.Blog{
			Title:    "title",
			Content:  "content",
		}
		suite.BlogUsecase.On("GetBlogByID", mock.Anything, blog.ID).Return(nil, errors.New("error")).Once()
		payload, _ := json.Marshal(blog)
		req, _ := http.NewRequest(http.MethodPost, "/blogs/:id", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func (suite *BlogControllerSuite) TestUpdateBlog() {
	suite.Run("update_blog_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		blog := domain.Blog{
			Title:    "title",
			Content:  "content",
		}
		resp := domain.Blog{}
		suite.BlogUsecase.On("UpdateBlog", mock.Anything, &blog).Return(&resp, nil).Once()
		payload, _ := json.Marshal(blog)
		req, _ := http.NewRequest(http.MethodPost, "/blogs/:id", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("update_blog_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		blog := domain.Blog{
			Title:    "title",
			Content:  "content",
		}
		suite.BlogUsecase.On("UpdateBlog", mock.Anything, &blog).Return(nil, errors.New("error")).Once()
		payload, _ := json.Marshal(blog)
		req, _ := http.NewRequest(http.MethodPost, "/blogs/:id", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func (suite *BlogControllerSuite) TestDeleteBlog() {
	suite.Run("delete_blog_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		blog := domain.Blog{
			Title:    "title",
			Content:  "content",
		}
		suite.BlogUsecase.On("DeleteBlog", mock.Anything, blog.ID).Return(nil).Once()
		payload, _ := json.Marshal(blog)
		req, _ := http.NewRequest(http.MethodPost, "/blogs/:id", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("delete_blog_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		blog := domain.Blog{
			Title:    "title",
			Content:  "content",
		}
		suite.BlogUsecase.On("DeleteBlog", mock.Anything, blog.ID).Return(errors.New("error")).Once()
		payload, _ := json.Marshal(blog)
		req, _ := http.NewRequest(http.MethodPost, "/blogs/:id", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func TestBlogControllerSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerSuite))
}
