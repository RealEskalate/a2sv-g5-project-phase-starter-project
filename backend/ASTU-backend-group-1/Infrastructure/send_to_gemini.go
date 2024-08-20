package infrastructure

import (
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func connectToGemini(apiKey, modelName string, ctx context.Context) (*genai.GenerativeModel, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return &genai.GenerativeModel{}, err
	}
	return client.GenerativeModel(modelName), nil
}

func SendPrompt(prompt, apiKey, modelName string, ctx context.Context) (string, error) {
	model, err := connectToGemini(apiKey, modelName, ctx)
	if err != nil {
		return "", fmt.Errorf("connecting to gemini: %v", err)
	}
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("no response from the model")
	}
	candidate := resp.Candidates[0]
	responseText := fmt.Sprint(candidate.Content.Parts[0])
	return responseText, nil
}
