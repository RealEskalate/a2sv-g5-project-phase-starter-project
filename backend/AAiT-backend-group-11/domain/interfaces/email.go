package interfaces


// EmailService is an interface for sending emails
type EmailService interface {
	SendEmail(to, subject, body string) error
	GenerateEmailTemplate(header, content string) string
}