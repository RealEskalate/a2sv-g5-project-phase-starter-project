package services

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
)

type IEmailService interface {
	SendResetEmail(to, link string) error
	SendVerificationEmail(to, link string) error
}

type EmailService struct {
	smtpHost string
	smtpPort int
	username string
	password string
}

func NewEmailService(smtpHost string, smtpPort int, username, password string) IEmailService {
	return &EmailService{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		username: username,
		password: password,
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", "artistrynexus.ov.server@gmail.com", "qbmx duiu brjw fjgx" , e.smtpHost)

	from := e.username
	toList := []string{to}
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", to, subject, body))

	addr := fmt.Sprintf("%s:%d", e.smtpHost, e.smtpPort)
	err := smtp.SendMail(addr, auth, from, toList, msg)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailService) SendResetEmail(to, link string) error {
	templatePath := filepath.Join("../internal/infrastructures/services/templates", "reset_password.html")
	body, err := parseTemplate(templatePath, link)
	if err != nil {
		return err
	}
	return e.SendEmail(to, "Reset Your Password", body)
}

func (e *EmailService) SendVerificationEmail(to, link string) error {
	templatePath := filepath.Join("../internal/infrastructures/services/templates", "verify_email.html")
	body, err := parseTemplate(templatePath, link)
	if err != nil {
		return err
	}
	return e.SendEmail(to, "Verify Your Email", body)
}

func parseTemplate(templatePath, link string) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	data := map[string]string{
		"Link": link,
	}

	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}
	return body.String(), nil
}
