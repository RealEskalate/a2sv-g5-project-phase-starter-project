package emailutil

import (
	"fmt"
	"net/smtp"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
)

func SendVerificationEmail(recipientEmail string, VerificationToken string, env *bootstrap.Env) error {
	// Email configuration
	from := env.SenderEmail
	password := env.SenderPassword
	smtpHost := env.SmtpHost
	smtpPort := env.SmtpPort

	subject := "Subject: Account Verification\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	url := fmt.Sprintf("http://localhost:8080/verify-email/%v",VerificationToken)
	body := fmt.Sprintf(
		"<html><body><p>Please click the link below to verify your account:</p><p><a href=\"%s\">Verify Email</a></p><p>Thank you for registering!</p></body></html>",
		url,
	)
	message := []byte(subject + "\n" + mime+body)
	auth := smtp.PlainAuth("", from, password, smtpHost)


	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipientEmail}, message)
	if err != nil {
		return err
	}
	return nil

}