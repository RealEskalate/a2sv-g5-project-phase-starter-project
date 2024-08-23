package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/delivery/controllers"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type BlogAssistantControllerTestSuite struct {
	suite.Suite
	usecase    *mocks.MockBlogAssistantUsecase
	controller domain.BlogAssistantController
}

func (suite *BlogAssistantControllerTestSuite) SetupSuite() {
	suite.usecase = new(mocks.MockBlogAssistantUsecase)
	suite.controller = controllers.NewBlogAssistantController(suite.usecase)
}

func TestBlogAssistantControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BlogAssistantControllerTestSuite))
}


func (suite *BlogAssistantControllerTestSuite) TestGenerateBlog() {
	suite.usecase.On("GenerateBlog", []string{"keyword"}, "tone", "audience").Return(map[string]interface{}{"key": "value"}, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/generate-blog", strings.NewReader(`{"keywords": ["keyword"], "tone": "tone", "audience": "audience"}`))

	suite.controller.GenerateBlog(ctx)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"key":"value"`)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *BlogAssistantControllerTestSuite) TestGenerateBlogWithEmptyKeywords() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/generate-blog", strings.NewReader(`{"keywords": [], "tone": "tone", "audience": "audience"}`))

	suite.controller.GenerateBlog(ctx)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "keywords are required")
}

func (suite *BlogAssistantControllerTestSuite) TestGenerateBlogWithInvalidJSON() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/generate-blog", strings.NewReader(`{}`))

	suite.controller.GenerateBlog(ctx)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *BlogAssistantControllerTestSuite) TestGenerateBlogWithUsecaseError() {
	suite.usecase.On("GenerateBlog", []string{"keyword"}, "tone", "audience").Return(nil, domain.CustomError{Code: http.StatusInternalServerError, Message: "internal server error"})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/generate-blog", strings.NewReader(`{"keywords": ["keyword"], "tone": "tone", "audience": "audience"}`))

	suite.controller.GenerateBlog(ctx)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "internal server error")
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *BlogAssistantControllerTestSuite) TestEnhanceBlog() {
	suite.usecase.On("EnhanceBlog", "content", "command").Return(map[string]interface{}{"key": "value"}, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/enhance-blog", strings.NewReader(`{"content": "content", "command": "command"}`))

	suite.controller.EnhanceBlog(ctx)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"key":"value"`)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *BlogAssistantControllerTestSuite) TestEnhanceBlogWithInvalidJSON() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/enhance-blog", strings.NewReader(`{}`))

	suite.controller.EnhanceBlog(ctx)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *BlogAssistantControllerTestSuite) TestEnhanceBlogWithUsecaseError() {
	suite.usecase.On("EnhanceBlog", "content", "command").Return(nil, domain.CustomError{Code: http.StatusInternalServerError, Message: "internal server error"})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/enhance-blog", strings.NewReader(`{"content": "content", "command": "command"}`))

	suite.controller.EnhanceBlog(ctx)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "internal server error")
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *BlogAssistantControllerTestSuite) TestSuggestBlog() {
	suite.usecase.On("SuggestBlog", "tech").Return([]map[string]interface{}{{"key": "value"}, {"key2": "value2"}}, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/suggest-blog", strings.NewReader(`{"industry": "tech"}`))

	suite.controller.SuggestBlog(ctx)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), `"key":"value"`)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *BlogAssistantControllerTestSuite) TestSuggestBlogWithInvalidJSON() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/suggest-blog", strings.NewReader(`{}`))

	suite.controller.SuggestBlog(ctx)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *BlogAssistantControllerTestSuite) TestSuggestBlogWithInvalidRequest() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/suggest-blog", strings.NewReader(`{"industry": ""}`))

	suite.controller.SuggestBlog(ctx)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "niche is required")
}

func (suite *BlogAssistantControllerTestSuite) TestSuggestBlogWithUsecaseError() {
	suite.usecase.On("SuggestBlog", "niche").Return(nil, domain.CustomError{Code: http.StatusInternalServerError, Message: "internal server error"})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("POST", "/suggest-blog", strings.NewReader(`{"industry": "niche"}`))

	suite.controller.SuggestBlog(ctx)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "internal server error")
	suite.usecase.AssertExpectations(suite.T())
}
