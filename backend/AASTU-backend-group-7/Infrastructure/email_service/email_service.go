package emailservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	YOUR_SERVICE_ID  = "service_67ycmnl"
	YOUR_TEMPLATE_ID = "template_fh4miqh"
	YOUR_PUBLIC_KEY  = "7MAmpSKJ-QMOiOhFH"
)

type MailService interface {
	SendEmail(toEmail string, subject string, text string, category string) error
}

type mailService struct {
	apiToken string
}

func NewMailService() *mailService {
	return &mailService{}
}
func (s *mailService) SendEmail(toEmail string, subject string, text string, category string) error {
	url := "https://api.emailjs.com/api/v1.0/email/send"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	data := map[string]interface{}{
		"service_id":  YOUR_SERVICE_ID,
		"template_id": YOUR_TEMPLATE_ID,
		"user_id":     YOUR_PUBLIC_KEY,
		"template_params": map[string]string{
			"message":  text,
			"subject":  subject,
			"category": category,
			"to_email": toEmail,
		},
	}

	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var responseBody map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
			return fmt.Errorf("error reading response: %s", resp.Status)
		}
		return fmt.Errorf("error sending email: %s", responseBody["message"])
	}

	return nil
}
