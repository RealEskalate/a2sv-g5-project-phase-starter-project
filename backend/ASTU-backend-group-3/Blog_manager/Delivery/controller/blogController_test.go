package controller_test

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/controller"
	"ASTU-backend-group-3/Blog_manager/Domain"
	"ASTU-backend-group-3/Blog_manager/mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogControllerTestSuite struct {
	suite.Suite
	blogUsecase    *mocks.BlogUsecase
	blogController *controller.BlogController
	router         *gin.Engine
}

func (suite *BlogControllerTestSuite) SetupTest() {
	// Initialize the mocks and the controller
	suite.blogUsecase = new(mocks.BlogUsecase)
	suite.blogController = controller.NewBlogController(suite.blogUsecase)

	// Set up the router with the controller routes
	suite.router = gin.Default()
	suite.router.POST("/blogs", suite.blogController.CreateBlog)
	suite.router.GET("/blogs", suite.blogController.RetrieveBlogs)
	suite.router.DELETE("/blogs/:id", suite.blogController.DeleteBlogByID)
	suite.router.GET("/blogs/search", suite.blogController.SearchBlogs)
	suite.router.PUT("/blogs/:id", suite.blogController.UpdateBlog)
	suite.router.PUT("/blogs/:id/view", suite.blogController.IncrementViewCount)
	suite.router.PUT("/blogs/:id/like", suite.blogController.ToggleLike)
	suite.router.PUT("/blogs/:id/dislike", suite.blogController.ToggleDislike)
	suite.router.POST("/blogs/:id/comment", suite.blogController.AddComment)
	suite.router.GET("/blogs/filter", suite.blogController.FilterBlogs)
}

func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}

func (suite *BlogControllerTestSuite) TestCreateBlog() {

	suite.blogUsecase.On("CreateBlog", mock.Anything).Return(&Domain.Blog{
		Id:        "1",
		Title:     "Test Blog",
		Content:   "This is a test blog",
		Tags:      []string{"test", "blog"},
		Author:    "testuser",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/blogs", strings.NewReader(`{
		"title": "Test Blog",
		"content": "This is a test blog",
		"tags": ["test", "blog"]
	}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("username", "testuser")

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"data":`)
}

func (suite *BlogControllerTestSuite) TestCreateBlog_BadRequest() {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/blogs", strings.NewReader(`{
		"title": "Test Blog",
		"content": ""
	}`))
	req.Header.Set("Content-Type", "application/json")

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), `"error":`)
}

func (suite *BlogControllerTestSuite) TestRetrieveBlogs() {
	suite.blogUsecase.On("RetrieveBlogs", 1, 20, "date").Return([]Domain.Blog{
		{
			Id:        "1",
			Title:     "Test Blog 1",
			Content:   "This is the first test blog",
			Tags:      []string{"test"},
			Author:    "user1",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
		{
			Id:        "2",
			Title:     "Test Blog 2",
			Content:   "This is the second test blog",
			Tags:      []string{"blog"},
			Author:    "user2",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
	}, int64(2), nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs", nil)

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"data":`)
	suite.Contains(w.Body.String(), `"totalPages":`)
}

func (suite *BlogControllerTestSuite) TestDeleteBlogByID() {

	suite.blogUsecase.On("DeleteBlogByID", "123").Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "123"}}
	// Manually set the username in the Gin context
	c.Set("username", "adminuser")
	c.Set("role", "admin")
	suite.blogController.DeleteBlogByID(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"message":"Blog post deleted successfully"`)
}

func (suite *BlogControllerTestSuite) TestSearchBlogs() {

	expectedTitle := "Go Testing"
	expectedAuthor := "Hamza"
	expectedTags := []string{"testing", "go"}

	expectedBlogs := []Domain.Blog{
		{Title: expectedTitle, Author: expectedAuthor, Tags: expectedTags},
	}

	suite.blogUsecase.On("SearchBlogs", expectedTitle, expectedAuthor, expectedTags).Return(expectedBlogs, nil).Once()

	req, _ := http.NewRequest("GET", "/blogs/search?title=Go+Testing&author=Hamza&tags=testing&tags=go", nil)

	w := httptest.NewRecorder()

	// Run the request through the router
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"data":[`)
	suite.blogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestUpdateBlog() {

	blogID := "123"
	input := Domain.UpdateBlogInput{
		Title:   "Updated Title",
		Content: "Updated Content",
		Tags:    []string{"tag1", "tag2"},
	}
	updatedBlog := &Domain.Blog{
		Id:      blogID,
		Title:   "Updated Title",
		Content: "Updated Content",
		Author:  "Hamza",
		Tags:    []string{"tag1", "tag2"},
	}

	suite.blogUsecase.On("UpdateBlog", blogID, input, "authoruser").Return(updatedBlog, nil).Once()

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("PUT", "/blogs/"+blogID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: blogID}}
	c.Set("username", "authoruser")

	suite.blogController.UpdateBlog(c)

	// Verify the response
	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"data"`)
	// Verify that the mock expectations were met
	suite.blogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestIncrementViewCount() {

	blogID := "123"

	suite.blogUsecase.On("IncrementViewCount", blogID).Return(nil).Once()

	// Create a new HTTP PUT request to increment the view count
	req, _ := http.NewRequest("PUT", "/blogs/"+blogID+"/view", nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: blogID}}

	suite.blogController.IncrementViewCount(c)

	suite.Equal(http.StatusOK, w.Code)

	// Assert that the response contains the expected message
	suite.Contains(w.Body.String(), `"message":"View count updated"`)

	// Verify that the mock expectations were met
	suite.blogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestToggleLike() {

	blogID := "123"
	username := "Hamza"

	suite.blogUsecase.On("ToggleLike", blogID, username).Return(nil).Once()

	req, _ := http.NewRequest("PUT", "/blogs/"+blogID+"/like", nil)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: blogID}}

	c.Set("username", username)
	suite.blogController.ToggleLike(c)
	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"message":"Like toggled"`)
	suite.blogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestToggleDislike() {
	blogID := "123"
	username := "Hamza"

	suite.blogUsecase.On("ToggleDislike", blogID, username).Return(nil).Once()

	req, _ := http.NewRequest("PUT", "/blogs/"+blogID+"/dislike", nil)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: blogID}}

	c.Set("username", username)
	suite.blogController.ToggleDislike(c)
	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"message":"Dislike toggled"`)
	suite.blogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestAddComment() {
	blogID := "123"
	username := "hamza"
	content := "This is test comment"

	expectedComment := Domain.Comment{
		Content:   content,
		PostID:    blogID,
		UserID:    username,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	suite.blogUsecase.On("AddComment", blogID, mock.MatchedBy(func(c Domain.Comment) bool {
		return c.Content == expectedComment.Content &&
			c.PostID == expectedComment.PostID &&
			c.UserID == expectedComment.UserID &&
			c.CreatedAt == expectedComment.CreatedAt
	})).Return(nil).Once()

	reqBody := `{"content": "` + content + `"}`
	req, _ := http.NewRequest("POST", "/blogs/"+blogID+"/comment", bytes.NewReader([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: blogID}}
	c.Set("username", username)

	suite.blogController.AddComment(c)

	suite.Equal(http.StatusOK, w.Code)

	// Debug the exact response
	fmt.Println("Response Body:", w.Body.String())

	// Parse the JSON response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)

	// Check that the "message" field is present and has the correct value
	suite.Equal("Comment added", response["message"])

	suite.blogUsecase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestFilterBlogs_Success() {

	tags := []string{"tech", "science"}
	startDate := "2024-01-01T00:00:00Z"
	endDate := "2024-12-31T23:59:59Z"
	sortBy := "date"

	startDateParsed, _ := time.Parse(time.RFC3339, startDate)
	endDateParsed, _ := time.Parse(time.RFC3339, endDate)

	expectedBlogs := []Domain.Blog{
		{
			Id:        "1",
			Title:     "Tech Blog",
			Content:   "Content of the tech blog",
			Tags:      tags,
			Author:    "Author Name",
			CreatedAt: startDateParsed.Format(time.RFC3339),
			UpdatedAt: endDateParsed.Format(time.RFC3339),
			ViewCount: 123,
			Likes:     []string{"user1", "user2"},
			Dislikes:  []string{"user3"},
			Comments: []Domain.Comment{
				{
					Id:        "c1",
					Content:   "Great blog!",
					CreatedAt: "",
					PostID:    "",
					UserID:    "",
				},
			},
		},
	}

	suite.blogUsecase.On("FilterBlogs", tags, startDateParsed, endDateParsed, sortBy).
		Return(expectedBlogs, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/filter-blogs?tags=tech,science&startDate="+startDate+"&endDate="+endDate+"&sortBy="+sortBy, nil)
	suite.blogController.FilterBlogs(c)
	suite.Equal(http.StatusOK, w.Code)

	// Expected JSON response
	expectedResponse := `[
        {
            "id":"1",
            "title":"Tech Blog",
            "content":"Content of the tech blog",
            "tags":["tech","science"],
            "author":"Author Name",
            "created_at":"` + startDateParsed.Format(time.RFC3339) + `",
            "updated_at":"` + endDateParsed.Format(time.RFC3339) + `",
            "view_count":123,
            "likes":["user1","user2"],
            "dislikes":["user3"],
            "comments":[
                {
                    "id":"c1",
                    "content":"Great blog!",
                    "created_at":"",
                    "post_id":"",
                    "user_id":""
                }
            ]
        }
    ]`
	suite.JSONEq(expectedResponse, w.Body.String())
}
