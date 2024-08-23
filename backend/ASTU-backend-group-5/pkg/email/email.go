package email

type EmailSender interface {
	SendVerificationEmail(userEmail string, token string) error
	SendPasswordResetEmail(userEmail string, token string) error
}

type EmailPyload struct {
	To      string
	Subject string
	Body    string
}
