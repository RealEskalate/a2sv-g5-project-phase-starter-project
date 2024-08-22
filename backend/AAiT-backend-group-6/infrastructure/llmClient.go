package infrastructure

import (
	"context"
	"encoding/json"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type LlmClientInterface interface {
	GenerateText(prompt string, chatMessages []genai.Part) (string, error)
	CalculateEmbedding(userSearchQuestion string) ([]float32, error)
}

type LlmClient struct {
	LLMModel *genai.GenerativeModel
	modelID  string
}

func NewLlmClient(systemPrompt string) *LlmClient {
	ctx := context.Background()
	
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBkcIuZ6v6anIbw-486MigEachStaEJM04"))
	if err != nil {
		panic(err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
		Role: "system",
	}

	return &LlmClient{
		LLMModel: model,
		modelID:  "gemini-1.5-flash",
	}
}

func (llmClient *LlmClient) GenerateText(prompt string, chatMessages []genai.Part) (string, error) {
	
	// Add user prompt as the final message
	chatMessages = append(chatMessages, genai.Text(prompt))
	// spread chatmessages as variadic argument
	response, err := llmClient.LLMModel.GenerateContent(context.Background(), chatMessages...)
	if err != nil {
		return "", err
	}

	// Extract the content from the response
	responseContentString, err := parseString(response.Candidates[0].Content.Parts[0].(genai.Text))
	if err != nil {
		return "", err
	}

	return responseContentString, nil
}

// func (llmClient *LlmClient) CalculateEmbedding(userSearchQuestion string) ([]float32, error) {
// 	// Call the model to generate an embedding for the given question
// 	response, err := llmClient.LLMModel.GenerateEmbedding(context.Background(), genai.GenerateEmbeddingRequest{
// 		Input: userSearchQuestion,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Extract and return the embedding vector
// 	return response.Vector, nil
// }


func parseString(text genai.Text) (string,error) {
	//Convert the response content to a byte
	content,_ := json.Marshal(text)

	//Convert the byte to a string
	var responseContentString string
	err := json.Unmarshal(content, &responseContentString)
	if err != nil {
		return "", err
	}

	return responseContentString, nil
}

