package infrastructure

import (
	"Blog_Starter/config"
	"Blog_Starter/utils"
	"context"
	"encoding/json"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// I need a utility to generate AI service so it accept a content and return ai genereted content
type AIService struct {
	// I need a utility to generate AI service so it accept a content and return ai genereted content
}

// NewAIService is a constructor to create AIService instance
func NewAIService() utils.AIService {
	// I need a utility to generate AI service so it accept a content and return ai genereted content
	return &AIService{}
}

// GenerateAIContent is a function to generate AI content
func (ai *AIService) GenerateAIContent(ctx context.Context, content string, preText string) (string, error) {
	// Access your API key as an environment variable (see "Set up your API key" above)

	// TODO: Add API key to .env file
	API_KEY := config.NewEnv().APIKEY
	// I need a utility to generate AI segirvice so it accept a content and return ai genereted content

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))

	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	// Construct the prompt with pre-text and user-provided content
	prompt := preText + " " + content

	// Generate text
	type Response struct {
		Candidates []struct {
			Content struct {
				Parts []string `json:"Parts"`
			} `json:"Content"`
		} `json:"Candidates"`
	}

	response, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	marshalResponse, _ := json.MarshalIndent(response, "", " ")

	var resp Response

	err = json.Unmarshal(marshalResponse, &resp)
	if err != nil {
		return "", err
	}

	text := resp.Candidates[0].Content.Parts[0]

	return text, nil
}
