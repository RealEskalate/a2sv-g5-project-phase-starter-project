package infrastructure

import (
	"astu-backend-g1/infrastructure"
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Prompts struct {
	Validate           string `json:"validate"`
	Refine             string `json:"refine"`
	RecommendTitle     string `json:"recommend_title"`
	RecommendContent   string `json:"recommend_content"`
	RecommendTags      string `json:"recommend_tags"`
	CheckPromptContent string `json:"check_prompt_content"`
	Summarize          string `json:"summarize"`
}
type GeminiModel struct {
	model   *genai.GenerativeModel
	prompts Prompts
}

func connectToGemini(apiKey, modelName string, ctx context.Context) (*genai.GenerativeModel, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return &genai.GenerativeModel{}, err
	}
	return client.GenerativeModel(modelName), nil
}

func NewGeminiModel(apiKey, modelName string, prompts Prompts) (*GeminiModel, error) {
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
	if strings.ToLower(resp) == "no" {
		return fmt.Errorf("prompt not about blogging")
	}
	return nil
}

func (g *GeminiModel) Refine(content string) (string, error) {
	prompt := fmt.Sprintf(g.prompts.Refine, content)
	refinedContent, err := g.SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return refinedContent, nil
}

func (g *GeminiModel) Validate(content string) error {
	prompt := fmt.Sprintf(g.prompts.Validate, content)
	validation, err := g.SendPrompt(prompt)
	if err != nil {
		return err
	}
	if validation != "yes" {
		return fmt.Errorf(validation)
	}
	return nil
}

func (g *GeminiModel) Recommend(data infrastructure.Data, opt string) (infrastructure.RecommendationResponse, error) {
	switch opt {
	case "content":
		content, err := g.recommendContent(data.Title, data.Tags)
		return infrastructure.RecommendationResponse{Content: content}, err
	case "title":
		recommendedTitle, err := g.recommendTitle(data.Content)
		return infrastructure.RecommendationResponse{Title: recommendedTitle}, err
	case "tag":
		tags, err := g.recommendTags(data.Title, data.Content)
		return infrastructure.RecommendationResponse{Tags: tags}, err
	}
	return infrastructure.RecommendationResponse{}, nil
}

func (g *GeminiModel) Summarize(data infrastructure.Data) (string, error) {
	blog := fmt.Sprintf("Title: %v, Content %v, Tags %v", data.Title, data.Content, data.Tags)
	prompt := fmt.Sprintf(g.prompts.Summarize, blog)
	summary, err := g.SendPrompt(prompt)
	return summary, err
}

func (g *GeminiModel) recommendTitle(content string) (string, error) {
	prompt := fmt.Sprintf(g.prompts.RecommendTitle, content)
	if err := g.CheckPromptContent(prompt); err != nil {
		return "", err
	}
	recommendedTitles, err := g.SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return recommendedTitles, nil
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
	prompt := fmt.Sprintf(g.prompts.RecommendContent, title, content)
	if err := g.CheckPromptContent(prompt); err != nil {
		return []string{}, err
	}
	t, err := g.SendPrompt(prompt)
	tags := strings.Split(t, ",")
	if err != nil {
		return []string{}, err
	}
	return tags, nil
}
