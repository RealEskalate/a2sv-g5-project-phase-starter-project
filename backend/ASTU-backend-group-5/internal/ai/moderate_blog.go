package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func ModerateBlog(blog_content string, blog_title string) (int, string, error) {
	// conf, err := config.Load()
	// if err != nil {
	// 	return false, "", errors.New("failed to load config: " + err.Error())
	// }

	endpoint := "http://127.0.0.1:8000/validate_post/" //conf.AI_API_DOMAIN + "/validate_post/"
	log.Printf("ModerateBlog: using endpoint %s", endpoint)

	data := map[string]string{
		"content": blog_content,
		"title":   blog_title,
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Printf("ModerateBlog: failed to marshal data: %v", err)
		return 100, "", errors.New("failed to marshal data to json: " + err.Error())
	}

	log.Printf("ModerateBlog: sending request with data %s", string(marshal))

	client := &http.Client{Timeout: 80 * time.Second}
	resp, err := client.Post(endpoint, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		log.Printf("ModerateBlog: failed to send request: %v", err)
		return 100, "", errors.New("failed to post request: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ModerateBlog: failed to read response body: %v", err)
		return 100, "", err
	}

	log.Printf("ModerateBlog: received response %s", string(body))

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("ModerateBlog: failed to unmarshal response: %v", err)
		return 100, "", err
	}

	// Ensure these types are correct in your JSON response
	gradeFloat, ok1 := result["grade"].(float64)
	message, ok2 := result["message"].(string)
	if !ok1 || !ok2 {
		log.Printf("ModerateBlog: unexpected response structure: %v", result)
		return 100, "", errors.New("unexpected response structure")
	}

	grade := int(gradeFloat) // Convert float64 to int

	return grade, message, nil
}
