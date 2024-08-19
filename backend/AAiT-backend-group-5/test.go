package main

import (
	"log"

	config "github.com/aait.backend.g5.main/backend/Config"
	infrastructure "github.com/aait.backend.g5.main/backend/Infrastructure"
)

func main(){
	env := config.NewEnv()
	emailConfig := config.NewEmailServer(*env)

	log.Println(env.SMTP_PORT)
	log.Println(env.SMTP_SERVER)
	emailServer := infrastructure.NewEmailService(emailConfig, *env)

	emailServer.SendEmail("dawitya21@gmail.com", "Test", "This is a test email")
}