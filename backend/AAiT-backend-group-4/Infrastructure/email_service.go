package infrastructure

import (
	"aait-backend-group4/Domain"
	"fmt"
	"gopkg.in/gomail.v2"
)

type emailService struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

func NewEmailService(host string, port int, user, password string) domain.EmailService {
	return &emailService{
		SMTPHost:     host,
		SMTPPort:     port,
		SMTPUser:     user,
		SMTPPassword: password,
	}
}

func (s *emailService) SendPasswordResetEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.SMTPUser)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Request")
	resetLink := fmt.Sprintf("https://blog.com/reset-password?token=%s", token)
	m.SetBody("text/plain", fmt.Sprintf("Click the following link to reset your password: %s", resetLink))

	d := gomail.NewDialer(s.SMTPHost, s.SMTPPort, s.SMTPUser, s.SMTPPassword)

	return d.DialAndSend(m)
}
