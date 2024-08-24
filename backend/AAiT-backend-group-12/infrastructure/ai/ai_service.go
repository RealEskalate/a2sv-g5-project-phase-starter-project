package ai_service

import (
	"blog_api/domain"
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)

// AIService provides an interface to interact with the AI service.
type AIService struct {
	Model domain.AIModelInterface
	Ctx   context.Context
}

// NewAIService creates a new AIService instance with the provided API key.
func NewAIService(apiKey string) *AIService {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-pro")

	return &AIService{
		Model: model,
		Ctx:   ctx,
	}
}

// CleanText removes unwanted characters and formatting from the provided text.
func (s *AIService) CleanText(value interface{}) string {
	text := s.ExtractText(value)

	cleanedText := strings.ReplaceAll(strings.ReplaceAll(text, "**", ""), "*", "")
	cleanedText = strings.ReplaceAll(cleanedText, "\n\n", "\n")

	return cleanedText
}

// extractText extracts the 'Text' field from a struct or returns the string value directly.
func (s *AIService) ExtractText(value interface{}) string {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Struct:
		field := v.FieldByName("Text")
		if !field.IsValid() {
			log.Printf("Field 'Text' not found in struct of type %T", value)
			return ""
		}
		return field.String()

	case reflect.String:
		return v.String()

	default:
		log.Printf("Unsupported type %T for field extraction", value)
		return ""
	}
}

// GenerateContent generates content based on topics provided.
func (s *AIService) GenerateContent(topics []string) (string, error) {
	prompt := "Generate a blog post about " + strings.Join(topics, ", ") + ". " +
		"The content should be engaging, include relevant subheadings, and provide useful insights. " +
		"Return the content in a well-structured format."

	resp, err := s.Model.GenerateContent(s.Ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error generating content: %v", err)
		return "", err
	}

	if len(resp.Candidates) == 0 {
		return "No candidates found", nil
	}

	if len(resp.Candidates[0].Content.Parts) == 0 {
		return "No content parts found", nil
	}

	generatedContent := s.CleanText(resp.Candidates[0].Content.Parts[0])
	if generatedContent == "" {
		return "Content extraction failed", nil
	}

	return generatedContent, nil
}

// ReviewContent analyzes the provided content and generates AI-based suggestions or enhancements.
func (s *AIService) ReviewContent(blogContent string) (string, error) {
	prompt := "Review the following blog content and provide suggestions or enhancements:\n\n" + blogContent +
		"\n\nProvide constructive feedback, highlight improvements, and suggest any enhancements."

	resp, err := s.Model.GenerateContent(s.Ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error generating review content: %v", err)
		return "", err
	}

	if len(resp.Candidates) == 0 {
		return "No suggestions found", nil
	}
	if len(resp.Candidates[0].Content.Parts) == 0 {
		return "No content parts found", nil
	}

	suggestions := s.CleanText(resp.Candidates[0].Content.Parts[0])
	if suggestions == "" {
		return "Suggestions extraction failed", nil
	}

	return suggestions, nil
}

// GenerateTrendingTopics generates a list of trending blog topics based on the provided keywords.
func (s *AIService) GenerateTrendingTopics(keywords []string) ([]string, error) {
	prompt := "Based on the following keywords: " + strings.Join(keywords, ", ") +
		", generate a list of trending blog topics that are currently popular."

	resp, err := s.Model.GenerateContent(s.Ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error generating trending topics: %v", err)
		return nil, err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no topics generated")
	}

	generatedText := s.CleanText(resp.Candidates[0].Content.Parts[0])
	topics := strings.Split(generatedText, "\n") // Assuming the AI returns a list separated by newlines

	return topics, nil
}
