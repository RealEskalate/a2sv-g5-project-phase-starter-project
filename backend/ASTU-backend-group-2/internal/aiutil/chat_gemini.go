package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIUtil interface {
	GenerateContentFromGemini(title string, description string, env bootstrap.Env) (string, error)
}

type AI struct {
	model *genai.GenerativeModel
}

func NewAIUtil(env *bootstrap.Env) AIUtil {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(env.GeminiAPIKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	return &AI{
		model: model,
	}
}

type Content struct {
	Parts []string `json:"parts"`
	Role  string   `json:"role"`
}

type Candidates struct {
	Content *Content `json:"content"`
}

type ContentResponse struct {
	Candidates *[]Candidates `json:"candidates"`
}

func (ai *AI) GenerateContentFromGemini(title string, description string, env bootstrap.Env) (string, error) {

	words, err := strconv.Atoi(env.GeminiWordCount)

	if err != nil {
		return "", fmt.Errorf("error converting word count to integer: %v", err)
	}

	prompt := []genai.Part{
		genai.Text(fmt.Sprintf("Create a blog post titled '%s' with approximately %d words of content. The post should focus on the following key points: '%s'. Please ensure the content is informative, engaging, and well-structured.",
			title,
			words,
			description)),
	}

	resp, err := ai.model.GenerateContent(context.TODO(), prompt...)
	if err != nil {
		return "", fmt.Errorf("error generating content: %v", err)
	}

	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")

	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	var blogContent string
	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				blogContent += part
			}
		}
	}

	return blogContent, nil
}

func cleanText(value interface{}) string {
	text := extractText(value)

	cleanedText := strings.ReplaceAll(strings.ReplaceAll(text, "**", ""), "*", "")
	cleanedText = strings.ReplaceAll(cleanedText, "\n\n", "\n")

	return cleanedText
}

func extractText(value interface{}) string {
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
		// Directly return the string if it's a string type
		return v.String()

	default:
		log.Printf("Unsupported type %T for field extraction", value)
		return ""
	}
}

func (ai *AI) ReviewContent(blogContent string) (string, error) {
	prompt := fmt.Sprintf("Please review the blog content below and offer constructive feedback, highlighting any areas for improvement or suggesting enhancements. Your feedback should be clear, concise, and actionable. \n \n %v", blogContent)

	resp, err := ai.model.GenerateContent(context.TODO(), genai.Text(prompt))
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

	suggestions := cleanText(resp.Candidates[0].Content.Parts[0])
	if suggestions == "" {
		return "Suggestions extraction failed", nil
	}

	return suggestions, nil
}

func (aiService *AI) SendMessage(ctx context.Context, history []entities.Message, message entities.Message) (entities.Message, error) {
	var chatSessionHistory []*genai.Content
	for _, message := range history {
		content := genai.Content{
			Parts: []genai.Part{
				genai.Text(message.Text),
			},

			Role: message.Role,
		}

		chatSessionHistory = append(chatSessionHistory, &content)
	}

	chatSession := aiService.model.StartChat()
	chatSession.History = chatSessionHistory

	response, err := chatSession.SendMessage(ctx, genai.Text(message.Text))
	if err != nil {
		return entities.Message{}, err
	}

	candidate := response.Candidates[0]
	return entities.Message{
		Text:      fmt.Sprintf("%s", candidate.Content.Parts[0]),
		Role:      candidate.Content.Role,
		CreatedAt: time.Now(),
	}, nil
}
