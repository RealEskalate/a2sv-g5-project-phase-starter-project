package config

import (
	"blogs/bootstrap"
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func getPassword() (string, error) {
	return bootstrap.GetEnv("EMAIL_PASSWORD")
}

func SendEmail(to, subject, body string) error {
	e := email.NewEmail()
	e.From = "eyouel.melkamu@a2sv.org"
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	password, err := getPassword()
	if err != nil {
		return err
	}

	err = e.SendWithTLS("smtp.gmail.com:465",
		smtp.PlainAuth("", "eyouel.melkamu@a2sv.org", password, "smtp.gmail.com"),
		&tls.Config{ServerName: "smtp.gmail.com"},
	)
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
