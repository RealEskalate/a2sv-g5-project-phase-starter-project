package tests

import (
	"blog_api/domain"
	"blog_api/delivery/controllers"
	"blog_api/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CommentControllerTestSuite struct {
	suite.Suite
	mockBlogUseCase *mocks.BlogUseCaseInterface
	controller      *controllers.BlogController
}

func (suite *CommentControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockBlogUseCase = new(mocks.BlogUseCaseInterface)
	suite.controller = controllers.NewBlogController(suite.mockBlogUseCase)
}

func (suite *CommentControllerTestSuite) TestHandleCreateComment_Success() {
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

func (suite *CommentControllerTestSuite) TestHandleDeleteComment_Success() {
	blogID := "blog123"
	commentID := "comment123"
	userName := "test_user"

	suite.mockBlogUseCase.On("DeleteComment", mock.Anything, blogID, commentID, userName).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/blogs/"+blogID+"/comments/"+commentID, nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("username", userName)
	ctx.Params = gin.Params{
		{Key: "blogId", Value: blogID},
		{Key: "commentId", Value: commentID},
	}

	suite.controller.HandleDeleteComment(ctx)

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	suite.Empty(w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func (suite *CommentControllerTestSuite) TestHandleUpdateComment_Success() {
	comment := domain.NewComment{
		Content: "Updated comment content",
	}
	blogID := "blog123"
	commentID := "comment123"
	userName := "test_user"

	suite.mockBlogUseCase.On("UpdateComment", mock.Anything, blogID, commentID, &comment, userName).Return(nil)

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

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	assert.Empty(suite.T(), w.Body.String())
	suite.mockBlogUseCase.AssertExpectations(suite.T())
}

func TestCommentControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CommentControllerTestSuite))
}
