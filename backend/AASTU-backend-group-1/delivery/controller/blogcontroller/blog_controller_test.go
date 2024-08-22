package blogcontroller_test

import (
	"blogs/config"
	"blogs/delivery/controller/blogcontroller"
	"blogs/domain"
	"blogs/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogControllerTestSuite struct {
	suite.Suite
	blogcontroller *blogcontroller.BlogController
	blogusecase    *mocks.BlogUsecase
	w              *httptest.ResponseRecorder
	ctx            *gin.Context
}

func (suite *BlogControllerTestSuite) SetupSuite() {
	suite.blogusecase = new(mocks.BlogUsecase)
	suite.blogcontroller = blogcontroller.NewBlogController(suite.blogusecase)
}

func (suite *BlogControllerTestSuite) TearDownTest() {
	suite.blogusecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) SetupTest() {
	suite.w = httptest.NewRecorder()
	suite.ctx, _ = gin.CreateTestContext(suite.w)
}

func (suite *BlogControllerTestSuite) TestAddComment_Success() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	body := gin.H{"content": "some comment"}

	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs/"+objectID.Hex()+"/comment",
		strings.NewReader(string(bodyJSON)),
	)

	suite.ctx.Set("claims", &domain.LoginClaims{Username: "user_1"})
	expectedBody := gin.H{"message": "Comment added successfully"}
	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogusecase.On("AddComment", mock.Anything).Return(nil).Once()
	suite.blogcontroller.AddComment(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestAddComment_Fail() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	body := gin.H{"content": "some comment"}

	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs/"+objectID.Hex()+"/comment",
		strings.NewReader(string(bodyJSON)),
	)

	suite.ctx.Set("claims", &domain.LoginClaims{Username: "user_1"})
	expectedBody := gin.H{"error": "blog not found"}
	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogusecase.On("AddComment", mock.Anything).Return(config.ErrBlogNotFound).Once()
	suite.blogcontroller.AddComment(suite.ctx)
	suite.Equal(http.StatusNotFound, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGetBlogComments_Empty() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	suite.blogusecase.On("GetBlogComments", mock.Anything).Return([]*domain.Comment{}, nil).Once()

	expectedBody := gin.H{"data": []*domain.Comment{}, "counts": 0}
	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.GetBlogComments(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGetBlogComments_Success() {
	objecID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objecID.Hex()}}
	comments := []*domain.Comment{
		{
			BlogID:  objecID,
			Author:  "user_1",
			Content: "some comment",
			Date:    time.Now(),
		},
		{
			BlogID:  objecID,
			Author:  "user_2",
			Content: "some other comment",
			Date:    time.Now().Add(time.Hour),
		},
	}

	expectedBody := gin.H{"data": comments, "counts": len(comments)}
	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogusecase.On("GetBlogComments", mock.Anything).Return(comments, nil).Once()
	suite.blogcontroller.GetBlogComments(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestDeleteBlogByID_Success() {
	objecID := primitive.NewObjectID()
	claims := &domain.LoginClaims{Username: "user_1"}

	suite.ctx.Params = []gin.Param{{Key: "id", Value: objecID.Hex()}}
	suite.ctx.Set("claims", claims)
	suite.blogusecase.On("DeleteBlogByID", objecID.Hex(), claims).Return(nil).Once()

	suite.blogcontroller.DeleteBlogByID(suite.ctx)
	suite.Equal(http.StatusNoContent, suite.w.Code)
}

func (suite *BlogControllerTestSuite) TestDeleteBlogByID_Fail() {
	objecID := primitive.NewObjectID()
	claims := &domain.LoginClaims{Username: "user_1"}

	suite.ctx.Params = []gin.Param{{Key: "id", Value: objecID.Hex()}}
	suite.ctx.Set("claims", claims)
	suite.blogusecase.On("DeleteBlogByID", objecID.Hex(), claims).Return(config.ErrBlogNotFound).Once()

	expectedBody := gin.H{"error": "blog not found"}
	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.DeleteBlogByID(suite.ctx)
	suite.Equal(http.StatusNotFound, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestInsertBlog_Success() {
	body := gin.H{
		"title":   "some title",
		"content": "this is a random content",
		"tags":    []string{"tag1", "tag2"},
	}

	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs",
		strings.NewReader(string(bodyJSON)),
	)

	claims := &domain.LoginClaims{
		Username: "test",
		Role:     "user",
		Type:     "login",
	}

	suite.ctx.Set("claims", claims)

	blogData := &domain.Blog{
		ID:            primitive.NewObjectID(),
		Title:         body["title"].(string),
		Content:       body["content"].(string),
		Author:        claims.Username,
		Tags:          body["tags"].([]string),
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount:    0,
		LikesCount:    0,
		CommentsCount: 0,
	}

	suite.blogusecase.On("InsertBlog", mock.Anything).Return(blogData, nil).Once()
	expected, err := json.Marshal(blogData)
	suite.NoError(err)

	suite.blogcontroller.InsertBlog(suite.ctx)
	suite.Equal(http.StatusCreated, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGetBlogs_Success() {
	suite.ctx.Request = httptest.NewRequest("GET", "/blogs", nil)
	blogs := []*domain.Blog{
		{
			ID:            primitive.NewObjectID(),
			Title:         "some title",
			Content:       "this is a random content",
			Author:        "test_user",
			Tags:          []string{"tag1", "tag2"},
			CreatedAt:     time.Now(),
			LastUpdatedAt: time.Now(),
			ViewsCount:    10,
			LikesCount:    2,
			CommentsCount: 3,
		},
		{
			ID:            primitive.NewObjectID(),
			Title:         "some other title",
			Content:       "this is another random content",
			Author:        "test_user",
			Tags:          []string{"tag3", "tag4"},
			CreatedAt:     time.Now(),
			LastUpdatedAt: time.Now(),
			ViewsCount:    20,
			LikesCount:    5,
			CommentsCount: 7,
		},
	}

	suite.ctx.Request = httptest.NewRequest(
		"GET",
		"/blogs?page=1&size=3&sort_by=date&reverse=false",
		nil,
	)

	suite.blogusecase.On("GetBlogs", "date", 1, 3, false).Return(blogs, nil)
	suite.blogcontroller.GetBlogs(suite.ctx)

	expectedBody := gin.H{"count": len(blogs), "data": blogs}
	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGetBlogByID_Success() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	suite.ctx.Request = httptest.NewRequest(
		"GET",
		"/blogs/"+objectID.Hex(),
		nil,
	)

	blog := &domain.Blog{
		ID:            objectID,
		Title:         "some title",
		Content:       "this is a random content",
		Author:        "test_user",
		Tags:          []string{"tag1", "tag2"},
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount:    10,
		LikesCount:    2,
		CommentsCount: 3,
	}

	suite.blogusecase.On("GetBlogByID", mock.Anything).Return(blog, nil)

	expected, err := json.Marshal(blog)
	suite.NoError(err)

	suite.blogcontroller.GetBlogByID(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}
