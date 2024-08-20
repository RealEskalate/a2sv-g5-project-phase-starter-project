package ai

import (
	"blogApp/internal/config"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func extractTextFromContent(content []interface{}) string {
	var textContent strings.Builder

	for _, item := range content {
		if str, ok := item.(string); ok {
			textContent.WriteString(str)
			textContent.WriteString("\n")
		}
	}

	return textContent.String()
}

func ModerateBlog(blog_content []interface{}, blog_title string) (bool, string, error) {
	conf, err := config.Load()
	if err != nil {
		return false, "", err
	}
	endpoint := conf.AI_API_DOMAIN + "/validate_post/"

	blogString := extractTextFromContent(blog_content)
	data := map[string]string{
		"content": blogString,
		"title":   blog_title,
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return false, "", err
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, "", err
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, "", err
	}
	return result["is_valid"].(bool), result["message"].(string), nil
}
