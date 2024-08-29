package Usecases

import (
    "context"
    "fmt"
    "os"
    "strings"

    "github.com/google/generative-ai-go/genai"
    "github.com/joho/godotenv"
    "google.golang.org/api/option"
)

func Chat(message string) (string, error) {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        return "", fmt.Errorf("error loading .env file: %v", err)
    }

    apiKey := os.Getenv("GOOGLE_API_KEY")
    if apiKey == "" {
        return "", fmt.Errorf("API key not found")
    }

    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        return "", fmt.Errorf("failed to create Gemini client: %v", err)
    }

    model := client.GenerativeModel("gemini-1.5-flash")

    // Send a text prompt to the Gemini model
    resp, err := model.GenerateContent(ctx, genai.Text(message))
    if err != nil {
        return "", fmt.Errorf("failed to generate content: %v", err)
    }

    // Prepare the response string
    var responseBuilder strings.Builder
    for i, candidate := range resp.Candidates {
        responseBuilder.WriteString(fmt.Sprintf("Candidate %d:\n%s\n", i+1, candidate.Content))
    }

    // Return the concatenated response as a single string
    return responseBuilder.String(), nil
}
