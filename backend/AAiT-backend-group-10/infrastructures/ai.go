package infrastructures

import (
	"context"
	"log"
	"reflect"
	"strings"

	"aait.backend.g10/domain"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AiService struct {
    model *genai.GenerativeModel
    ctx   context.Context
}

func NewAIService(apiKey string) *AiService {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        log.Fatal(err)
    }
    model := client.GenerativeModel("gemini-pro")
    return &AiService{
        model: model,
        ctx:   ctx,
    }
}

func (a *AiService) GenerateContent(topic string, keywords []string) (*domain.BlogContentResponse, error) {
    prompt := generatePrompt(topic, keywords)

    resp, err := a.model.GenerateContent(a.ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error generating content: %v", err)
		return nil, err
	}

    if len(resp.Candidates) == 0 {
		return &domain.BlogContentResponse{
            SuggestedContent: "No candidates found",
        }, nil
	}
	if len(resp.Candidates[0].Content.Parts) == 0 {
		return &domain.BlogContentResponse{
            SuggestedContent: "No content parts found",
        }, nil
	}

    generatedContent := a.CleanText(resp.Candidates[0].Content.Parts[0])
	if generatedContent == "" {
		return &domain.BlogContentResponse{
            SuggestedContent: "Content extraction failed",
        }, nil
	}

    return &domain.BlogContentResponse{
        SuggestedContent: generatedContent,
    }, nil
}

func (a *AiService) SuggestImprovements(content string) (*domain.SuggestionResponse, error) {
    prompt := "Review the following blog content and provide suggestions or enhancements:\n\n" + content +
		"\n\nProvide constructive feedback, highlight improvements, and suggest any enhancements."

    resp, err := a.model.GenerateContent(a.ctx, genai.Text(prompt))
    if err != nil {
        log.Printf("Error generating suggestions: %v", err)
        return nil, err
    }

    if len(resp.Candidates) == 0 {
		return &domain.SuggestionResponse{
            Suggestions: "No suggestions found",
        }, nil
	}
	if len(resp.Candidates[0].Content.Parts) == 0 {
		return &domain.SuggestionResponse{
            Suggestions: "No content parts found",
        }, nil
	}

	suggestions := a.CleanText(resp.Candidates[0].Content.Parts[0])
	if suggestions == "" {
		return &domain.SuggestionResponse{
            Suggestions: "Suggestion extraction failed",
        }, nil
	}

	return &domain.SuggestionResponse{
        Suggestions: suggestions,
    }, nil
    
}

func generatePrompt(topic string, keywords []string) string {
    return "Generate a blog post on the topic '" + topic + "' with the following keywords: " + strings.Join(keywords, ", ") +
		". The content should be engaging, include relevant subheadings, and provide useful insights. " +
		"Return the content in a well-structured format."
}

func (a *AiService) CleanText(value interface{}) string {
	text := a.ExtractText(value)

	cleanedText := strings.ReplaceAll(strings.ReplaceAll(text, "**", ""), "*", "")
	cleanedText = strings.ReplaceAll(cleanedText, "\n\n", "\n")

	return cleanedText
}

func (a *AiService) ExtractText(value interface{}) string {
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
