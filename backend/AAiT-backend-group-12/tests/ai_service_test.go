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
	mockModel *mocks.AIModelInterface
	aiService *ai_service.AIService
}

func (suite *AIServicesTestSuite) SetupSuite() {
	suite.mockModel = new(mocks.AIModelInterface)
	suite.aiService = &ai_service.AIService{
		Model: suite.mockModel,
		Ctx:   context.Background(),
	}
}

func (suite *AIServicesTestSuite) TestGenerateContent_Negative_NoCandidates() {
	suite.mockModel.On("GenerateContent", suite.aiService.Ctx, genai.Text("Generate a blog post about test. The content should be engaging, include relevant subheadings, and provide useful insights. Return the content in a well-structured format.")).Return(&genai.GenerateContentResponse{}, nil).Once()

	res, err := suite.aiService.GenerateContent([]string{"test"})
	suite.Nil(err)
	suite.Equal("No candidates found", res)
}

func (suite *AIServicesTestSuite) TestGenerateContent_Negative_NoContentParts() {
	item := &genai.Candidate{
		Content: &genai.Content{
			Parts: []genai.Part{},
		},
	}

	suite.mockModel.On("GenerateContent", suite.aiService.Ctx, genai.Text("Generate a blog post about test. The content should be engaging, include relevant subheadings, and provide useful insights. Return the content in a well-structured format.")).Return(&genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			item,
		},
	}, nil).Once()

	res, err := suite.aiService.GenerateContent([]string{"test"})
	suite.Nil(err)
	suite.Equal("No content parts found", res)
}

func (suite *AIServicesTestSuite) TestCleanText() {
	input := "T***h*i**s *is* s*ome te********xt with unw******ante*d *ch*aracters**"
	expectedOutput := "This is some text with unwanted characters"

	output := suite.aiService.CleanText(input)
	suite.Equal(expectedOutput, output)
}

func (suite *AIServicesTestSuite) TestExtractText_Struct() {
	type TestData struct {
		Text string
	}

	input := TestData{
		Text: "This is some text",
	}
	expectedOutput := "This is some text"

	output := suite.aiService.ExtractText(input)
	suite.Equal(expectedOutput, output)
}

func (suite *AIServicesTestSuite) TestExtractText_Direct() {
	input2 := "This is another text"
	expectedOutput2 := "This is another text"

	output2 := suite.aiService.ExtractText(input2)
	suite.Equal(expectedOutput2, output2)
}

func (suite *AIServicesTestSuite) TestExtractText_Negative() {
	input2 := 0

	output2 := suite.aiService.ExtractText(input2)
	suite.Equal("", output2)
}

func TestAIServices(t *testing.T) {
	suite.Run(t, new(AIServicesTestSuite))
}
