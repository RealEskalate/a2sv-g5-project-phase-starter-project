package mail_service

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendMail(from string, to string, smtpAddress string, smtpPassword string, mailContent string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", from, smtpAddress)
	e.To = []string{to}
	e.Subject = "Email Verification"
	e.HTML = []byte(mailContent)

	return e.Send("smtp.gmail.com:587", smtp.PlainAuth("", smtpAddress, smtpPassword, "smtp.gmail.com"))
}
