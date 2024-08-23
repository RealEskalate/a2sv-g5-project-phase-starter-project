package infrastructure

import (
	"fmt"
	"net/url"
	"strconv"

	"gopkg.in/gomail.v2"
)

var (
	smtpHost       = DotEnvLoader("SMTPHOST")
	smtpPort, _    = strconv.Atoi(DotEnvLoader("SMTPPORT"))
	senderEmail    = DotEnvLoader("SMTPUSER")
	senderPassword = DotEnvLoader("SMTPPASS")
)

// Sends the password reset email using gomail
func sendResetEmail(userEmail, resetToken string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "A2SV G55-G2 Blog <"+senderEmail+">")
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Password Reset Request")

	// Ensure the token is correctly embedded in the URL
	resetURL := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", url.QueryEscape(resetToken))
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

func SendVerificationEmail(userEmail, verificationToken string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "A2SV G55-G2 Blog <"+senderEmail+">")
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Email Verification")

	verificationURL := fmt.Sprintf("http://localhost:8080/user/verify-email?token=%s", verificationToken)
	body := fmt.Sprintf(
		"Click the following link to verify your email:\n%s\n\n"+
			"If you did not sign up for this account, please ignore this email.",
		verificationURL)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}

	return nil
}
