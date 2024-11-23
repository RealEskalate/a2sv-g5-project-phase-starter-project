package infrastructure

import (
	"blog-api/domain"
	"context"
	"errors"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GenAIService struct{}

// NewGenAIService creates a new instance of the AI content generator.
func NewGenAIService() domain.AIContentGenerator {
	return &GenAIService{}
}

// GenerateContent calls the Gemini AI API to generate content based on the given prompt.
func (s *GenAIService) GenerateContent(prompt string, API_Key string) (genai.Part, error) {
	if prompt == "" {
		return nil, errors.New("prompt cannot be empty") // Ensuring prompt is not empty before making the API call
	}

	ctx := context.Background()
	apiKey := API_Key
	if apiKey == "" {
		return nil, errors.New("API key is not set in the environment variables") // Check if API key is available
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close() // Ensure client is closed after the operation

	model := client.GenerativeModel("gemini-1.5-flash")
	response, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	// Check the actual structure of the response and extract the text accordingly
	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		response := response.Candidates[0].Content.Parts[0]

		return response, nil
	}
	return nil, errors.New("no content generated")
}
