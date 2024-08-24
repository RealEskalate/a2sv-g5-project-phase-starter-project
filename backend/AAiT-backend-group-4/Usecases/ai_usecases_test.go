package usecases_test

import (
	domain "aait-backend-group4/Domain"
	usecases "aait-backend-group4/Usecases"
	"aait-backend-group4/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AiUsecaseTestSuite struct {
	suite.Suite
	mockAiRepo *mocks.AiRepository
	aiUsecase  domain.AiUsecase
}

func (suite *AiUsecaseTestSuite) SetupTest() {
	suite.mockAiRepo = new(mocks.AiRepository)
	suite.aiUsecase = usecases.NewAiUsecase(2*time.Second, suite.mockAiRepo)
}

func (suite *AiUsecaseTestSuite) TestGenerateTextWithTags_Success() {
	ctx := context.Background()
	tags := []domain.Tag{"technology", "AI"}
	expectedPrompt := "Generate a blog post using the following tags: technology, AI. The blog should be engaging, informative, and well-structured."
	expectedContent := "Generated blog content"

	suite.mockAiRepo.On("GenerateText", mock.Anything, expectedPrompt).Return(expectedContent, nil)

	content, err := suite.aiUsecase.GenerateTextWithTags(ctx, tags)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedContent, content)
	suite.mockAiRepo.AssertExpectations(suite.T())
}

func (suite *AiUsecaseTestSuite) TestGenerateTextWithTags_Error() {
	ctx := context.Background()
	tags := []domain.Tag{"technology", "AI"}
	expectedPrompt := "Generate a blog post using the following tags: technology, AI. The blog should be engaging, informative, and well-structured."
	expectedError := errors.New("generation failed")

	suite.mockAiRepo.On("GenerateText", mock.Anything, expectedPrompt).Return("", expectedError)

	content, err := suite.aiUsecase.GenerateTextWithTags(ctx, tags)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "", content)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockAiRepo.AssertExpectations(suite.T())
}

func (suite *AiUsecaseTestSuite) TestGenerateTextWithPrompt_Success() {
	ctx := context.Background()
	prompt := "Generate a blog post about AI"
	expectedContent := "Generated blog content about AI"

	suite.mockAiRepo.On("GenerateText", mock.Anything, prompt).Return(expectedContent, nil)

	content, err := suite.aiUsecase.GenerateTextWithPrompt(ctx, prompt)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedContent, content)
	suite.mockAiRepo.AssertExpectations(suite.T())
}

func (suite *AiUsecaseTestSuite) TestGenerateTextWithPrompt_Error() {
	ctx := context.Background()
	prompt := "Generate a blog post about AI"
	expectedError := errors.New("generation failed")

	suite.mockAiRepo.On("GenerateText", mock.Anything, prompt).Return("", expectedError)

	content, err := suite.aiUsecase.GenerateTextWithPrompt(ctx, prompt)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "", content)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockAiRepo.AssertExpectations(suite.T())
}

func (suite *AiUsecaseTestSuite) TestGenerateSuggestions_Success() {
	ctx := context.Background()
	textContent := "This is a blog post about AI"
	expectedPrompt := "Provide suggestions to improve the following blog content: This is a blog post about AI. The suggestions should be constructive, specific, and aimed at enhancing the quality and readability of the blog."
	expectedSuggestions := "Here are some suggestions to improve the blog post..."

	suite.mockAiRepo.On("GenerateSuggestions", mock.Anything, expectedPrompt).Return(expectedSuggestions, nil)

	suggestions, err := suite.aiUsecase.GenerateSuggestions(ctx, textContent)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedSuggestions, suggestions)
	suite.mockAiRepo.AssertExpectations(suite.T())
}

func (suite *AiUsecaseTestSuite) TestGenerateSuggestions_Error() {
	ctx := context.Background()
	textContent := "This is a blog post about AI"
	expectedPrompt := "Provide suggestions to improve the following blog content: This is a blog post about AI. The suggestions should be constructive, specific, and aimed at enhancing the quality and readability of the blog."
	expectedError := errors.New("suggestion generation failed")

	suite.mockAiRepo.On("GenerateSuggestions", mock.Anything, expectedPrompt).Return("", expectedError)

	suggestions, err := suite.aiUsecase.GenerateSuggestions(ctx, textContent)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "", suggestions)
	assert.Equal(suite.T(), expectedError, err)
	suite.mockAiRepo.AssertExpectations(suite.T())
}

func TestAiUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AiUsecaseTestSuite))
}
