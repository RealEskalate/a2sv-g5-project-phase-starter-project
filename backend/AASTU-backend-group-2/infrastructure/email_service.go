package infrastructure

import (
	"fmt"
	"net/url"
	"strconv"

	"gopkg.in/gomail.v2"
)

// EmailConfig holds the configuration required to send emails.
type EmailConfig struct {
	SMTPHost       string
	SMTPPort       int
	SenderEmail    string
	SenderPassword string
}

// NewEmailConfig initializes and returns a new EmailConfig instance.
func NewEmailConfig() (*EmailConfig, error) {
	port, err := strconv.Atoi(DotEnvLoader("SMTPPORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid SMTP port: %w", err)
	}

	return &EmailConfig{
		SMTPHost:       DotEnvLoader("SMTPHOST"),
		SMTPPort:       port,
		SenderEmail:    DotEnvLoader("SMTPUSER"),
		SenderPassword: DotEnvLoader("SMTPPASS"),
	}, nil
}

// EmailService provides methods to send different types of emails.
type EmailService struct {
	config *EmailConfig
}

// NewEmailService initializes and returns a new EmailService instance.
func NewEmailService(config *EmailConfig) *EmailService {
	return &EmailService{config: config}
}

// SendResetEmail sends the password reset email using gomail.
func (es *EmailService) SendResetEmail(userEmail, resetToken string) error {
	resetURL := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", url.QueryEscape(resetToken))
	body := fmt.Sprintf(
		"Click the following link to reset your password:\n%s\n\n"+
			"If you did not request a password reset, please ignore this email.",
		resetURL)

	return es.sendEmail(userEmail, "Password Reset Request", body)
}

// SendVerificationEmail sends the email verification email using gomail.
func (es *EmailService) SendVerificationEmail(userEmail, verificationToken string) error {
	verificationURL := fmt.Sprintf("http://localhost:8080/user/verify-email?token=%s", verificationToken)
	body := fmt.Sprintf(
		"Click the following link to verify your email:\n%s\n\n"+
			"If you did not sign up for this account, please ignore this email.",
		verificationURL)

	return es.sendEmail(userEmail, "Email Verification", body)
}

// sendEmail is a helper method to send an email with the given subject and body.
func (es *EmailService) sendEmail(toEmail, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("A2SV G55-G2 Blog <%s>", es.config.SenderEmail))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(es.config.SMTPHost, es.config.SMTPPort, es.config.SenderEmail, es.config.SenderPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
