package infrastructure

import (
	"astu-backend-g1/config"
	"context"
	"fmt"
)

func Refine(content string) (string, error) {
	prompt := fmt.Sprintf(`Refine this blog: %v`, content)
	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	refinedContent, err := SendPrompt(prompt, config.Gemini.ApiKey, config.Gemini.Model, context.Background())
	if err != nil {
		return "", err
	}
	return refinedContent, nil
}
