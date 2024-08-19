package infrastructure

import (
	"log"
	"strconv"

	"gopkg.in/gomail.v2"

	config "github.com/aait.backend.g5.main/backend/Config"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/asaskevich/govalidator"
)

type emailService struct {
	emailConfig models.EmailConfig
	env         config.Env
}

func NewEmailService(emailConfig models.EmailConfig, env config.Env) interfaces.EmailService {
	return &emailService{
		emailConfig: emailConfig,
		env:         env,
	}
}

func (es *emailService) IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func (es *emailService) SendEmail(to string, subject string, body string) *models.ErrorResponse {
	m := gomail.NewMessage()
	m.SetHeader("From", es.emailConfig.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	port, err := strconv.Atoi(es.emailConfig.Port)
	if err != nil {
		log.Println("Invalid SMTP port configuration", err.Error())
		return models.InternalServerError("Invalid SMTP port configuration")
	}

	log.Println("Sending email...", port)
	d := gomail.NewDialer(es.emailConfig.SMTPServer, port, es.emailConfig.Username, es.emailConfig.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Invalid SMTP port configuration", err.Error())

		return models.InternalServerError("Error occurred while sending email")
	}

	log.Println("Email sent successfully")

	return nil
}
