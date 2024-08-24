package infrastructure

import (
	"astu-backend-g1/config"
	"log"
	"net/smtp"
)

func SendEmail(toEmail string, title string, body string, link string) error {
	config, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	//in route to handle email related confirmation is domain/confirmation/:email/:pwd
	//also make

	message := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>` + title + `</title>
	</head>
	<body>
		<h1>` + title + `</h1>
		<p>` + body + `</p>
		<a href="` + link + `">Click the Link</a>
	</body>
	</html>
	`
	
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	key := config.Email.EmailKey
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("", "abel.wendmu@a2sv.org", key, "smtp.gmail.com")

	port := "587"
	address := host + ":" + port
	messages := []byte(mime+message)

	err = smtp.SendMail(address, auth, "abel.wendmu@a2sv.org", []string{toEmail}, messages)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
