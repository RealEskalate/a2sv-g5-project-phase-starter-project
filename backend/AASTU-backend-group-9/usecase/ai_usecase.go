package usecase

import (
    "context"
    "blog/domain"
    "net/http"
    "bytes"
    "encoding/json"
    "errors"
    "io"
    "time"
)

type aiUsecase struct {
    apiKey         string
    contextTimeout time.Duration
}

func NewAIUsecase(apiKey string, timeout time.Duration) domain.AIUsecase {
    return &aiUsecase{
        apiKey:         apiKey,
        contextTimeout: timeout,
    }
}

func (au *aiUsecase) GenerateBlogContent(ctx context.Context, keywords string) (string, error) {
    ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
    defer cancel()

    requestBody, err := json.Marshal(map[string]string{
        "prompt": keywords,
        "max_tokens": "150",
    })
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci-codex/completions", bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+au.apiKey)

    client := &http.Client{}
	req = req.WithContext(ctx)
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var result map[string]interface{}
    json.Unmarshal(body, &result)

    if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
        if text, ok := choices[0].(map[string]interface{})["text"].(string); ok {
            return text, nil
        }
    }

    return "", errors.New("failed to generate content")
}