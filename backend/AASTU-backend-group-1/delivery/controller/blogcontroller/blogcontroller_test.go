package blogcontroller_test

import (
	"blogs/delivery/controller/blogcontroller"
	"blogs/domain"
	"blogs/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogControllerTestSuite struct {
	suite.Suite
	blogcontroller *blogcontroller.BlogController
	blogusecase    *mocks.BlogUsecase
}

func (suite *BlogControllerTestSuite) SetupTest() {
	suite.blogusecase = new(mocks.BlogUsecase)
	suite.blogcontroller = blogcontroller.NewBlogController(suite.blogusecase)
}

func (suite *BlogControllerTestSuite) TearDownTest() {
	suite.blogusecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestAddComment() {

	suite.Run("TestAddComment", func() {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = []gin.Param{{Key: "id", Value: "5f3f1b9b7f4b3b1b4c7f1b3b"}}
		body := struct {
			Content string `json:"content"`
		}{
			Content: "test",
		}
		bodyJSON, _ := json.Marshal(body)

		ctx.Request = httptest.NewRequest("POST", "/blogs/5f3f1b9b7f4b3b1b4c7f1b3b/comment", strings.NewReader(string(bodyJSON)))

		ctx.Set("claims", &domain.LoginClaims{
			Username: "test",
		})

		suite.blogusecase.On("AddComment", mock.Anything).Return(nil)
		suite.blogcontroller.AddComment(ctx)
		suite.Equal(suite.T(), http.StatusOK, w.Code)
	})

}

func (suite *BlogControllerTestSuite) TestGetBlogComments() {
	suite.Run("TestGetBlogComments", func() {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = []gin.Param{{Key: "id", Value: "5f3f1b9b7f4b3b1b4c7f1b3b"}}

		suite.blogusecase.On("GetBlogComments", mock.Anything).Return([]domain.Comment{}, nil)
		suite.blogcontroller.GetBlogComments(ctx)
		suite.Equal(suite.T(), http.StatusOK, w.Code)
	})
}

func (suite *BlogControllerTestSuite) TestDeleteLogByID() {
	suite.Run("TestDeleteLogByID", func() {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = []gin.Param{{Key: "id", Value: "5f3f1b9b7f4b3b1b4c7f1b3b"}}

		ctx.Set("claims", &domain.LoginClaims{
			Username: "test",
			Role:     "admin",
			Type:     "login",
		})

		suite.blogusecase.On("DeleteBlogByID", mock.Anything, mock.Anything).Return(nil)
		suite.blogcontroller.DeleteLogByID(ctx)
		suite.Equal(suite.T(), http.StatusOK, w.Code)
	})
}

func (suite *BlogControllerTestSuite) TestInsertBlog() {
	suite.Run("TestInsertBlog", func() {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		body := struct {
			Title   string   `json:"title"`
			Content string   `json:"content"`
			Tags    []string `json:"tags"`
		}{
			Title:   "test",
			Content: "test",
			Tags:    []string{"test"},
		}
		bodyJSON, _ := json.Marshal(body)

		ctx.Request = httptest.NewRequest("POST", "/blogs", strings.NewReader(string(bodyJSON)))

		ctx.Set("claims", &domain.LoginClaims{
			Username: "test",
			Role:     "admin",
			Type:     "login",
		})

		suite.blogusecase.On("InsertBlog", mock.Anything).Return(nil)
		suite.blogcontroller.InsertBlog(ctx)
		suite.Equal(suite.T(), http.StatusOK, w.Code)
	})

}

func (suite *BlogControllerTestSuite) TestGetBlogs() {
	suite.Run("GetBlogs", func() {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("GET", "/blogs", nil)

		ctx.Params = []gin.Param{{Key: "page", Value: "1"}, {Key: "size", Value: "10"}, {Key: "sort_by", Value: "date"}, {Key: "reverse", Value: "false"}}
		suite.blogusecase.On("GetBlogs", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]domain.Blog{}, nil)
		suite.blogcontroller.GetBlogs(ctx)
		suite.Equal(suite.T(), http.StatusOK, w.Code)
	})
}

func (suite *BlogControllerTestSuite) TestGetBlogByID() {
	suite.Run("GetBlogByID", func() {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = []gin.Param{{Key: "id", Value: "5f3f1b9b7f4b3b1b4c7f1b3b"}}

		ctx.Request = httptest.NewRequest("GET", "/blogs/5f3f1b9b7f4b3b1b4c7f1b3b", nil)

		suite.blogusecase.On("GetBlogByID", mock.Anything).Return(domain.Blog{}, nil)
		suite.blogcontroller.GetBlogByID(ctx)
		suite.Equal(suite.T(), http.StatusOK, w.Code)
	})
}

func (suite *BlogControllerTestSuite) TestAddLike() {
	suite.Run("AddLike", func() {

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = []gin.Param{{Key: "id", Value: "5f3f1b9b7f4b3b1b4c7f1b3b"}}

		body := struct {
			Like bool `bson:"like" json:"like"`
		}{
			Like: true,
		}

		bodyJSON, _ := json.Marshal(body)

		ctx.Request = httptest.NewRequest("POST", "/blogs/5f3f1b9b7f4b3b1b4c7f1b3b/like", strings.NewReader(string(bodyJSON)))

		ctx.Set("claims", &domain.LoginClaims{
			Username: "test",
		})

		suite.blogusecase.On("AddLike", mock.Anything, mock.Anything).Return(nil)
		suite.blogcontroller.AddLike(ctx)
		suite.Equal(suite.T(), http.StatusOK, w.Code)
	})
}
