package infrastructure

import (
	"context"
	"encoding/json"
	"meleket/domain"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIService interface {
	GenerateAIContent(prompt string) (string, error)
}

type aiService struct {
	client *genai.Client
}

func NewAIService() AIService {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		panic(err)
	}

	return &aiService{
		client: client,
	}
}

func (s *aiService) GenerateAIContent(prompt string) (string, error) {
	ctx := context.Background()
	model := s.client.GenerativeModel("gemini-1.5-flash")
	model.ResponseMIMEType = "text/plain"

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}


	respBytes, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}
	
	var response domain.Response
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return "", err
	}

	// Extract the generated content
	if len(response.Candidates) > 0 && response.Candidates[0].Content != nil {
		return response.Candidates[0].Content.Parts[0], nil
	}

	return "", nil
}
