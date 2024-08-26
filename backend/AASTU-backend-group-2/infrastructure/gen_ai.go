package infrastructure

import (
	"blog_g2/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiAIService struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

type ContentResponse struct {
	Candidates *[]struct {
		Content *struct {
			Parts []string `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// NewGeminiAIService creates a new instance of GeminiAIService
func NewGeminiAIService() (domain.AIService, *domain.AppError) {
	apiKey := DotEnvLoader("GEMINI_API_KEY")

	if apiKey == "" {
		return nil, domain.ErrGeminiAPIKeyMissing
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, domain.ErrGeminiClientCreation
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	if model == nil {
		return nil, domain.ErrGeminiModelNotAvailable
	}

	return &GeminiAIService{
		client: client,
		model:  model,
	}, nil
}

// GeneratePost generates a blog post using the Gemini AI
func (s *GeminiAIService) GeneratePost(title, description string) (*domain.PostResponse, *domain.AppError) {
	prompt := []genai.Part{
		genai.Text(fmt.Sprintf("Title: %s\n\nDescription: %s\n\nWrite a blog post based on the above title and description:", title, description)),
	}

	resp, err := s.model.GenerateContent(context.Background(), prompt...)
	if err != nil {
		return nil, domain.ErrGeminiContentGeneration
	}

	// Marshal and unmarshal response to process the content
	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		return nil, domain.ErrGeminiResponseParsing
	}

	var generatedContent string

	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				generatedContent += part
			}
		}
	}

	splitContent := strings.SplitN(generatedContent, "\n\n", 2)
	var generatedTitle, body string

	if len(splitContent) > 1 {
		generatedTitle = strings.TrimSpace(splitContent[0])
		body = strings.TrimSpace(splitContent[1])
	} else {
		body = strings.TrimSpace(splitContent[0])
	}

	return &domain.PostResponse{
		Title:   generatedTitle,
		Content: body,
	}, nil
}

// Validate_Comment checks if a comment contains offensive language
func (s *GeminiAIService) Validate_Comment(comment string) *domain.AppError {
	prompt := []genai.Part{
		genai.Text(fmt.Sprintf("Is the following comment offensive or contains offensive languages? Respond with Yes or No: \"%s\"", comment)),
	}

	resp, err := s.model.GenerateContent(context.Background(), prompt...)
	if err != nil {
		if strings.Contains(err.Error(), "FinishReasonSafety") {
			return domain.ErrOffensiveComment

		}
		return domain.ErrGeminiContentGeneration
	}

	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(resp)
	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		log.Fatal(err)
	}

	var generatedContent string

	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				generatedContent += part
			}
		}
	}

	fmt.Println(string(generatedContent))

	if strings.HasPrefix(generatedContent, "Yes") {
		return domain.ErrOffensiveComment
	}

	return nil
}

// Validate_Blog checks if a blog post contains offensive language
func (s *GeminiAIService) Validate_Blog(blog string) *domain.AppError {
	prompt := []genai.Part{
		genai.Text(fmt.Sprintf("Is the following blog post offensive or contains offensive languages? Respond with Yes or No: \"%s\"", blog)),
	}

	resp, err := s.model.GenerateContent(context.Background(), prompt...)
	if err != nil {
		if strings.Contains(err.Error(), "FinishReasonSafety") {
			return domain.ErrOffensiveBlogContent
		}
		return domain.ErrGeminiContentGeneration
	}

	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(resp)
	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		log.Fatal(err)
	}

	var generatedContent string

	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				generatedContent += part
			}
		}
	}

	fmt.Println(string(generatedContent))

	if strings.HasPrefix(generatedContent, "Yes") {
		return domain.ErrOffensiveBlogContent
	}

	return nil
}
