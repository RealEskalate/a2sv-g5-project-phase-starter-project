package infrastructure

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)


type LlmClientInterface interface {
	GenerateText(prompt string) (*genai.GenerateContentResponse, error)
}

type LlmClient struct {
	LLMModel *genai.GenerativeModel
}

func  NewLlmClient() *LlmClient {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	
	model := client.GenerativeModel("gemini-1.5-flash")

	return &LlmClient{
		LLMModel: model,
	}
}

func (llmClient *LlmClient) GenerateText(prompt string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	response, err := llmClient.LLMModel.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}
	return response, nil
}

