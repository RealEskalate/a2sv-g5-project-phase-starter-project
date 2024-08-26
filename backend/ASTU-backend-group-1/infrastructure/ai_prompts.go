package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Prompts struct {
	Validate           string `json:"validate"`
	Refine             string `json:"refine"`
	RecommendTitle     string `json:"recommend_title"`
	RecommendContent   string `json:"recommend_content"`
	RecommendTags      string `json:"recommend_tags"`
	CheckPromptContent string `json:"check_prompt_content"`
	Summarize          string `json:"summarize_blog"`
}

func LoadPrompt(path string) (Prompts, error) {
	file, err := os.Open(path)
	if err != nil {
		return Prompts{}, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return Prompts{}, err
	}
	prompts := Prompts{}
	err = json.Unmarshal(data, &prompts)
	if err != nil {
		return Prompts{}, fmt.Errorf("unmarshaling json: %v", err)
	}
	return prompts, nil
}
