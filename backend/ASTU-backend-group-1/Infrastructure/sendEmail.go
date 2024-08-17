package infrastructure

import (
	"astu-backend-g1/config"
	"log"
	"net/smtp"
)

func SendEmail(toEmail string, title string, body string, hashedpwd string) error {
	config, err := config.LoadConfig()
	log.Println(config)
	if err != nil {
		log.Println(err)
	}
	log.Println(config.Email.EmailKey, config.Port)

	//in route to handle email reliated comfirmation is domain/confirmation/:email/:pwd
	//also make
	link := "http://localhost:8000/confirmation/email/" + toEmail + "/pwd/" + hashedpwd + "/"
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
	key := config.Email.EmailKey
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("", "abel.wendmu@a2sv.org", key, "smtp.gmail.com")

	port := "587"
	address := host + ":" + port
	messages := []byte(message)

	err = smtp.SendMail(address, auth, "abel.wendmu@a2sv.org", []string{toEmail}, messages)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
