package controllers

import (
	domain "blogs/Domain"
	"blogs/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ControllerTestSuite struct {
	suite.Suite
	mockBlogUseCase *mocks.BlogUsecase
	mockValidator   *mocks.ValidateInterface
	router          *gin.Engine
	controller      *BlogController
}

func MockAuthMiddleware(role string, userID string, user_name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("role", role)
		c.Set("user_id", userID)
		c.Set("user_name", user_name)
		c.Next()
	}
}

func (suite *ControllerTestSuite) SetupTest() {
	suite.mockBlogUseCase = new(mocks.BlogUsecase)
	suite.mockValidator = new(mocks.ValidateInterface)
	suite.router = gin.Default()
	controller := NewBlogController(suite.mockBlogUseCase, suite.mockValidator)
	suite.controller = &controller
}
func (suite *ControllerTestSuite) TestGetBlogsHandler_Success() {
	mockBlogs := []domain.Blog{
		{
			ID:      primitive.NewObjectID(),
			Title:   "Blog 1",
			Author:  "The author",
			Content: "This is a blog",
			Tags:    []string{"tag1", "tag2"},
		},
		{
			ID:      primitive.NewObjectID(),
			Title:   "Blog 1",
			Author:  "The author",
			Content: "This is a blog",
			Tags:    []string{"tag1", "tag2"},
		},
	}
	pagination := domain.Pagination{
		CurrentPage: 0,
		PageSize:    1,
		TotalPages:  2,
		TotatRecord: 2,
	}

	suite.router.GET("/blogs", suite.controller.GetBlogs)

	suite.mockBlogUseCase.On("GetBlogs", "0", "0", "").Return(mockBlogs, pagination, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Blog 1")
	suite.Contains(w.Body.String(), "The author")
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetBlogByIdHandler_Success() {
	blog_id := primitive.NewObjectID()
	mockBlog := domain.Blog{
		ID:      blog_id,
		Title:   "Blog 1",
		Author:  "The author",
		Content: "This is a blog",
		Tags:    []string{"tag1", "tag2"},
	}
	suite.mockBlogUseCase.On("GetBlogByID", mockBlog.ID.Hex(), false).Return(mockBlog, nil)
	suite.router.GET("/blogs/:id", suite.controller.GetBlogByID)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs/"+mockBlog.ID.Hex(), nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Blog 1")
	suite.Contains(w.Body.String(), "The author")
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestReactOnBlogHandler_Success() {
	blog_id := primitive.NewObjectID().Hex()
	user_id := primitive.NewObjectID().Hex()
	suite.mockBlogUseCase.On("ReactOnBlog", user_id, "true", blog_id).Return(domain.ErrorResponse{})
	suite.router.POST("/blogs/react/:id", MockAuthMiddleware("user", user_id, "user123"), suite.controller.ReactOnBlog)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/blogs/react/"+blog_id+"?isLiked=true", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Reaction saved successfully")
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestCommentOnBlogHandler_Success() {
	comment := domain.Comment{
		Blog_ID: primitive.NewObjectID(),
		Content: "hi very good blog",
	}
	user_id := primitive.NewObjectID().Hex()
	suite.mockBlogUseCase.On("CommentOnBlog", user_id, "user123", comment).Return(nil)
	suite.mockValidator.On("ValidateStruct", comment).Return(nil)
	suite.router.POST("/blogs/comment/create", MockAuthMiddleware("user", user_id, "user123"), suite.controller.CommentOnBlog)

	body, _ := json.Marshal(comment)
	req, _ := http.NewRequest("POST", "/blogs/comment/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusCreated, w.Code)
	suite.Contains(w.Body.String(), "Comment created successfully")
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestCreateBlogHandler_Success() {
	blog_id := primitive.NewObjectID()
	created_at := time.Now()
	updated_at := time.Now()
	mockBlog := domain.Blog{
		ID:      blog_id,
		Title:   "Blog 1",
		Author:  "The author",
		Content: "This is a blog",
		Tags:    []string{"tag1", "tag2"},
	}
	newBlog := domain.Blog{
		ID:            blog_id,
		Title:         "Blog 1",
		Author:        "The author",
		Content:       "This is a blog",
		Tags:          []string{},
		CreatedAt:     created_at,
		UpdatedAt:     updated_at,
		Commenters_ID: []primitive.ObjectID{},
		ViewCount:     0,
		Popularity:    0,
	}
	user_id := primitive.NewObjectID().Hex()
	suite.mockBlogUseCase.On("CreateBlog", user_id, mockBlog, "user").Return(newBlog, nil)
	suite.mockValidator.On("ValidateStruct", mockBlog).Return(nil)
	suite.router.POST("/blogs/create", MockAuthMiddleware("user", user_id, "user123"), suite.controller.CreateBlog)

	body, _ := json.Marshal(mockBlog)
	req, _ := http.NewRequest("POST", "/blogs/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusCreated, w.Code)
	suite.Contains(w.Body.String(), "Blog created successfully")
	suite.Contains(w.Body.String(), blog_id.Hex())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestDeleteHandler_Success() {
	user_id := primitive.NewObjectID().Hex()
	blog_id := primitive.NewObjectID().Hex()
	suite.mockBlogUseCase.On("DeleteBlogByID", user_id, blog_id, "user").Return(domain.ErrorResponse{})
	suite.router.DELETE("/blogs/delete:id", MockAuthMiddleware("user", user_id, "user123"), suite.controller.DeleteBlogByID)
	req, _ := http.NewRequest("DELETE", "/blogs/delete"+blog_id, nil)

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Blog deleted successfully")
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetMyBlogByIdHandler_Success() {
	blog_id := primitive.NewObjectID()
	user_id := primitive.NewObjectID().Hex()
	mockBlog := domain.Blog{
		ID:      blog_id,
		Title:   "Blog 1",
		Author:  "The author",
		Content: "This is a blog",
		Tags:    []string{"tag1", "tag2"},
	}
	suite.mockBlogUseCase.On("GetMyBlogByID", user_id, blog_id.Hex(), "user").Return(mockBlog, nil)
	suite.router.GET("/blogs/my/:id", MockAuthMiddleware("user", user_id, "user123"), suite.controller.GetMyBlogByID)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs/my/"+blog_id.Hex(), nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Blog 1")
	suite.Contains(w.Body.String(), "The author")
	suite.mockBlogUseCase.AssertExpectations(suite.T())

}
func (suite *ControllerTestSuite) TestGetMyBlogsHandler_Success() {
	user_id := primitive.NewObjectID()
	mockBlogs := []domain.Blog{
		{
			ID:      primitive.NewObjectID(),
			Title:   "Blog 1",
			Author:  "The author",
			Content: "This is a blog",
			Tags:    []string{"tag1", "tag2"},
		},
		{
			ID:      primitive.NewObjectID(),
			Title:   "Blog 1",
			Author:  "The author",
			Content: "This is a blog",
			Tags:    []string{"tag1", "tag2"},
		},
	}
	pagination := domain.Pagination{
		CurrentPage: 0,
		PageSize:    1,
		TotalPages:  2,
		TotatRecord: 2,
	}

	suite.router.GET("/blogs/my", MockAuthMiddleware("user", user_id.Hex(), "user123"), suite.controller.GetMyBlogs)

	suite.mockBlogUseCase.On("GetMyBlogs", user_id.Hex(), "0", "0", "").Return(mockBlogs, pagination, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs/my", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Blog 1")
	suite.Contains(w.Body.String(), "The author")
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestUpdateTaskHandler_UserSuccess() {
	blog_id := primitive.NewObjectID()
	user_id := primitive.NewObjectID()
	mockBlog := domain.Blog{
		ID:      blog_id,
		Title:   "Blog 1",
		Author:  "The author",
		Content: "This is a blog",
		Tags:    []string{"tag1", "tag2"},
	}

	suite.router.PUT("/update/:id", MockAuthMiddleware("user", user_id.Hex(), "user123"), suite.controller.UpdateBlogByID)
	suite.mockBlogUseCase.On("UpdateBlogByID", user_id.Hex(), blog_id.Hex(), mockBlog, "user").Return(mockBlog, nil)

	body, _ := json.Marshal(mockBlog)
	req, _ := http.NewRequest("PUT", "/update/"+blog_id.Hex(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusAccepted, w.Code)
	suite.Contains(w.Body.String(), blog_id.Hex())
	suite.Contains(w.Body.String(), "This is a blog")
	suite.Contains(w.Body.String(), "Blog 1")
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestSearchBlogByTitleAndAuthorHandler_Success() {
	mockBlogs := []domain.Blog{
		{
			ID:      primitive.NewObjectID(),
			Title:   "Blog 1",
			Author:  "The author",
			Content: "This is a blog",
			Tags:    []string{"tag1", "tag2"},
		},
		{
			ID:      primitive.NewObjectID(),
			Title:   "Blog 1",
			Author:  "The author",
			Content: "This is a blog",
			Tags:    []string{"tag1", "tag2"},
		},
	}
	pagination := domain.Pagination{
		CurrentPage: 0,
		PageSize:    1,
		TotalPages:  2,
		TotatRecord: 2,
	}

	suite.router.GET("/search", suite.controller.SearchBlogByTitleAndAuthor)

	suite.mockBlogUseCase.On("SearchBlogByTitleAndAuthor", "xx", "xx", "", "", "").Return(mockBlogs, pagination, domain.ErrorResponse{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?title=xx&author=xx", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "Blogs fetched successfully.")
	suite.Contains(w.Body.String(), mockBlogs[0].Title)
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
