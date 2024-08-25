package infrastructure

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
)

type EmailService interface {
	SendEmail(from, to, body, subjectRelated string) error
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

func (e *Email) SendEmail(from, to, body, subjectRelated string) error {
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)
	fromMsg := fmt.Sprintf("From: <%s>\r\n", from)
	toMSG := fmt.Sprintf("To: <%s>\r\n", to)
	subject := fmt.Sprintf("Subject: %s\r\n", subjectRelated)

	message := []byte(fromMsg + toMSG + subject + "\r\n" + body)

	addr := fmt.Sprintf("%s:%s", e.Host, e.Port)
	err := smtp.SendMail(addr, auth, from, []string{to}, message)
	if err != nil {
		log.Default().Println("Error while sending email", err)
		return errors.New("failed to send an email")
	}
	return nil
}
