package tests

import (
	"log"
	"os"
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/usecases"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BlogAssistantUseCaseTestSuite struct {
	suite.Suite
	usecase domain.BlogAssistantUseCase
}

func (suite *BlogAssistantUseCaseTestSuite) SetupTest() {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	suite.usecase = usecases.NewBlogAssistantUsecase(os.Getenv("GEMINI_API_KEY"))
}

func (suite *BlogAssistantUseCaseTestSuite) TestGenerateBlog() {
	keywords := []string{"AI", "Machine Learning", "Data Science"}
	tone := "professional"
	audience := "data scientists"

	blog, err := suite.usecase.GenerateBlog(keywords, tone, audience)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), blog)
}

func (suite *BlogAssistantUseCaseTestSuite) TestEnhanceBlog() {
	content := "AI is the future of technology."
	command := "Make it more engaging and informative."

	blog, err := suite.usecase.EnhanceBlog(content, command)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), blog)
}

func (suite *BlogAssistantUseCaseTestSuite) TestSuggestBlog() {
	industry := "AI"
	blog, err := suite.usecase.SuggestBlog(industry)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), blog)
}

func TestBlogAssistantUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BlogAssistantUseCaseTestSuite))
}

