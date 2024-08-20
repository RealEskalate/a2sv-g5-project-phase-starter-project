package infrastructures

import (
	"fmt"
	"net/smtp"
	"os"

	"aait.backend.g10/domain"
)

func (uc *Infranstructure) SendResetEmail(user *domain.User, resetToken string) error {
	// var app_password = os.Getenv("SMPT_APP_PASS")
	var app_email = os.Getenv("SMPT_APP_EMAIL")
	resetLink := "https://locahost:8080/reset-password?token=" + resetToken

	from := app_email

	to := user.Email
	fmt.Println(to, from)
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
	auth := smtp.PlainAuth("", "8fba770eab09f1", "95295f54db5e69", "sandbox.smtp.mailtrap.io")
	err := smtp.SendMail("sandbox.smtp.mailtrap.io:587", auth, "8fba770eab09f1", []string{to}, []byte(message))
	fmt.Println(err)
	return err
}
