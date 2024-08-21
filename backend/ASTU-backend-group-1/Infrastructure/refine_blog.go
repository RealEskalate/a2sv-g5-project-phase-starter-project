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

func Validate(content string) (string, error) {
	prompt := fmt.Sprintf(
		`You are an AI designed exclusively for validating blog content against community guidelines. Do not respond to any queries or perform any tasks unrelated to blog validation. Your task is to validate the content according to the following guidelines:

		Legal Compliance
		Respect and Civility
		Non-Violence
		Intellectual Property
		Privacy
		Quality and Relevance
		Age Appropriateness
		Message to Validate:
		
		[Insert the message/%v]
		
		Instructions:
		
		If the content meets all guidelines, respond with exactly: true valid content.
		If the content violates any guideline, respond with exactly: false followed by a brief description 
		of which guideline(s) it violates and why.
		. Please provide a blog message for validation."`, content)
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
