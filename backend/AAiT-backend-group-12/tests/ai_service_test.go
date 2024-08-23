package tests

import (
	ai_service "blog_api/infrastructure/ai"
	"blog_api/mocks"
	"context"
	"fmt"
	"testing"

	"github.com/google/generative-ai-go/genai"
	"github.com/stretchr/testify/suite"
)

const (
	ERROR_INPUT = "error-genesis"
)

type MockModel genai.GenerativeModel

func (model *MockModel) GenerateContent(c context.Context, g genai.Text) (*genai.GenerateContentResponse, error) {
	if g == ERROR_INPUT {
		return nil, fmt.Errorf("error")
	}

	return nil, nil
}

type AIServicesTestSuite struct {
	suite.Suite
	model     *mocks.AIModelInterface
	aiService *ai_service.AIService
}

func (suite *AIServicesTestSuite) SetupSuite() {
	suite.model = new(mocks.AIModelInterface)
	suite.aiService = &ai_service.AIService{
		Model: suite.model,
		Ctx:   context.Background(),
	}
}

func (suite *AIServicesTestSuite) TestNewAIService() {

}

func TestAIServices(t *testing.T) {
	suite.Run(t, new(AIServicesTestSuite))
}
