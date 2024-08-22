package infrastructure

import (
	"blog_project/domain"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GeminiService struct {
	ApiKey string
}

func NewGeminiService(apiKey string) domain.AiService {
	return &GeminiService{
		ApiKey: apiKey,
	}
}

func (s *GeminiService) GenerateContent(ctx context.Context, keywords string) (string, error) {
	prompt := fmt.Sprintf("Generate blog content for the following keywords: %s", keywords)

	fmt.Println("Prompt is:", prompt)

	reqBody, err := json.Marshal(map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": prompt,
					},
				},
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=%s", s.ApiKey)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read HTTP response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status code: %d, response body: %s", resp.StatusCode, string(body))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal HTTP response: %w", err)
	}

	// Navigating through the nested response structure
	candidates, ok := response["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		return "", fmt.Errorf("candidates are missing or not an array in the response: %v", response)
	}

	firstCandidate, ok := candidates[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("first candidate is not a map in the response: %v", candidates[0])
	}

	content, ok := firstCandidate["content"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("content is missing or not a map in the first candidate: %v", firstCandidate)
	}

	parts, ok := content["parts"].([]interface{})
	if !ok || len(parts) == 0 {
		return "", fmt.Errorf("parts are missing or not an array in the content: %v", content)
	}

	firstPart, ok := parts[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("first part is not a map in the parts array: %v", parts[0])
	}

	generatedText, ok := firstPart["text"].(string)
	if !ok {
		return "", fmt.Errorf("text is missing or not a string in the first part: %v", firstPart)
	}

	return generatedText, nil
}
