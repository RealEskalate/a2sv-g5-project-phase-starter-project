package email

import "blog-api/domain"

type emailService struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	FromAddress  string
}

func NewEmailService(server, port, user, password, fromAddress string) domain.EmailService {
	return &emailService{
		SMTPServer:   server,
		SMTPPort:     port,
		SMTPUser:     user,
		SMTPPassword: password,
		FromAddress:  fromAddress,
	}
}
