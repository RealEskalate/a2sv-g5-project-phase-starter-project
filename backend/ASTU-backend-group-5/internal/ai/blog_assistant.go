package ai

import (
	"blogApp/internal/config"
	"blogApp/internal/domain"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetAiBlog(chat_id, query string) (domain.Blog, error) {

	conf, err := config.Load()
	if err != nil {
		return domain.Blog{}, err
	}

	endpoint := conf.AI_API_DOMAIN + "/blog_assistant/"
	data := map[string]string{
		"query":   query,
		"chat_id": chat_id,
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return domain.Blog{}, err
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return domain.Blog{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Blog{}, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return domain.Blog{}, err
	}
	return domain.Blog{
		Title:   result["title"].(string),
		Content: result["content"].(string),
	}, nil

}
