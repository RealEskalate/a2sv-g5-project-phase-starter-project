package infrastructure

import (
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
func (ai *AIService) GenerateAIContent(ctx context.Context, content string) (string, error) {
	API_KEY := "AIzaSyBfy1w8ZR4C6xOq43duYKz4RU1wbmsKl18"
	// I need a utility to generate AI service so it accept a content and return ai genereted content

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	// Generate text

	type Response struct {
		Candidates []struct {
			Content struct {
				Parts []string `json:"Parts"`
			} `json:"Content"`
		} `json:"Candidates"`
	}

	response, err := model.GenerateContent(ctx, genai.Text(content))
	if err != nil {
		return "", err
	}
	marshalResponse, _ := json.MarshalIndent(response, "", "  ")

	var resp Response

	err = json.Unmarshal(marshalResponse, &resp)

	if err != nil {
		return "", err
	}

	text := resp.Candidates[0].Content.Parts[0]

	return text, nil

}
