package controller_test

import (
	"blog/delivery/controller"
	"blog/domain/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AIControllerSuite struct {
	suite.Suite
	router       *gin.Engine
	AIUsecase    *mocks.AIUsecase
	AIController *controller.AIController
}

func (suite *AIControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.AIUsecase = new(mocks.AIUsecase)
	suite.AIController = &controller.AIController{
		AIUsecase: suite.AIUsecase,
	}
	suite.router = gin.Default()
	suite.router.POST("/generate_content", suite.AIController.GenerateContent)
}
func (suite *AIControllerSuite) TearDownTest() {
	suite.AIUsecase.AssertExpectations(suite.T())
}

func (suite *AIControllerSuite) TestGenerateContent() {
	suite.Run("generate_content_success", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		var request struct {
			Prompt string `json:"keywords" binding:"required"`
		}
		request.Prompt = "test"
		resp := []genai.Part{}
		suite.AIUsecase.On("GenerateBlogContent", mock.Anything, request.Prompt).Return(resp, nil).Once()
		payload, _ := json.Marshal(request)
		req, _ := http.NewRequest(http.MethodPost, "/generate_content", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("generate_content_error", func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		var request struct {
			Keywords string `json:"keywords" binding:"required"`
		}
		request.Keywords = "test"
		suite.AIUsecase.On("GenerateBlogContent", mock.Anything, request.Keywords).Return(nil, errors.New("error")).Once()
		payload, _ := json.Marshal(request)
		req, _ := http.NewRequest(http.MethodPost, "/generate_content", bytes.NewBuffer(payload))
		c.Request = req
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusInternalServerError, w.Code)
	})
}

func TestAIControllerSuite(t *testing.T) {
	suite.Run(t, new(AIControllerSuite))
}
