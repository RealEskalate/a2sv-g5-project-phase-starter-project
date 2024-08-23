package controller_test

import (
	"backend-starter-project/delivery/controller"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockAIContentService struct {
	mock.Mock
}

func (m *MockAIContentService) GenerateContentSuggestions(keywords []string) (string, error) {
	args := m.Called(keywords)
	return args.String(0), args.Error(1)
}

func (m *MockAIContentService) SuggestContentImprovements(blogPostId, instruction string) (string, error) {
	args := m.Called(blogPostId, instruction)
	return args.String(0), args.Error(1)
}


type AIContentControllerTestSuite struct {
	suite.Suite
	mockAIContentService *MockAIContentService
	controller           *controller.AIContentController
	router               *gin.Engine
}

func (suite *AIContentControllerTestSuite) SetupTest() {
	suite.mockAIContentService = new(MockAIContentService)
	suite.controller = controller.NewAIContentController(suite.mockAIContentService)

	suite.router = gin.Default()
	suite.router.POST("/generate-suggestions", suite.controller.GenerateContentSuggestions)
	suite.router.POST("/suggest-improvements", suite.controller.SuggestContentImprovements)
}

func (suite *AIContentControllerTestSuite) TestGenerateContentSuggestions_Success() {
	keywords := []string{"AI", "content", "generation"}
	suggestion := "Here are some content suggestions."

	suite.mockAIContentService.On("GenerateContentSuggestions", keywords).Return(suggestion, nil)

	reqBody, _ := json.Marshal(map[string]interface{}{
		"keywords": keywords,
	})
	req, _ := http.NewRequest(http.MethodPost, "/generate-suggestions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), suggestion)
}

func (suite *AIContentControllerTestSuite) TestGenerateContentSuggestions_BadRequest() {
	req, _ := http.NewRequest(http.MethodPost, "/generate-suggestions", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid request payload")
}

func (suite *AIContentControllerTestSuite) TestSuggestContentImprovements_Success() {
	blogPostId := "12345"
	instruction := "Improve readability"
	suggestion := "Here are some improvements."

	suite.mockAIContentService.On("SuggestContentImprovements", blogPostId, instruction).Return(suggestion, nil)

	reqBody, _ := json.Marshal(map[string]interface{}{
		"blogPostId":  blogPostId,
		"instruction": instruction,
	})
	req, _ := http.NewRequest(http.MethodPost, "/suggest-improvements", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), suggestion)
}

func (suite *AIContentControllerTestSuite) TestSuggestContentImprovements_BadRequest() {
	req, _ := http.NewRequest(http.MethodPost, "/suggest-improvements", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid request payload")
}

func TestAIContentControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AIContentControllerTestSuite))
}
