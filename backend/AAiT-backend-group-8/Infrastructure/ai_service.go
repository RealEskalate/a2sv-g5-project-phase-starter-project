package Infrastructure

import (
	"context"
	"fmt"

	interfaces "AAiT-backend-group-8/Interfaces"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GenAIService struct {
	client *genai.Client
}

func NewGenAIService(apiKey string) (interfaces.IAiService, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &GenAIService{client: client}, nil
}

func (s *GenAIService) GenerateContent(ctx context.Context, userInput string) (string, error) {
	model := s.client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("Generate a blog post based on the title: %s. Output should follow this structure:1. First line is the title.2. Second line is the body text, with a maximum of 1000 words.3. After the body, write Tags: followed by a comma-separated list of tags.Ensure the body and tags are clearly separated.", userInput)))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}

func (s *GenAIService) SuggestImprovements(ctx context.Context, title, body, tags string) (string, error) {
	model := s.client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf(
		"Here is a blog with title: '%s', body: '%s', and tags: '%s'.\n\n"+
			"Provide a comment not more than 25 word  to improve it and a revised version with body not more than 1000 word.",
		title, body, tags)))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}
