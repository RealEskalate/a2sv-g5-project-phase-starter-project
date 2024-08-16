package infrastructure


import (
	"ASTU-backend-group-3/Blog_manager/Delivery/config"
	"fmt"
	"net/smtp"
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
	auth := smtp.PlainAuth("", es.smtpConfig.Username, es.smtpConfig.Password, es.smtpConfig.Host)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body + "\r\n")

	err := smtp.SendMail(es.smtpConfig.Host+":"+es.smtpConfig.Port, auth, es.smtpConfig.From, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
