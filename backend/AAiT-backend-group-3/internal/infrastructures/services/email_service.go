package services

import (
	"log"

	"gopkg.in/gomail.v2"
)

type IEmailService interface {
    SendEmail(to string, subject string, body string) error
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
	m := gomail.NewMessage()
	m.SetHeader("From", e.username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.smtpHost, e.smtpPort, e.username, e.password)
	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}
	return nil
}

