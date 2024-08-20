package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

)

type Content struct {
	Parts []string `json:"Parts"`
	Role  string   `json:"Role"`
}

type Candidates struct {
	Content *Content `json:"Content"`
}

type ContentResponse struct {
	Candidates *[]Candidates `json:"Candidates"`
}

type AIUseCase struct {
	geminiApiKey string
}

func NewAIUseCase(geminiApiKey string) *AIUseCase {
	return &AIUseCase{geminiApiKey: geminiApiKey}
}

func (u *AIUseCase) GenerateContent(prompt string) ([]string, error) {
	// Initializing the AI client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(u.geminiApiKey))
	if err != nil {
		return nil, fmt.Errorf("error initializing AI client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	promptData := []genai.Part{
		genai.Text(prompt),
	}

	resp, err := model.GenerateContent(ctx, promptData...)
	if err != nil {
		return nil, fmt.Errorf("error generating content: %w", err)
	}

	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")

	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	var parts []string
	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			parts = append(parts, cad.Content.Parts...)
		}
	}

	return parts, nil
}
