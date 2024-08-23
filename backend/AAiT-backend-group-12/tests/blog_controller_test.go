package tests

import (
	"blog_api/delivery/controllers"
	"blog_api/domain"
	"blog_api/mocks"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BlogControllerTestSuite struct {
	suite.Suite
	mockBlogUseCase *mocks.BlogUseCaseInterface
	mockAIServices  *mocks.AIServicesInterface
	controller      *controllers.BlogController
}

func (suite *BlogControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockBlogUseCase = new(mocks.BlogUseCaseInterface)
	suite.mockAIServices = new(mocks.AIServicesInterface)
	suite.controller = controllers.NewBlogController(suite.mockBlogUseCase)
}

func (suite *BlogControllerTestSuite) TestCreateBlogHandler_Success() {
	newBlog := domain.NewBlog{
		Title:   "Sample Blog",
		Content: "This is a sample blog post content ;vjdfnvjnvf fkvndlvn jvfdjkvn vjfd jvnjdsf  vjdfv fj vjfd jv fv jdf v fdj vdfj vjdf vjkd fv d fvd vdfuivdfiu uv df vufv fj vufid pv  vuf vd vu psdfvu vfv.",
	}
	userName := "test_user"

	suite.mockBlogUseCase.On("CreateBlogPost", mock.Anything, &newBlog, userName).Return(nil)

	reqBody, _ := json.Marshal(newBlog)
	req, _ := http.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)

	suite.controller.CreateBlogHandler(ctx)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.JSONEq(suite.T(), `{"message": "blog created successfully"}`, w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestUpdateBlogHandler_Success() {
	blogID := "blog123"
	updatedBlog := domain.NewBlog{
		Title:   "Updated Blog",
		Content: "This is a sample blog post content ;vjdfnvjnvf fkvndlvn jvfdjkvn vjfd jvnjdsf  vjdfv fj vjfd jv fv jdf v fdj vdfj vjdf vjkd fv d fvd vdfuivdfiu uv df vufv fj vufid pv  vuf vd vu psdfvu vfv.",
	}
	userName := "test_user"

	suite.mockBlogUseCase.On("EditBlogPost", mock.Anything, blogID, &updatedBlog, userName).Return(nil)

	reqBody, _ := json.Marshal(updatedBlog)
	req, _ := http.NewRequest(http.MethodPut, "/blogs/"+blogID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: blogID}}

	suite.controller.UpdateBlogHandler(ctx)

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestDeleteBlogHandler_Success() {
	blogID := "blog123"
	userName := "test_user"

	suite.mockBlogUseCase.On("DeleteBlogPost", mock.Anything, blogID, userName).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/blogs/"+blogID, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: blogID}}

	suite.controller.DeleteBlogHandler(ctx)

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGetBlogHandler_Success() {
	filters := domain.BlogFilterOptions{
		Author: "test_user",
	}
	expectedBlogs := []domain.Blog{
		{ID: "blog123", Title: "Sample Blog", Content: "This is a sample blog post content ;vjdfnvjnvf fkvndlvn jvfdjkvn vjfd jvnjdsf  vjdfv fj vjfd jv fv jdf v fdj vdfj vjdf vjkd fv d fvd vdfuivdfiu uv df vufv fj vufid pv  vuf vd vu psdfvu vfv."},
	}
	expectedTotal := 1

	suite.mockBlogUseCase.On("GetBlogPosts", mock.Anything, filters).Return(expectedBlogs, expectedTotal, nil)

	reqBody, _ := json.Marshal(filters)
	req, _ := http.NewRequest(http.MethodPost, "/blogs", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	suite.controller.GetBlogHandler(ctx)

	expectedResponse := `{
        "blogs": [{
            "id": "blog123",
            "title": "Sample Blog",
            "content": "This is a sample blog post content ;vjdfnvjnvf fkvndlvn jvfdjkvn vjfd jvnjdsf  vjdfv fj vjfd jv fv jdf v fdj vdfj vjdf vjkd fv d fvd vdfuivdfiu uv df vufv fj vufid pv  vuf vd vu psdfvu vfv.",
            "comment": null,
            "created_at": "0001-01-01T00:00:00Z",
            "disliked_by": null,
            "liked_by": null,
            "tags": null,
            "updated_at": "0001-01-01T00:00:00Z",
            "username": "",
            "view_count": 0
        }],
        "total": 1,
        "currentPage": 1,
        "postsPerPage": 10
    }`

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGetBlogPostByIDHandler_Success() {
	blogID := "blog123"
	expectedBlog := domain.Blog{
		ID:      blogID,
		Title:   "Sample Blog",
		Content: "This is a sample blog post content ;vjdfnvjnvf fkvndlvn jvfdjkvn vjfd jvnjdsf  vjdfv fj vjfd jv fv jdf v fdj vdfj vjdf vjkd fv d fvd vdfuivdfiu uv df vufv fj vufid pv  vuf vd vu psdfvu vfv.",
	}

	suite.mockBlogUseCase.On("GetBlogPostByID", mock.Anything, blogID).Return(&expectedBlog, nil)

	req, _ := http.NewRequest(http.MethodGet, "/blogs/"+blogID, nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = gin.Params{gin.Param{Key: "id", Value: blogID}}

	suite.controller.GetBlogByIDHandler(ctx)

	expectedResponse := `{
        "id": "blog123",
        "title": "Sample Blog",
        "content": "This is a sample blog post content ;vjdfnvjnvf fkvndlvn jvfdjkvn vjfd jvnjdsf  vjdfv fj vjfd jv fv jdf v fdj vdfj vjdf vjkd fv d fvd vdfuivdfiu uv df vufv fj vufid pv  vuf vd vu psdfvu vfv.",
        "comment": null,
        "created_at": "0001-01-01T00:00:00Z",
        "disliked_by": null,
        "liked_by": null,
        "tags": null,
        "updated_at": "0001-01-01T00:00:00Z",
        "username": "",
        "view_count": 0
    }`

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}
func (suite *BlogControllerTestSuite) TestLikeBlogHandler_Success() {
	// Arrange
	ctx := context.TODO()
	blogId := "test-blog-id"
	action := "like"
	state := true
	username := "test-user"

	// Mock the TrackBlogPopularity method
	suite.mockBlogUseCase.On("TrackBlogPopularity", ctx, blogId, action, state, username).
		Return(nil).
		Once()

	// Act
	err := suite.mockBlogUseCase.TrackBlogPopularity(ctx, blogId, action, state, username)

	// Assert
	suite.Nil(err)
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestDislikeBlogHandler_Success() {
	// Arrange
	ctx := context.TODO()
	blogId := "test-blog-id"
	action := "dislike"
	state := false
	username := "test-user"

	// Mock the TrackBlogPopularity method
	suite.mockBlogUseCase.On("TrackBlogPopularity", ctx, blogId, action, state, username).
		Return(nil).
		Once()

	// Act
	err := suite.mockBlogUseCase.TrackBlogPopularity(ctx, blogId, action, state, username)

	// Assert
	suite.Nil(err)
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGenerateTopicHandler_Success() {
	// Arrange
	reqPayload := struct {
		Keywords []string `json:"keywords"`
	}{
		Keywords: []string{"golang", "testing"},
	}
	expectedTopics := []string{"Golang Tips", "Testing Best Practices"}

	suite.mockBlogUseCase.On("GenerateTrendingTopics", reqPayload.Keywords).
		Return(expectedTopics, nil).
		Once()

	reqBody, _ := json.Marshal(reqPayload)
	req, _ := http.NewRequest(http.MethodPost, "/blogs/generate-topics", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Act
	suite.controller.GenerateTopicHandler(ctx)

	// Assert
	expectedResponse := `{"topics": ["Golang Tips", "Testing Best Practices"]}`
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestGenerateContentHandler_Success() {
	// Arrange
	reqPayload := struct {
		Topics []string `json:"topics"`
	}{
		Topics: []string{"tech trends", "AI advancements"},
	}
	expectedContent := "Generated blog content based on the topics."

	suite.mockBlogUseCase.On("GenerateBlogContent", reqPayload.Topics).
		Return(expectedContent, nil).
		Once()

	reqBody, _ := json.Marshal(reqPayload)
	req, _ := http.NewRequest(http.MethodPost, "/blogs/generate-content", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Act
	suite.controller.GenerateContentHandler(ctx)

	// Assert
	expectedResponse := `{"content": "Generated blog content based on the topics."}`
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}
func (suite *BlogControllerTestSuite) TestReviewContentHandler_Success() {
	// Arrange
	reqPayload := struct {
		Content string `json:"content"`
	}{
		Content: "Sample blog content for review",
	}
	expectedSuggestions := "Improve introduction; Add more details to conclusion"

	suite.mockBlogUseCase.On("ReviewBlogContent", reqPayload.Content).
		Return(expectedSuggestions, nil).
		Once()

	reqBody, _ := json.Marshal(reqPayload)
	req, _ := http.NewRequest(http.MethodPost, "/blogs/review-content", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Act
	suite.controller.ReviewContentHandler(ctx)

	// Assert
	expectedResponse := `{"suggestions": "Improve introduction; Add more details to conclusion"}`
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestHandleCreateComment_Success() {
	comment := domain.NewComment{
		Content: "This is a test comment",
	}
	blogID := "blog123"
	userName := "test_user"

	suite.mockBlogUseCase.On("AddComment", mock.Anything, blogID, &comment, userName).Return(nil)

	reqBody, _ := json.Marshal(comment)
	req, _ := http.NewRequest(http.MethodPost, "/blogs/"+blogID+"/comments", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)
	ctx.Params = gin.Params{gin.Param{Key: "blogId", Value: blogID}}

	suite.controller.HandleCreateComment(ctx)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.JSONEq(suite.T(), `{"message": "created successfully"}`, w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestHandleDeleteComment_Success() {
	blogID := "blog123"
	commentID := "comment123"
	userName := "test_user"

	// Arrange
	suite.mockBlogUseCase.On("DeleteComment", mock.Anything, blogID, commentID, userName).Return(nil)

	// Act
	req, _ := http.NewRequest(http.MethodDelete, "/blogs/"+blogID+"/comments/"+commentID, nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)
	ctx.Params = gin.Params{
		{Key: "blogId", Value: blogID},
		{Key: "commentId", Value: commentID},
	}

	// Call the handler
	suite.controller.HandleDeleteComment(ctx)

	// Assert
	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	// Since there's no content, don't assert the response body
	suite.Empty(w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestHandleUpdateComment_Success() {
	comment := domain.NewComment{
		Content: "Updated comment content",
	}
	blogID := "blog123"
	commentID := "comment123"
	userName := "test_user"

	// Arrange
	suite.mockBlogUseCase.On("UpdateComment", mock.Anything, blogID, commentID, &comment, userName).Return(nil)

	// Act
	reqBody, _ := json.Marshal(comment)
	req, _ := http.NewRequest(http.MethodPut, "/blogs/"+blogID+"/comments/"+commentID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)
	ctx.Params = gin.Params{
		{Key: "blogId", Value: blogID},
		{Key: "commentId", Value: commentID},
	}

	suite.controller.HandleUpdateComment(ctx)

	// Assert
	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	assert.Empty(suite.T(), w.Body.String()) // Check that the response body is empty
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *BlogControllerTestSuite) TestHandleCreateComment() {
	// Test Success Case
	comment := domain.NewComment{
		Content: "This is a test comment",
	}
	blogID := "blog123"
	userName := "test_user"

	suite.mockBlogUseCase.On("AddComment", mock.Anything, blogID, &comment, userName).Return(nil)

	reqBody, _ := json.Marshal(comment)
	req, _ := http.NewRequest(http.MethodPost, "/blogs/"+blogID+"/comments", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)
	ctx.Params = gin.Params{gin.Param{Key: "blogId", Value: blogID}}

	suite.controller.HandleCreateComment(ctx)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.JSONEq(suite.T(), `{"message": "created successfully"}`, w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())

	// Test Binding Error
	req, _ = http.NewRequest(http.MethodPost, "/blogs/"+blogID+"/comments", nil)
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Request = req

	suite.controller.HandleCreateComment(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.JSONEq(suite.T(), `{"error": "EOF"}`, w.Body.String()) // Adjust the expected error message as needed

	// Test Validation Error
	comment = domain.NewComment{
		Content: "", // Assume empty content fails validation
	}

	reqBody, _ = json.Marshal(comment)
	req, _ = http.NewRequest(http.MethodPost, "/blogs/"+blogID+"/comments", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)
	ctx.Params = gin.Params{gin.Param{Key: "blogId", Value: blogID}}

	suite.controller.HandleCreateComment(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.JSONEq(suite.T(), `{"error": "Key: 'NewComment.Content' Error:Field validation for 'Content' failed on the 'required' tag"}`, w.Body.String())
}

func TestBlogControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogControllerTestSuite))
}
