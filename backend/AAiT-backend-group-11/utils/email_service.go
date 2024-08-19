package utils

import (
	"backend-starter-project/domain/interfaces"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
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
	ctx := context.Background()

	b, err := os.ReadFile("credentials.json")
	// Unable to read client secret file
	if err != nil {
		return err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	// Unable to parse client secret file to config
	if err != nil {
		return err
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	// Unable to retrieve Gmail client
	if err != nil {
		return err
	}

	// Create MIME message with HTML content
	email := &gmail.Message{
		Raw: base64.URLEncoding.EncodeToString([]byte("To: " + emailAddress + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n" +
			"\r\n" +
			"<!DOCTYPE html>" + body)),
	}

	_, err = srv.Users.Messages.Send("me", email).Do()
	// Error trying to send the message
	if err != nil {
		return err
	}

	return nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
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