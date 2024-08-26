package gemini

import (
	"astu-backend-g1/infrastructure"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiModel struct {
	model   *genai.GenerativeModel
	prompts infrastructure.Prompts
}

func connectToGemini(apiKey, modelName string, ctx context.Context) (*genai.GenerativeModel, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return &genai.GenerativeModel{}, err
	}
	return client.GenerativeModel(modelName), nil
}

func NewGeminiModel(apiKey, modelName string, prompts infrastructure.Prompts) (*GeminiModel, error) {
	model, err := connectToGemini(apiKey, modelName, context.Background())
	if err != nil {
		return &GeminiModel{}, err
	}
	return &GeminiModel{model: model, prompts: prompts}, nil
}

func (g *GeminiModel) SendPrompt(prompt string) (string, error) {
	resp, err := g.model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("no response from the model")
	}
	candidate := resp.Candidates[0]
	responseText := fmt.Sprint(candidate.Content.Parts[0])
	return responseText, nil
}

func (g *GeminiModel) CheckPromptContent(content string) error {
	p := fmt.Sprintf(g.prompts.CheckPromptContent, content)
	resp, err := g.SendPrompt(p)
	if err != nil {
		return fmt.Errorf("getting response error: %v", err.Error())
	}
	if strings.ToLower(resp[:len(resp)-2]) == "no" {
		return fmt.Errorf("prompt not about blogging")
	}
	return nil
}

func (g *GeminiModel) Refine(data infrastructure.Data) (infrastructure.Data, error) {
	prompt := fmt.Sprintf(g.prompts.Refine, data.Content, data.Title, data.Tags)
	jsonString, err := g.SendPrompt(prompt)
	if err != nil {
		return infrastructure.Data{}, err
	}
	jsonString = strings.TrimPrefix(jsonString, "```json\n")
	jsonString = strings.TrimSuffix(jsonString, "\n```")
	var d infrastructure.Data
	err = json.Unmarshal([]byte(jsonString), &d)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	if err != nil {
		return infrastructure.Data{}, err
	}
	return d, nil
}

func (g *GeminiModel) Validate(data infrastructure.Data) error {
	prompt := fmt.Sprintf(g.prompts.Validate, data.Content)
	validation, err := g.SendPrompt(prompt)
	if err != nil {
		return err
	}
	if validation != "yes" {
		return fmt.Errorf("%v", validation)
	}
	return nil
}

func (g *GeminiModel) Recommend(data infrastructure.Data, opt string) (interface{}, error) {
	switch opt {
	case "content":
		content, err := g.recommendContent(data.Title, data.Tags)
		return map[string]string{"content": content}, err
	case "title":
		return g.recommendTitle(data.Content, data.Tags)
	case "tags":
		tags, err := g.recommendTags(data.Title, data.Content)
		return tags, err
	}
	return []string{}, nil
}

func (g *GeminiModel) Summarize(data infrastructure.Data) (string, error) {
	blog := fmt.Sprintf("Title: %v, Content %v, Tags %v", data.Title, data.Content, data.Tags)
	prompt := fmt.Sprintf(g.prompts.Summarize, blog)
	summary, err := g.SendPrompt(prompt)
	return summary[:len(summary)-2], err
}

func (g *GeminiModel) recommendTitle(content string, tags []string) ([]string, error) {
	prompt := fmt.Sprintf(g.prompts.RecommendTitle, content, strings.Join(tags, ", "))
	if err := g.CheckPromptContent(prompt); err != nil {
		return []string{}, err
	}
	resp, err := g.SendPrompt(prompt)
	titles := strings.Split(resp, "\n")
	if err != nil {
		return []string{}, err
	}
	return titles[:len(titles)-1], nil
}

func (g *GeminiModel) Chat(content string) (string, error) {
	if err := g.CheckPromptContent(content); err != nil {
		return "", err
	}
	resp, err := g.SendPrompt(content)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (g *GeminiModel) recommendContent(title string, tags []string) (string, error) {
	prompt := fmt.Sprintf(g.prompts.RecommendContent, title, tags)
	if err := g.CheckPromptContent(prompt); err != nil {
		return "", err
	}
	recommendedContent, err := g.SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return recommendedContent, nil
}

func (g *GeminiModel) recommendTags(title string, content string) ([]string, error) {
	prompt := fmt.Sprintf(g.prompts.RecommendTags, title, content)
	if err := g.CheckPromptContent(prompt); err != nil {
		return []string{}, err
	}
	t, err := g.SendPrompt(prompt)
	tags := strings.Split(t[:len(t)-2], ", ")
	if err != nil {
		return []string{}, err
	}
	return tags, nil
}
