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

// SendEmail sends an email with the specified content type (text or HTML)
func SendEmail(to, subject, body string, isHTML bool) error {
	e := email.NewEmail()
	e.From = "eyouel.melkamu@a2sv.org"
	e.To = []string{to}
	e.Subject = subject

	if isHTML {
		e.HTML = []byte(body) // Set HTML content
	} else {
		e.Text = []byte(body) // Set plain text content
	}

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
