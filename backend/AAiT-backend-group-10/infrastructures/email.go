package infrastructures

import (
	"fmt"
	"net/smtp"

	"aait.backend.g10/domain"
)

type EmailService struct {
	AppEmail    string
	AppUsername string
	AppPass     string
	AppHost     string
}

func (s *EmailService) SendResetEmail(email string, resetToken string) *domain.CustomError {
	resetLink := "https://localhost:8080/reset-password?token=" + resetToken

	from := s.AppEmail
	to := email

	subject := "Password Reset Request"
	body := `
		<html>
		<head>
			<style>
				.container {
					font-family: Arial, sans-serif;
					line-height: 1.6;
					color: #333;
				}
				.button {
					display: inline-block;
					padding: 10px 20px;
					margin: 20px 0;
					background-color: #28a745;
					color: #ffffff;
					text-decoration: none;
					border-radius: 5px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Password Reset Request</h1>
				<p>Hello,</p>
				<p>We received a request to reset your password. Click the button below to reset it:</p>
				<a href="` + resetLink + `" class="button">Reset Password</a>
				<p>If you did not request a password reset, please ignore this email.</p>
				<p>Thank you!</p>
			</div>
		</body>
		</html>
	`

	// MIME headers
	message := "MIME-Version: 1.0\r\n"
	message += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	message += "From: " + from + "\r\n"
	message += "To: " + to + "\r\n"
	message += "Subject: " + subject + "\r\n\r\n"
	message += body

	// Example using Gmail's SMTP server
	auth := smtp.PlainAuth("", s.AppUsername, s.AppPass, s.AppHost)
	err := smtp.SendMail(s.AppHost+":587", auth, s.AppUsername, []string{to}, []byte(message))
	if err != nil {
		fmt.Println(err)
		return domain.ErrEmailSendingFailed
	}

	return nil
}

func (s *EmailService) SendActivationEmail(email string, activationToken string) *domain.CustomError {
	verifyEmail := "https://localhost:8080/reset-password?token=" + activationToken

	from := s.AppEmail
	to := email

	subject := "Password Reset Request"
	body := `
		<html>
		<head>
			<style>
				.container {
					font-family: Arial, sans-serif;
					line-height: 1.6;
					color: #333;
				}
				.button {
					display: inline-block;
					padding: 10px 20px;
					margin: 20px 0;
					background-color: #28a745;
					color: #ffffff;
					text-decoration: none;
					border-radius: 5px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Password Reset Request</h1>
				<p>Hello,</p>
				<p>Thankyou for joining us; Click the button below to verify your email:</p>
				<a href="` + verifyEmail + `" class="button">Verify Email</a>
				<p>If you did not request a password reset, please ignore this email.</p>
				<p>Thank you!</p>
			</div>
		</body>
		</html>
	`

	// MIME headers
	message := "MIME-Version: 1.0\r\n"
	message += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	message += "From: " + from + "\r\n"
	message += "To: " + to + "\r\n"
	message += "Subject: " + subject + "\r\n\r\n"
	message += body

	// Example using Gmail's SMTP server
	auth := smtp.PlainAuth("", s.AppUsername, s.AppPass, s.AppHost)
	err := smtp.SendMail(s.AppHost+":587", auth, s.AppUsername, []string{to}, []byte(message))
	if err != nil {
		fmt.Println(err)
		return domain.ErrEmailSendingFailed
	}

	return nil
}
