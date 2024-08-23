package infrastructure

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiModel struct {
	model *genai.GenerativeModel
	topic string
}

func connectToGemini(apiKey, modelName string, ctx context.Context) (*genai.GenerativeModel, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return &genai.GenerativeModel{}, err
	}
	return client.GenerativeModel(modelName), nil
}

func NewGeminiModel(apiKey, modelName, topic string) (*GeminiModel, error) {
	model, err := connectToGemini(apiKey, modelName, context.Background())
	if err != nil {
		return &GeminiModel{}, err
	}
	return &GeminiModel{model: model, topic: topic}, nil
}

func (g *GeminiModel) SendPrompt(prompt string) (string, error) {
	resp, err := g.model.GenerateContent(context.Background(), genai.Text(prompt))
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

func (g *GeminiModel) CheckPromptContent(prompt string) error {
	wrapperPrompt := `"Please provide a specific answer to the following question: [Is this prompt about blogging and nothing else: "%v"]. Please use only yes or no in your response."`
	finalPrompt := fmt.Sprintf(wrapperPrompt, prompt)
	resp, err := g.SendPrompt(finalPrompt)
	if err != nil {
		return fmt.Errorf("getting response error: %v", err.Error())
	}
	if strings.ToLower(resp) == "no" {
		return fmt.Errorf("prompt not about blogging")
	}
	return nil
}

func (g *GeminiModel) Refine(content string) (string, error) {
	prompt := fmt.Sprintf(`Please refine the following content to make it more engaging, clear, and concise. Focus on improving the flow, enhancing readability, and ensuring that the main points are emphasized effectively. Feel free to rephrase sentences, restructure paragraphs, and add any necessary transitions. The tone should remain professional yet approachable. And make sure that you don add any title please and no comments.: %v`, content)
	refinedContent, err := g.SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return refinedContent, nil
}

func (g *GeminiModel) Validate(content string) (string, error) {
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
	refinedContent, err := g.SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return refinedContent, nil
}

func (g *GeminiModel) RecommendTitle(content string) (string, error) {
	prompt := fmt.Sprintf(`Please provide a single, concise, and engaging title for the following blog content in plain text, without using any markdown symbols, asterisks, bullet points, or other formatting elements. The response should consist only of the title. And include multiple titles for choosing don't include any of the mentioned symbols and I want maximum of five: %v`, content)
	if err := g.CheckPromptContent(prompt); err != nil {
		return "", err
	}
	recommendedTitles, err := g.SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return recommendedTitles, nil
}

func (g *GeminiModel) Chat(content string) (string, error) {
	if err := g.CheckPromptContent(content); err != nil {
		return "", err
	}
	resp, err := g.SendPrompt(content)
	if err != nil {
		return "", err
	}
	return resp, nil
}

/* func (g *GeminiModel) RecommendContent(title string, tags []string) (string, error) {
	wrapperPrompt := `Task: Recommend blog posts based on title and tags.
Input:
Title: [%v]
Tags: [%v]
Output:
List of recommended blog posts (include title, URL, brief description, author, publication date) without any asterisks or markdown formatting in plain text format`
	prompt := fmt.Sprintf(wrapperPrompt, title, tags)
} */
