package infrastructure

import (
	"errors"
	"fmt"
	"net/smtp"
)

type EmailService interface {
	SendEmail(from, to, body string) error
}

type Email struct {
	Username string
	Password string
	Host     string
	Port     string
}

func NewEmail(username, password, host, port string) EmailService {
	return &Email{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func (e *Email) SendEmail(from, to, body string) error {
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)
	message := []byte(fmt.Sprintf("This Message of From %s To %s", from, to) +
		"Subject: Account Activation \r\n" +
		fmt.Sprintf("%s\r\n", body))
	addr := fmt.Sprintf("%s:%s", e.Host, e.Port)
	err := smtp.SendMail(addr, auth, from, []string{to}, message)
	if err != nil {
		return errors.New("failed to send an email")
	}
	return nil
}
