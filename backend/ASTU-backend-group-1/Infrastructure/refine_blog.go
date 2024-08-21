package infrastructure

import (
	"astu-backend-g1/config"
	"context"
	"fmt"
)

func Refine(content string) (string, error) {
	prompt := fmt.Sprintf(`Please refine the following content to make it more engaging, clear, and concise. Focus on improving the flow, enhancing readability, and ensuring that the main points are emphasized effectively. Feel free to rephrase sentences, restructure paragraphs, and add any necessary transitions. The tone should remain professional yet approachable. And make sure that you don add any title please and no comments.: %v`, content)
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
