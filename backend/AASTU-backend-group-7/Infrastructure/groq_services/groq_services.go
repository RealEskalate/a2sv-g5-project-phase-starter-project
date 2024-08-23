package groqservice

import (
	"blogapp/Config"
	"blogapp/Dtos"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ChatCompletionInterface defines the interface for chat completion
type ChatCompletionInterface interface {
	GetChatCompletion(prompt string) (string, error)
	GetChatCompletionByTags(tags []string) (string, error)
	GetChatCompletionEnhancements(prompt string) (string, error)
	GenerateBlog(postDTO Dtos.PostDTO) (string, error)
}

const GROQAPIURL = "https://api.groq.com/openai/v1/chat/completions" // Placeholder URL

// Structs for request and response
type ChatCompletionRequest struct {
	Messages []ChatMessage `json:"messages"`
	Model    string        `json:"model"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message ChatMessage `json:"message"`
}

// GroqAI service struct
var APIKey = "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"

type GroqAI struct {
	APIKey string
	test   string
}

func NewGroqAI(apiKey string) *GroqAI {
	Config.Envinit()
	return &GroqAI{APIKey: apiKey,
		test: "test"}
}

// GetChatCompletion method to call the GROQ API
func (g *GroqAI) GetChatCompletion(prompt string) (string, error) {
	requestBody := ChatCompletionRequest{
		Messages: []ChatMessage{{Role: "user", Content: prompt}},
		Model:    "mixtral-8x7b-32768",
	}

	// 	"model": "mixtral-8x7b-32768",

	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", GROQAPIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	fmt.Println("apikeys match:", g.APIKey == APIKey, g.APIKey, APIKey, g.test, Config.GROQ_API_KEY)
	req.Header.Set("Authorization", "Bearer "+g.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("failed to get a valid response")
	}

	var response ChatCompletionResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}

	return (response.Choices[0].Message.Content), nil

}

func (g *GroqAI) GetChatCompletionByTags(tags []string) (string, error) {
	prompt := "Generate a blog based on tags: " + strings.Join(tags, ", ")
	return g.GetChatCompletion(prompt)
}

func (g *GroqAI) GetChatCompletionEnhancements(prompt string) (string, error) {
	prompt = "Enhance this blog and provide numbered tips: " + prompt
	return g.GetChatCompletion(prompt)
}

func (g *GroqAI) GenerateBlog(postDTO Dtos.PostDTO) (string, error) {
	setDefaultValues(&postDTO)
	fmt.Println(postDTO)
	prompt := "Generate a blog with the following details:\n" +
		"Title: " + postDTO.Title + "\n" +
		"Categories: " + strings.Join(postDTO.Categories, ", ") + "\n" +
		"Keywords: " + strings.Join(postDTO.Keywords, ", ") + "\n" +
		"Tone: " + postDTO.Tone + "\n" +
		"Format: " + postDTO.Format + "\n" +
		"ParagraphLimit: " + strconv.Itoa(postDTO.ParagraphLimit) + "\n" +
		"Max Word Limit: " + strconv.Itoa(postDTO.MaxWordLimit) + "\n" +
		"Additional Context: " + postDTO.AdditionalContext

	return g.GetChatCompletion(prompt)
}

var setDefaultValues = func(postDTO *Dtos.PostDTO) {
	if postDTO.Title == "" {
		postDTO.Title = "Default Title"
	}
	if postDTO.Categories == nil {
		postDTO.Categories = []string{"Default Category"}
	}
	if postDTO.Keywords == nil {
		postDTO.Keywords = []string{"Default Keyword"}
	}
	if postDTO.Tone == "" {
		postDTO.Tone = "Formal"
	}
	if postDTO.Format == "" {
		postDTO.Format = "Article"
	}
	if postDTO.ParagraphLimit == 0 {
		postDTO.ParagraphLimit = 5
	}
	if postDTO.MaxWordLimit == 0 {
		postDTO.MaxWordLimit = 500
	}
}
