package services

import (
	"context"
	"encoding/json"

	"github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)


var (
	systemPrompt = "`Imagine that you are a bot that assists users in creating content for their blogs. Users rely on you to generate text content or improve their existing content. After asking the necessary questions, you will provide them with the correct content. Users cannot ask questions unrelated to blog content creation.`"
)

type AIService interface {
	GenerateText(propmt string, chatMesages []genai.Part) (string, error)
}

type aiService struct {
	model *genai.GenerativeModel
	ApiKey string
}

func NewAIService(apiKey string) AIService {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))

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

	return &aiService{
		model: model,
		ApiKey: apiKey,
	}

}

func (client *aiService) GenerateText(propmt string, chatMesages []genai.Part) (string, error) {
	chatMesages = append(chatMesages, genai.Text(propmt))

	response, err := client.model.GenerateContent(context.Background(), chatMesages...)
	if err != nil {
		return "", err
	}

	responseContentString, err := parseString(response.Candidates[0].Content.Parts[0].(genai.Text))

	if err != nil {	
		return "", err
	}

	return responseContentString, nil
}

func parseString(text genai.Text) (string, error) {
	content, err := json.Marshal(text)
	if err != nil {
		return "", err
	}

	var responseContentString string
	err = json.Unmarshal(content, &responseContentString)
	if err != nil {
		return "", err
	}

	return responseContentString, nil
}


