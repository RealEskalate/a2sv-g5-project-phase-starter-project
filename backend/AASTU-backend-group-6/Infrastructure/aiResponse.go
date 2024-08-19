package infrastructure

import (
	domain "blogs/Domain"
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)


type AIService struct {
	client *genai.Client
	model  *genai.GenerativeModel
}
func NewAIService(config *Config) domain.AIConfig {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GeminiAPIKey))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")

	return &AIService{
		client: client,
		model:  model,
	}
}


func (s *AIService) Ask(ctx context.Context , formattedPrompt string) (*genai.GenerateContentResponse , error) {	
	
	result, err := s.model.GenerateContent(ctx, genai.Text(formattedPrompt))
	if err != nil {
		return nil , err 
	}
	
	return result , nil
}