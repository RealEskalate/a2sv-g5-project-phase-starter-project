package email

import (
	"fmt"
	"net/smtp"
)

type SimpleEmailSender struct {
	smtpHost     string
	smtpPort     string
	senderEmail  string
	senderPasswd string
}

func NewSimpleEmailSender(smtpHost, smtpPort, senderEmail, senderPasswd string) *SimpleEmailSender {
	return &SimpleEmailSender{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		senderEmail:  senderEmail,
		senderPasswd: senderPasswd,
	}
}

func (s *SimpleEmailSender) SendVerificationEmail(userEmail string, token string) error {
	subject := "Subject: Verify Your Email\n"
	body := fmt.Sprintf("Please verify your email by clicking the link: https://yourapp.com/verify-email?token=%s", token)
	message := subject + "\n" + body

	return s.sendMail(userEmail, message)
}

func (s *SimpleEmailSender) SendPasswordResetEmail(userEmail string, token string) error {
	subject := "Subject: Reset Your Password\n"
	body := fmt.Sprintf("You can reset your password using the following link: https://yourapp.com/reset-password?token=%s", token)
	message := subject + "\n" + body

	return s.sendMail(userEmail, message)
}

func (s *SimpleEmailSender) sendMail(to string, message string) error {
	auth := smtp.PlainAuth("", s.senderEmail, s.senderPasswd, s.smtpHost)
	err := smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, s.senderEmail, []string{to}, []byte(message))
	return err
}
