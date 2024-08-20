package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// Define a type for the send email function
type SendEmailFunc func(to string, subject string, body string) error

// Declare a global variable for the send email function
var SendEmail SendEmailFunc = defaultSendEmail

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

func defaultSendEmail(emailAddress, subject, body string) error {
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

func GenerateEmailTemplate(header, content string) string {
	tmpl := `
	<html>
	<head>
		<style>
			body {
				font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
				line-height: 1.5;
				background-color: #f4f4f4;
				margin: 0;
				padding: 0;
			}

			.header {
				background-color: #2D298E;
				padding: 20px;
				text-align: center;
			}

			.header h1 {
				color: #ffffff;
				margin: 0;
				font-size: 28px;
			}

			.content {
				padding: 20px;
				background-color: #ffffff;
			}

			.content p {
				margin-bottom: 15px;
				font-size: 18px;
				color: #333333;
			}

			.content h3 {
				color: #2D298E;
				margin: 0 0 10px;
				font-size: 20px;
			}

			.footer {
				background-color: #f9f9f9;
				padding: 20px;
				text-align: center;
			}

			.footer p {
				margin: 0;
				font-size: 16px;
				color: #888888;
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
			<p>Best regards,</p>
		</div>
	</body>
	</html>`

	return tmpl
}
