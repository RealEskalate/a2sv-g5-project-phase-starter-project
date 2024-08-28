package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
)

type emailService struct {
	auth         smtp.Auth
	templatePath string
}

func NewEmailService(templatePath string) domain.EmailService {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
		"smtp.gmail.com",
	)
	return &emailService{auth: auth, templatePath: templatePath}
}

func (service *emailService) SendMail(to, subject, templateName string, body interface{}) error {
	from := os.Getenv("SMTP_EMAIL")
	tmplt, errLoadingTmplt := template.ParseFiles(service.templatePath + templateName)
	if errLoadingTmplt != nil {
		return fmt.Errorf("error loading the template: %v", errLoadingTmplt)
	}

	var bodyWritten bytes.Buffer
	if errBuffer := tmplt.Execute(&bodyWritten, body); errBuffer != nil {
		return fmt.Errorf("error excuting template :%w", errBuffer)
	}
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", from, to, subject, bodyWritten.String())
	errSmtp := smtp.SendMail("smtp.gmail.com:587", service.auth, from, []string{to}, []byte(msg))
	if errSmtp != nil {
		return fmt.Errorf("error sending email: %w", errSmtp)
	}
	return nil
}

func (service *emailService) SendVerificationEmail(to, name, verificationLink string) error {
	data := map[string]string{
		"Name":             name,
		"VerificationLink": verificationLink,
	}

	return service.SendMail(to, "Email Verification", "verification.html", data)
}

func (service *emailService) SendPasswordResetEmail(to, name, resetLink, resetCode string) error {
	data := map[string]string{
		"Name":      name,
		"ResetLink": resetLink,
		"ResetCode": resetCode,
	}

	return service.SendMail(to, "Password Reset", "password_reset.html", data)
}
