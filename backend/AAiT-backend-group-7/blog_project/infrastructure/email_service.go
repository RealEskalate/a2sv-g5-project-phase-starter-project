
package infrastructure

import (
	"fmt"
	"net/smtp"
)

type EmailService struct {
	smtpHost string
	smtpPort string
	username string
	password string
}

func NewEmailService(smtpHost, smtpPort, username, password string) *EmailService {
	return &EmailService{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		username: username,
		password: password,
	}
}

func (service *EmailService) SendPasswordResetEmail(email, resetToken string) error {
	from := service.username
	to := []string{email}
	subject := "Password Reset Request"
	body := fmt.Sprintf("Click on the following link to reset your password: \n\nhttp://localhost:8080/reset-password?token=%s\n", resetToken)

	msg := "From: " + from + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", service.username, service.password, service.smtpHost)

	err := smtp.SendMail(service.smtpHost+":"+service.smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Password reset email sent to:", email)
	return nil
}
