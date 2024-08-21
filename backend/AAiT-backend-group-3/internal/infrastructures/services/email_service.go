package services

import (
	"fmt"
	"net/smtp"
	//"strings"
)

type IEmailService interface {
    SendResetEmail(to, subject, body string) error
	SendVerificationEmail(to, subject, body string) error
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
	auth := smtp.PlainAuth("", e.username, e.password, e.smtpHost)

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

func (e *EmailService) SendResetEmail(to, subject, body string) error {
	return e.SendEmail(to, subject, body)
}

func (e *EmailService) SendVerificationEmail(to, subject, body string) error {
	return e.SendEmail(to, subject, body)
}
