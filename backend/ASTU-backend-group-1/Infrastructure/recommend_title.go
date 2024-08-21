package infrastructure

import (
	"astu-backend-g1/config"
	"context"
	"fmt"
)

func RecommendTitle(content string) (string, error) {
	prompt := fmt.Sprintf(`lease provide a single, concise, and engaging title for the following blog content in plain text, without using any markdown symbols, asterisks, bullet points, or other formatting elements. The response should consist only of the title. And include multiple titles for choosing don't include any of the mentioned symbols and I want maximum of five: %v`, content)
	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	recommendedTitles, err := SendPrompt(prompt, config.Gemini.ApiKey, config.Gemini.Model, context.Background())
	if err != nil {
		return "", err
	}
	return recommendedTitles, nil
}
