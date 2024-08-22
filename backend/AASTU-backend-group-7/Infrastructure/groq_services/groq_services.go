// package groqservice

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// )

// const GROQAPIURL = "https://api.groq.com/openai/v1/chat/completions" // Placeholder URL

// // Structs for request and response
// type ChatCompletionRequest struct {
// 	Messages []ChatMessage `json:"messages"`
// 	Model    string        `json:"model"`
// }

// type ChatMessage struct {
// 	Role    string `json:"role"`
// 	Content string `json:"content"`
// }

// type ChatCompletionResponse struct {
// 	Choices []Choice `json:"choices"`
// }

// type Choice struct {
// 	Message ChatMessage `json:"message"`
// }

// // GroqAI service struct
// type GroqAI struct {
// 	APIKey string
// }

// func NewGroqAI(apiKey string) *GroqAI {
// 	return &GroqAI{APIKey: apiKey}
// }

// // GetChatCompletion method to call the GROQ API
// func (g *GroqAI) GetChatCompletion(prompt string) (string, error) {
// 	requestBody := ChatCompletionRequest{
// 		Messages: []ChatMessage{{Role: "user", Content: prompt}},
// 		Model:    "mixtral-8x7b-32768",
// 	}

// 	// 	"model": "mixtral-8x7b-32768",

// 	reqBody, err := json.Marshal(requestBody)
// 	if err != nil {
// 		return "", err
// 	}

// 	req, err := http.NewRequest("POST", GROQAPIURL, bytes.NewBuffer(reqBody))
// 	if err != nil {
// 		return "", err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+g.APIKey)
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		return "", errors.New("failed to get a valid response")
// 	}

// 	var response ChatCompletionResponse
// 	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
// 		return "", err
// 	}

// 	return response.Choices[0].Message.Content, nil
// }

package groqservice

import (
	"blogapp/Config"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
type GroqAI struct {
	APIKey string
}

var APIKey = Config.GROQ_API_KEY

func NewGroqAI(apiKey string) *GroqAI {
	fmt.Println("NewGroqAI", apiKey)
	return &GroqAI{APIKey: apiKey}
}

// GetChatCompletion method to call the GROQ API
func (g *GroqAI) GetChatCompletion(prompt string) (string, error) {
	requestBody := ChatCompletionRequest{
		Messages: []ChatMessage{{Role: "user", Content: prompt}},
		Model:    "mixtral-8x7b-32768",
	}

	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", GROQAPIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("failed to get a valid response: " + res.Status)
	}

	var response ChatCompletionResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", errors.New("no choices received in the response")
	}

	return response.Choices[0].Message.Content, nil
}
