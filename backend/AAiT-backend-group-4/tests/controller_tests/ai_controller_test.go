package controller_tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "aait-backend-group4/Domain"
    "aait-backend-group4/Delivery/controllers"
    "aait-backend-group4/Mocks/usecase_mocks"

)

type AIControllerTestSuite struct {
    suite.Suite
    Router *gin.Engine
    MockUsecase *mocks.MockAiUsecase
    Controller *controllers.AIController
}

func (suite *AIControllerTestSuite) SetupSuite() {
    // Set Gin to Test mode
    gin.SetMode(gin.TestMode)

    // Initialize router and controller
    suite.Router = gin.Default()
    suite.MockUsecase = &mocks.MockAiUsecase{}
    suite.Controller = &controllers.AIController{AiUsecase: suite.MockUsecase}

    // Set up routes
    suite.Router.POST("/blog/generateWithTags", suite.Controller.GenerateTextWithTags)
    suite.Router.POST("/blog/generateWithPrompt", suite.Controller.GenerateTextWithPrompt)
    suite.Router.POST("/blog/generateSuggestions", suite.Controller.GenerateSuggestions)
}

func (suite *AIControllerTestSuite) TestGenerateTextWithTags() {
    tests := []struct {
        name         string
        requestBody  interface{}
        expectedCode int
        expectedBody string
    }{
        {
            name:         "Valid tags",
            requestBody:  []domain.Tag{{Name: "tag1"}, {Name: "tag2"}},
            expectedCode: http.StatusOK,
            expectedBody: `{"result":"Mocked blog content"}`,
        },
        {
            name:         "Empty tags",
            requestBody:  []domain.Tag{},
            expectedCode: http.StatusBadRequest,
            expectedBody: `{"error":"no tags provided"}`,
        },
    }

    for _, tt := range tests {
        suite.Run(tt.name, func() {
            body, _ := json.Marshal(tt.requestBody)
            req, _ := http.NewRequest(http.MethodPost, "/blog/generateWithTags", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")

            w := httptest.NewRecorder()
            suite.Router.ServeHTTP(w, req)

            assert.Equal(suite.T(), tt.expectedCode, w.Code)
            assert.JSONEq(suite.T(), tt.expectedBody, w.Body.String())
        })
    }
}

func (suite *AIControllerTestSuite) TestGenerateTextWithPrompt() {
    tests := []struct {
        name         string
        requestBody  interface{}
        expectedCode int
        expectedBody string
    }{
        {
            name:         "Valid prompt",
            requestBody:  map[string]string{"prompt": "Some prompt"},
            expectedCode: http.StatusOK,
            expectedBody: `{"result":"Mocked text content"}`,
        },
        {
            name:         "Empty prompt",
            requestBody:  map[string]string{"prompt": ""},
            expectedCode: http.StatusBadRequest,
            expectedBody: `{"error":"empty prompt"}`,
        },
    }

    for _, tt := range tests {
        suite.Run(tt.name, func() {
            body, _ := json.Marshal(tt.requestBody)
            req, _ := http.NewRequest(http.MethodPost, "/blog/generateWithPrompt", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")

            w := httptest.NewRecorder()
            suite.Router.ServeHTTP(w, req)

            assert.Equal(suite.T(), tt.expectedCode, w.Code)
            assert.JSONEq(suite.T(), tt.expectedBody, w.Body.String())
        })
    }
}

func (suite *AIControllerTestSuite) TestGenerateSuggestions() {
    tests := []struct {
        name         string
        requestBody  interface{}
        expectedCode int
        expectedBody string
    }{
        {
            name:         "Valid text content",
            requestBody:  map[string]string{"textContent": "Sample content"},
            expectedCode: http.StatusOK,
            expectedBody: `{"suggestions":"Mocked suggestions"}`,
        },
        {
            name:         "Empty text content",
            requestBody:  map[string]string{"textContent": ""},
            expectedCode: http.StatusBadRequest,
            expectedBody: `{"error":"empty text content"}`,
        },
    }

    for _, tt := range tests {
        suite.Run(tt.name, func() {
            body, _ := json.Marshal(tt.requestBody)
            req, _ := http.NewRequest(http.MethodPost, "/blog/generateSuggestions", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")

            w := httptest.NewRecorder()
            suite.Router.ServeHTTP(w, req)

            assert.Equal(suite.T(), tt.expectedCode, w.Code)
            assert.JSONEq(suite.T(), tt.expectedBody, w.Body.String())
        })
    }
}

func TestAIControllerTestSuite(t *testing.T) {
    suite.Run(t, new(AIControllerTestSuite))
}
