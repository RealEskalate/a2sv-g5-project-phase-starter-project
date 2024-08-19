package infrastructure

import (
	"net/smtp"
)

type EmailService struct {
    smtpServer string
    email      string
    password   string
}

func NewEmailService(smtpServer string, email string, password string) *EmailService {
    return &EmailService{
        smtpServer: smtpServer,
        email: email,
        password: password,
    }
}

func (e *EmailService) SendEmail(to, name, code string) error {
	
    msg := "From: " + e.email + "\n" +
        "To: " + to + "\n" +
        "Subject: " + "Vefification Mail" + "\n\n" +
        "Your verification code is" + "\n\n" +
		code

    auth := smtp.PlainAuth(
        "",
        e.email,
        e.password,
        "smtp.gmail.com",
    )

    err := smtp.SendMail(
        e.smtpServer,
        auth,
        e.email,
        []string{to},
        []byte(msg),

    )

    return err
}
