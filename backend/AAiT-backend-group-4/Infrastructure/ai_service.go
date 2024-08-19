package infrastructure

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type aiService struct {
	Env    *bootstrap.Env
	Client *genai.Client
	Model  *genai.GenerativeModel
}

func NewAiService(env *bootstrap.Env) domain.AiRepository {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(env.GeminiApiKey))
	if err != nil {
		log.Fatalf("Failed to create GenAI client: %v", err)
	}
	return &aiService{
		Env:    env,
		Client: client,
		Model:  client.GenerativeModel("gemini-1.5-flash"),
	}
}

func (as *aiService) GenerateText(c context.Context, prompt string) (string, error) {
	resp, err := as.Model.GenerateContent(c, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var text string

	if resp != nil {
		candindates := resp.Candidates
		for _, candidate := range candindates {
			content := candidate.Content
			if content != nil && len(content.Parts) > 0 {
				text = fmt.Sprintf(text, content.Parts[0])
			}
		}
	}

	if text == "" {
		return "", fmt.Errorf("no content generated")
	}

	return text, nil

}

func (as *aiService) GenerateSuggestions(c context.Context, textContent string) (string, error) {
	resp, err := as.Model.GenerateContent(c, genai.Text("Can you give me a suggestion,places where i can fix and add more details for the following blog: "+textContent))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var text string

	if resp != nil {
		candindates := resp.Candidates
		for _, candidate := range candindates {
			content := candidate.Content
			if content != nil && len(content.Parts) > 0 {
				text = fmt.Sprintf(text, content.Parts[0])
			}
		}
	}

	if text == "" {
		return "", fmt.Errorf("no suggestions")
	}

	return text, nil
}

func (as *aiService) Chat(c context.Context, textContent string) (string, error) {
	return "", nil
}
