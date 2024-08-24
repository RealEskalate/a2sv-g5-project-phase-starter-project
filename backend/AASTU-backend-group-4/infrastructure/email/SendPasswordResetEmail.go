package email

import (
	"context"
	"fmt"
	"net/smtp"
)

func (e *emailService) SendPasswordResetEmail(ctx context.Context, email, resetLink string) error {
	// Set up the email message
	subject := "Password Reset Request"
	body := fmt.Sprintf(
		"Hi,\n\nTo reset your password, please click the following link:\n\n%s\n\nIf you did not request this, please ignore this email.\n\nThanks,\nYour Company",
		resetLink,
	)
	message := fmt.Sprintf(
		"From: %s\nTo: %s\nSubject: %s\n\n%s",
		e.FromAddress, email, subject, body,
	)

	// Set up authentication information.
	auth := smtp.PlainAuth("", e.SMTPUser, e.SMTPPassword, e.SMTPServer)

	// Send the email.
	return smtp.SendMail(
		fmt.Sprintf("%s:%s", e.SMTPServer, e.SMTPPort),
		auth,
		e.FromAddress,
		[]string{email},
		[]byte(message),
	)
}
