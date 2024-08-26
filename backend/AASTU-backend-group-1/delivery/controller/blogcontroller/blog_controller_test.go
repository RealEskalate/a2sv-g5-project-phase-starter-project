package blogcontroller_test

import (
	"blogs/config"
	"blogs/delivery/controller/blogcontroller"
	"blogs/domain"
	"blogs/mocks"
	"encoding/json"
	"errors"
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
	expectedBody := domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Comment added",
	}

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
	expectedBody := domain.APIResponse{
		Status:  http.StatusNotFound,
		Message: "Error adding comment",
		Error:   "blog not found",
	}

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

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Comments retrieved",
		Data:    []domain.Comment{},
	}

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

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Comments retrieved",
		Count:   len(comments),
		Data:    comments,
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogusecase.On("GetBlogComments", mock.Anything).Return(comments, nil).Once()
	suite.blogcontroller.GetBlogComments(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestDeleteCommentByID_Success() {
	objecID := primitive.NewObjectID()
	objecID2 := primitive.NewObjectID()
	claims := &domain.LoginClaims{Username: "user_1"}

	suite.ctx.Params = []gin.Param{{Key: "id", Value: objecID.Hex()}, {Key: "commentid", Value: objecID2.Hex()}}
	suite.ctx.Set("claims", claims)
	suite.blogusecase.On("DeleteComment", objecID2.Hex(), claims).Return(nil).Once()

	suite.blogcontroller.DeleteComment(suite.ctx)
	suite.Equal(http.StatusNoContent, suite.w.Code)
	suite.Equal("", suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestDeleteCommentByID_Fail() {
	objecID := primitive.NewObjectID()
	objecID2 := primitive.NewObjectID()
	claims := &domain.LoginClaims{Username: "user_1"}

	suite.ctx.Params = []gin.Param{{Key: "id", Value: objecID.Hex()}, {Key: "commentid", Value: objecID2.Hex()}}
	suite.ctx.Set("claims", claims)
	suite.blogusecase.On("DeleteComment", objecID2.Hex(), claims).Return(config.ErrCommentNotFound).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusNotFound,
		Message: "Error deleting comment",
		Error:   "comment not found",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.DeleteComment(suite.ctx)

	suite.Equal(http.StatusNotFound, suite.w.Code)
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

	expectedBody := domain.APIResponse{
		Status:  http.StatusNotFound,
		Message: "Error deleting blog",
		Error:   "blog not found",
	}
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

	expectedBody := domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Blog created",
		Data:    blogData,
	}
	expected, err := json.Marshal(expectedBody)
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

	suite.blogusecase.On("GetBlogs", "date", 1, 3, false).Return(blogs, 2, nil).Once()
	suite.blogcontroller.GetBlogs(suite.ctx)

	expectedBody := domain.APIResponse{
		Status:    http.StatusOK,
		Message:   "Blogs retrieved",
		PageCount: 2,
		Count:     len(blogs),
		Data:      blogs,
	}

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

	suite.blogusecase.On("GetBlogByID", mock.Anything).Return(blog, nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Blog retrieved",
		Data:    blog,
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.GetBlogByID(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGetBlogByID_Fail() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	suite.ctx.Request = httptest.NewRequest(
		"GET",
		"/blogs/"+objectID.Hex(),
		nil,
	)

	suite.blogusecase.On("GetBlogByID", mock.Anything).Return(nil, config.ErrBlogNotFound).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusNotFound,
		Message: "Cannot get blog",
		Error:   "blog not found",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.GetBlogByID(suite.ctx)
	suite.Equal(http.StatusNotFound, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestAddLike_Success() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	suite.ctx.Set("claims", &domain.LoginClaims{Username: "user_1"})

	body := gin.H{"like": true}
	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs/"+objectID.Hex()+"/like",
		strings.NewReader(string(bodyJSON)),
	)

	addedLike := &domain.Like{
		BlogID: objectID,
		User:   "user_1",
		Like:   true,
	}

	suite.blogusecase.On("AddLike", addedLike).Return(nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Like added",
		Data:    addedLike,
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.AddLike(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestAddLike_Fail() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	suite.ctx.Set("claims", &domain.LoginClaims{Username: "user_1"})

	body := gin.H{"like": true}
	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs/"+objectID.Hex()+"/like",
		strings.NewReader(string(bodyJSON)),
	)

	addedLike := &domain.Like{
		BlogID: objectID,
		User:   "user_1",
		Like:   true,
	}

	suite.blogusecase.On("AddLike", addedLike).Return(config.ErrBlogNotFound).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusNotFound,
		Message: "Error adding like",
		Error:   "blog not found",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.AddLike(suite.ctx)
	suite.Equal(http.StatusNotFound, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestSearchBlog_Success() {
	suite.ctx.Request = httptest.NewRequest(
		"GET",
		"/blogs/search?title=some%20title&author=test_user&tags=tag1&tags=tag2",
		nil,
	)

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

	suite.blogusecase.On("SearchBlog", "some title", "test_user", []string{"tag1", "tag2"}).Return(blogs, nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Count:   len(blogs),
		Data:    blogs,
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.SearchBlog(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestSearchBlog_Fail() {
	suite.ctx.Request = httptest.NewRequest(
		"GET",
		"/blogs/search?title=some%20title&author=test_user&tags=tag1&tags=tag2",
		nil,
	)

	suite.blogusecase.On("SearchBlog", "some title", "test_user", []string{"tag1", "tag2"}).Return(nil, config.ErrBlogNotFound).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusNotFound,
		Message: "Cannot search blog",
		Error:   "blog not found",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.SearchBlog(suite.ctx)
	suite.Equal(http.StatusNotFound, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestUpdateBlog_Success() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	claims := &domain.LoginClaims{Username: "user_1"}
	suite.ctx.Set("claims", claims)

	body := gin.H{
		"title":   "some new title",
		"content": "this is a new random content",
		"tags":    []string{"tag3", "tag4"},
	}

	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"PUT",
		"/blogs/"+objectID.Hex(),
		strings.NewReader(string(bodyJSON)),
	)

	sentBlog := &domain.Blog{
		Title:   body["title"].(string),
		Content: body["content"].(string),
		Tags:    body["tags"].([]string),
	}

	newBlog := &domain.Blog{
		ID:            objectID,
		Title:         body["title"].(string),
		Content:       body["content"].(string),
		Tags:          body["tags"].([]string),
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		ViewsCount:    0,
		LikesCount:    0,
		CommentsCount: 0,
	}

	suite.blogusecase.On("UpdateBlogByID", objectID.Hex(), sentBlog, claims).Return(newBlog, nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Blog updated",
		Data:    newBlog,
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.UpdateBlogByID(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestUpdateBlog_Fail() {
	objectID := primitive.NewObjectID()
	suite.ctx.Params = []gin.Param{{Key: "id", Value: objectID.Hex()}}
	claims := &domain.LoginClaims{Username: "user_1"}
	suite.ctx.Set("claims", claims)

	body := gin.H{
		"title":   "some new title",
		"content": "this is a new random content",
		"tags":    []string{"tag3", "tag4"},
	}

	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"PUT",
		"/blogs/"+objectID.Hex(),
		strings.NewReader(string(bodyJSON)),
	)

	sentBlog := &domain.Blog{
		Title:   body["title"].(string),
		Content: body["content"].(string),
		Tags:    body["tags"].([]string),
	}

	suite.blogusecase.On("UpdateBlogByID", objectID.Hex(), sentBlog, claims).Return(nil, config.ErrBlogNotFound).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusNotFound,
		Message: "Cannot update blog",
		Error:   "blog not found",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.UpdateBlogByID(suite.ctx)
	suite.Equal(http.StatusNotFound, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestAddView_Success() {
	objectID := primitive.NewObjectID()
	objectID2 := primitive.NewObjectID()
	objectID3 := primitive.NewObjectID()

	blogIDs := []string{
		objectID.Hex(),
		objectID2.Hex(),
		objectID3.Hex(),
	}

	body := gin.H{"ids": blogIDs}
	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs/view",
		strings.NewReader(string(bodyJSON)),
	)

	claims := &domain.LoginClaims{Username: "user_1"}
	suite.ctx.Set("claims", claims)

	objectIDs := []primitive.ObjectID{
		objectID,
		objectID2,
		objectID3,
	}

	suite.blogusecase.On("AddView", objectIDs, *claims).Return(nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Views added",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.AddView(suite.ctx)
	suite.Equal(http.StatusCreated, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestAddView_Fail() {
	objectID := primitive.NewObjectID()
	objectID2 := primitive.NewObjectID()

	blogIDs := []string{
		objectID.Hex(),
		objectID2.Hex(),
	}

	body := gin.H{"ids": blogIDs}
	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs/view",
		strings.NewReader(string(bodyJSON)),
	)

	claims := &domain.LoginClaims{Username: "user_1"}
	suite.ctx.Set("claims", claims)

	objectIDs := []primitive.ObjectID{
		objectID,
		objectID2,
	}

	suite.blogusecase.On("AddView", objectIDs, *claims).Return(config.ErrBlogNotFound).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusNotFound,
		Message: "Failed to add views",
		Error:   "blog not found",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.AddView(suite.ctx)
	suite.Equal(http.StatusNotFound, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGenerateContent_Success() {
	body := gin.H{
		"title": "some title",
		"tags":  []string{"tag1", "tag2"},
	}

	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs/generate",
		strings.NewReader(string(bodyJSON)),
	)

	content := "Generate content for this title :" + body["title"].(string) + " by considering this tags :" + strings.Join(body["tags"].([]string), ", ") + ". If the content contains any inappropriate content, please remove it, and state the reason for the removal."
	result := "some content that is generated by AI or LLM"

	suite.blogusecase.On("GenerateAiContent", content).Return(result, nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.GenerateContent(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGenerateContent_Fail() {
	body := gin.H{
		"title": "some title",
		"tags":  []string{"tag1", "tag2"},
	}

	bodyJSON, err := json.Marshal(body)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest(
		"POST",
		"/blogs/generate",
		strings.NewReader(string(bodyJSON)),
	)

	content := "Generate content for this title :" + body["title"].(string) + " by considering this tags :" + strings.Join(body["tags"].([]string), ", ") + ". If the content contains any inappropriate content, please remove it, and state the reason for the removal."

	suite.blogusecase.On("GenerateAiContent", content).Return("", errors.New("some error")).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Error:   "some error",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.GenerateContent(suite.ctx)
	suite.Equal(http.StatusInternalServerError, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGetTags_Success() {
	suite.ctx.Request = httptest.NewRequest("GET", "/blogs/tags", nil)

	tags := []*domain.Tag{
		{ID: "tag1"},
		{ID: "tag2"},
		{ID: "tag3"},
	}

	suite.blogusecase.On("GetTags").Return(tags, nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Tags retrieved",
		Data:    tags,
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.GetTags(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestGetTags_Fail() {
	suite.ctx.Request = httptest.NewRequest("GET", "/blogs/tags", nil)

	suite.blogusecase.On("GetTags").Return(nil, errors.New("some error")).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusInternalServerError,
		Message: "Error getting tags",
		Error:   "some error",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.GetTags(suite.ctx)
	suite.Equal(http.StatusInternalServerError, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestInsertTag_Success() {
	tag := domain.Tag{ID: "tag1"}

	body, err := json.Marshal(tag)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest("POST", "/blogs/tags", strings.NewReader(string(body)))

	claims := &domain.LoginClaims{Username: "user_1"}
	suite.ctx.Set("claims", claims)

	suite.blogusecase.On("InsertTag", &tag, claims).Return(nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Tag added",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.InsertTag(suite.ctx)
	suite.Equal(http.StatusCreated, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestInsertTag_Fail() {
	tag := domain.Tag{ID: "tag1"}

	body, err := json.Marshal(tag)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest("POST", "/blogs/tags", strings.NewReader(string(body)))

	claims := &domain.LoginClaims{Username: "user_1"}
	suite.ctx.Set("claims", claims)

	suite.blogusecase.On("InsertTag", &tag, claims).Return(errors.New("some error")).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusInternalServerError,
		Message: "Error adding tag",
		Error:   "some error",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.InsertTag(suite.ctx)
	suite.Equal(http.StatusInternalServerError, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestRemoveTags_Success() {
	tag := domain.Tag{ID: "tag1"}

	body, err := json.Marshal(tag)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest("DELETE", "/blogs/tags", strings.NewReader(string(body)))

	claims := &domain.LoginClaims{Username: "user_1"}
	suite.ctx.Set("claims", claims)

	suite.blogusecase.On("DeleteTag", &tag, claims).Return(nil).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Tag removed",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.RemoveTags(suite.ctx)
	suite.Equal(http.StatusOK, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func (suite *BlogControllerTestSuite) TestRemoveTags_Fail() {
	tag := domain.Tag{ID: "tag1"}

	body, err := json.Marshal(tag)
	suite.NoError(err)

	suite.ctx.Request = httptest.NewRequest("DELETE", "/blogs/tags", strings.NewReader(string(body)))

	claims := &domain.LoginClaims{Username: "user_1"}
	suite.ctx.Set("claims", claims)

	suite.blogusecase.On("DeleteTag", &tag, claims).Return(errors.New("some error")).Once()

	expectedBody := domain.APIResponse{
		Status:  http.StatusInternalServerError,
		Message: "Error removing tag",
		Error:   "some error",
	}

	expected, err := json.Marshal(expectedBody)
	suite.NoError(err)

	suite.blogcontroller.RemoveTags(suite.ctx)
	suite.Equal(http.StatusInternalServerError, suite.w.Code)
	suite.Equal(string(expected), suite.w.Body.String())
}

func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}
