package infrastructure

import (
    "context"
    "encoding/json"
    "fmt"
    "os"

    "github.com/google/generative-ai-go/genai"
    "github.com/joho/godotenv"
    "google.golang.org/api/option"
)

type AI struct {}

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

func (AI)GenerateContentFromGemini(title string, description string) (string, error) {
    err := godotenv.Load()
    if err != nil {
        return "", err
    }

    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        return "", fmt.Errorf("API key not found in environment variables")
    }

    client, err := genai.NewClient(context.TODO(), option.WithAPIKey(apiKey))
    if err != nil {
        return "", fmt.Errorf("error initializing Gemini client: %v", err)
    }
    defer client.Close()

    model := client.GenerativeModel("gemini-pro")

    prompt := []genai.Part{
        genai.Text(fmt.Sprintf("Generate a blog post with the title '%s' and the following description: '%s'", title, description)),
    }


    resp, err := model.GenerateContent(context.TODO(), prompt...)
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