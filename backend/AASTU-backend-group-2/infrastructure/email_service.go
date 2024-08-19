package infrastructure

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

const (
	smtpHost       = "smtp.gmail.com"
	smtpPort       = 587
	senderEmail    = "danielababu0966@gmail.com"
	senderPassword = "toyweqxpdatwhlhf"
)

// Sends the password reset email using gomail
func sendResetEmail(userEmail, resetToken string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Password Reset Request")

	resetURL := fmt.Sprintf("lochalhost:8080/reset-password?token=%s", resetToken)
	body := fmt.Sprintf(
		"Click the following link to reset your password:\n%s\n\n"+
			"If you did not request a password reset, please ignore this email.",
		resetURL)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send reset email: %w", err)
	}

	return nil
}
