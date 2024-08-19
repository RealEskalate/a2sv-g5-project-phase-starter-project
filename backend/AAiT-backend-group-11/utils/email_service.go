package utils

import (
	"backend-starter-project/domain/interfaces"

	"fmt"
	"log"

	"net/smtp"
	"strings"

	"github.com/badoux/checkmail"

)

type emailService struct {
	smtpServer string
	auth       smtp.Auth
	sender     string
}

// VerifyEmailAddress implements interfaces.EmailService.
func (e *emailService) VerifyEmailAddress(emailAddress string) (bool, error) {
    err := checkmail.ValidateHostAndUser(e.smtpServer, e.sender, emailAddress)
    if err != nil {
        if smtpErr, ok := err.(checkmail.SmtpError); ok {
            // Print detailed SMTP error and return false
            log.Printf("SMTP Error - Code: %s, Msg: %s\n", smtpErr.Code(), smtpErr)
            return false, smtpErr
        }
        // Return false and the generic error
        return false, err
    }
    // Email is valid
    return true, nil
}

// GenerateEmailTemplate implements interfaces.EmailService.
func (e *emailService) GenerateEmailTemplate(header string, content string) string {

	tmpl := `
	<html>
	<head>
		<style>
			body {
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				line-height: 1.6;
				background-color: #f8f9fa;
				margin: 0;
				padding: 0;
			}

			.header {
				background-color: #4CAF50;
				padding: 20px;
				text-align: center;
			}

			.header h1 {
				color: #ffffff;
				margin: 0;
				font-size: 32px;
				text-transform: uppercase;
			}

			.content {
				padding: 20px;
				background-color: #ffffff;
				box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
				border-radius: 10px;
				margin: 20px;
			}

			.content p {
				margin-bottom: 15px;
				font-size: 16px;
				color: #333333;
			}

			.content h3 {
				color: #4CAF50;
				margin: 0 0 10px;
				font-size: 22px;
			}

			.footer {
				background-color: #4CAF50;
				padding: 20px;
				text-align: center;
				color: #ffffff;
			}

			.footer p {
				margin: 0;
				font-size: 16px;
				color: #ffffff;
			}
		</style>
	</head>
	<body>
		<div class="header">
			<h1>` + header + `</h1>
		</div>

		<div class="content">
		` + content + `
		</div>

		<div class="footer">
			<p>Read, and Write </p>
			<p>The next medium</p>
		</div>
	</body>
	</html>
`
	return tmpl
}

// SendEmail implements interfaces.EmailService.
func (e *emailService) SendEmail(emailAddress string, subject string, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"haloitisme0912@gmail.com", // Replace with your email
		"btnb soyo xqpm ooxw",      // Replace with your app-specific password
		"smtp.gmail.com",
	)

	// Create the email message.
	msg := "To: " + emailAddress + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n" +
		"\r\n" + body

	// Send the email.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"haloitisme0912@gmail.com", // Replace with your email
		[]string{emailAddress},
		[]byte(msg),
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil

}

// NewEmailService creates a new instance of emailService.
func NewEmailService(smtpServer, password, sender string) interfaces.EmailService {
	auth := smtp.PlainAuth("", sender, password, smtpServer[:strings.Index(smtpServer, ":")])
	return &emailService{
		smtpServer: smtpServer,
		auth:       auth,
		sender:     sender,
	}
}