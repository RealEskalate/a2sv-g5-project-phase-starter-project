package infrastructure

import (
	"ASTU-backend-group-3/Blog_manager/Delivery/config"
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type EmailService struct {
	smtpConfig config.SMTPConfig
}

func NewEmailService() *EmailService {
	return &EmailService{
		smtpConfig: config.LoadSMTPConfig(),
	}
}
func (es *EmailService) SendEmail(to, subject, body string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load environment variables for email: %w", err)
	}
	es.smtpConfig.Host = os.Getenv("SMTP_HOST")
	es.smtpConfig.Port = os.Getenv("SMTP_PORT")
	es.smtpConfig.Username = os.Getenv("SMTP_USERNAME")
	es.smtpConfig.Password = os.Getenv("SMTP_PASSWORD")
	es.smtpConfig.From = os.Getenv("SMTP_FROM")
	// SMTP server address format should include the hostname only, not the port
	auth := smtp.PlainAuth("", es.smtpConfig.Username, es.smtpConfig.Password, es.smtpConfig.Host)

	// The message should include the From header
	msg := []byte("From: " + es.smtpConfig.From + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body + "\r\n")

	// SMTP server address format should include the hostname and port separated by a colon
	err = smtp.SendMail(es.smtpConfig.Host+":"+es.smtpConfig.Port, auth, es.smtpConfig.From, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
