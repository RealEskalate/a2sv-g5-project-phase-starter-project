package emailservice

import (
	"blogapp/Config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	fromEmail = "mailtrap@demomailtrap.com"
	fromName  = "Mailtrap Blog Test"
)

// MailTrapService defines methods for interacting with the Mailtrap API
type MailTrapService interface {
	SendEmail(toEmail string, subject string, text string, category string) error
	// GetEmailMessages(emailAddress string) ([]byte, error)
	// DeleteEmailMessage(emailAddress, messageID string) error
}

// mailTrapService implements the MailTrapService interface
type mailTrapService struct {
	apiToken string
}

// NewMailTrapService creates a new instance of mailTrapService
func NewMailTrapService() *mailTrapService {
	return &mailTrapService{
		apiToken: Config.Mail_TRAP_API_KEY,
	}
}

// SendEmail sends an email using the Mailtrap API
func (s *mailTrapService) SendEmail(toEmail string, subject string, text string, category string) error {
	url := "https://send.api.mailtrap.io/api/send"
	headers := map[string]string{
		"Authorization": "Bearer " + s.apiToken,
		"Content-Type":  "application/json",
	}

	emailPayload := map[string]interface{}{
		"from": map[string]string{
			"email": fromEmail,
			"name":  fromName,
		},
		"to": []map[string]string{
			{
				"email": toEmail,
			},
		},
		"subject":  subject,
		"text":     text,
		"category": category,
	}

	jsonPayload, err := json.Marshal(emailPayload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error sending email: %s", resp.Status)
	}
	return nil
}

// Optionally implement GetEmailMessages and DeleteEmailMessage here
