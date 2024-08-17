package config

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

func getPassword() string {
	godotenv.Load(".env")
	return os.Getenv("EMAIL_PASSWORD")
}

func SendEmail(to, subject, body string) error {
	e := email.NewEmail()
	e.From = "eyouel.melkamu@a2sv.org"
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	err := e.SendWithTLS("smtp.gmail.com:465",
		smtp.PlainAuth("", "eyouel.melkamu@a2sv.org", getPassword(), "smtp.gmail.com"),
		&tls.Config{ServerName: "smtp.gmail.com"},
	)
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
