package infrastructure

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type aiService struct {
	Env    *bootstrap.Env
	Client *genai.Client
	Model  *genai.GenerativeModel
}

// NewAiService creates a new instance of AiService that implements the AiRepository interface.
// It takes an environment configuration as input and returns an AiRepository.
// The AiService uses the GenAI client to interact with the AI service.
// It initializes the client with the provided Gemini API key from the environment configuration.
// If the client initialization fails, it logs an error and terminates the program.
// The AiService also sets the model to "gemini-1.5-flash" for generating AI models.
func NewAiService(env *bootstrap.Env) domain.AiRepository {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(env.GeminiApiKey))
	if err != nil {
		log.Fatalf("Failed to create GenAI client: %v", err)
	}
	return &aiService{
		Env:    env,
		Client: client,
		Model:  client.GenerativeModel("gemini-1.5-flash"),
	}
}

// GenerateText generates text based on the given prompt using the AI model.
// It takes a context and a prompt string as input and returns the generated text and an error.
// The generated text is obtained by calling the GenerateContent method of the AI model with the provided prompt.
// If an error occurs during the generation process, it will be returned.
// The function checks if the response is not nil and iterates over the candidates to find the content.
// If the content is not nil and has at least one part, it appends the first part to the generated text.
// If no content is generated, an error is returned with the message "no content generated".
func (as *aiService) GenerateText(c context.Context, prompt string) (string, error) {
	resp, err := as.Model.GenerateContent(c, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var text string

	if resp != nil {
		candindates := resp.Candidates
		for _, candidate := range candindates {
			content := candidate.Content
			if content != nil && len(content.Parts) > 0 {
				text = fmt.Sprintf(text, content.Parts[0])
			}
		}
	}

	if text == "" {
		return "", fmt.Errorf("no content generated")
	}

	return text, nil

}

// GenerateSuggestions generates suggestions for improving a blog by providing places where more details can be added or fixed.
// It takes a context and the text content of the blog as input and returns the generated suggestions as a string.
// If there are no suggestions, it returns an error with the message "no suggestions".
func (as *aiService) GenerateSuggestions(c context.Context, textContent string) (string, error) {
	resp, err := as.Model.GenerateContent(c, genai.Text("Can you give me a suggestion,places where i can fix and add more details for the following blog: "+textContent))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var text string

	if resp != nil {
		candindates := resp.Candidates
		for _, candidate := range candindates {
			content := candidate.Content
			if content != nil && len(content.Parts) > 0 {
				text = fmt.Sprintf(text, content.Parts[0])
			}
		}
	}

	if text == "" {
		return "", fmt.Errorf("no suggestions")
	}

	return text, nil
}

func (as *aiService) Chat(c context.Context, textContent string) (string, error) {
	return "", nil
}
