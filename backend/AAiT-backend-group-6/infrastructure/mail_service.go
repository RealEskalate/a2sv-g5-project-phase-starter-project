package infrastructure

import (
	"fmt"
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

func (e *EmailService) EmailVerificationMsg(to, name, code string) string {
    body := fmt.Sprintf("Dear %s,\n\nPlease use the following code to verify your email address:\n\n%s\n\nThank you.", name, code)
    msg := "From: " + e.email + "\n" +
        "To: " + to + "\n" +
        "Subject: " + "Email Verification" + "\n\n" +
        body

	return msg
}

func (e *EmailService) PWRecoveryMsg(to, name, resetLink string) string {
	body := fmt.Sprintf("Dear %s,\n\nWe received a request to reset your password. You can reset it by clicking the link below:\n\n%s\n\nIf you did not request a password reset, please ignore this email.", name, resetLink)

    msg := "From: " + e.email + "\n" +
        "To: " + to + "\n" +
        "Subject: " + "Password Reset Request" + "\n\n" +
        body
        
	return msg
}

func (e *EmailService) SendEmail(to, msg string) error {

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
