package entities

type EmailTemplate struct {
	Subject string
	Body    string
}

type SMTPConfig struct {
    Server   string
    Username string
    Password string
}