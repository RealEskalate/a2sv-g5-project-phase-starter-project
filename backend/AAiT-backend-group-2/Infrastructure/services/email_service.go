package services

import (
	domain "AAiT-backend-group-2/Domain"
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"net/url"
	"path/filepath"
	"time"
)

type EmailService interface {
	SendEmail(receiver, subject, templateName string, model interface{}) error
	GeneratePasswordResetTemplate(email, userName, token string) (domain.PasswordResetModel, error)
}


type emailService struct {
	Host       string
	Port       string
	SenderEmail    string
	SenderPassword string
	FrontendHost string
	TemplateDir   string
}

func NewEmailService(emailHost string, emailPort string, senderEmail string, senderPassword string) EmailService {
	return &emailService{
		Host:		   emailHost,
		Port:		   emailPort,
		SenderEmail:   senderEmail,
		SenderPassword: senderPassword,
		FrontendHost:  "http://localhost:3000/",
		TemplateDir:   "Infrastructure/templates",
	}
}

func (service *emailService) SendEmail(receiver, subject, templateName string, model interface{}) error {
	templ, err := template.ParseFiles(filepath.Join(service.TemplateDir, templateName))

	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := templ.Execute(&body, model); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", service.SenderEmail, service.SenderPassword, service.Host)

	msg := fmt.Sprintf(
		"From: %s\nTo: %s\nSubject: %s\nMIME-Version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\nDate: %s\n\n%s",
		service.SenderEmail,
		receiver,
		subject,
		time.Now().Format(time.RFC1123Z),
		body.String(),
	)
	

	err = smtp.SendMail(service.Host+":"+service.Port, auth, service.SenderEmail, []string{receiver}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func (service *emailService) GeneratePasswordResetTemplate(email, userName, token string) (domain.PasswordResetModel, error) {
	resetLink := service.FrontendHost + "auth/reset-password?email=" + url.QueryEscape(email) + "&token=" + url.QueryEscape(token)
	resetModel := domain.PasswordResetModel{
		UserName:  userName,
		ResetLink: resetLink,
	}

	return resetModel, nil
}