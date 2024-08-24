package email

import (
	"log"

	iemail "github.com/group13/blog/usecase/common/i_email"
	"gopkg.in/gomail.v2"
)

type MailTrapService struct {
	port     int
	host     string
	username string
	password string
}

type Config struct {
	Port     int
	Host     string
	Username string
	Password string
}

func NewMailTrapService(config Config) *MailTrapService {
	return &MailTrapService{
		port:     config.Port,
		host:     config.Host,
		username: config.Username,
		password: config.Password,
	}
}

func (es *MailTrapService) Send(mail *iemail.Mail) error {
	log.Printf("sending mail to %v -- MailTrapService", mail.To)
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@EBS.com")
	for _, reciever := range mail.To {
		m.SetHeader("To", reciever)
	}
	m.SetHeader("Subject", "Activate your account")
	m.SetBody("text/html", mail.Body)

	d := gomail.NewDialer(es.host, es.port, es.username, es.password)

	return d.DialAndSend(m)
}
