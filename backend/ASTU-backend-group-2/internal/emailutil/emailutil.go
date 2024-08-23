package emailutil

import (
	"fmt"
	"log"
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
	mime := "MIME-Version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	url := fmt.Sprintf("http://localhost:8080/verify-email/%v", VerificationToken)
	body := Emailtemplate(url)
	message := []byte(subject + mime + "\n" + body)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipientEmail}, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	return nil

}

func SendOtpVerificationEmail(recipientEmail string, otp string, env *bootstrap.Env) error {
	// Email configuration
	from := env.SenderEmail
	password := env.SenderPassword
	smtpHost := env.SmtpHost
	smtpPort := env.SmtpPort

	subject := "Subject: Account Verification\n"
	mime := "MIME-Version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := OTPEmailTemplate(otp,env)
	message := []byte(subject + mime + "\n" + body)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipientEmail}, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	return nil

}