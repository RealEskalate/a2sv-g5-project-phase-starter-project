package email

type EmailSender interface {
	SendVerificationEmail(userEmail string, token string) error
	SendPasswordResetEmail(userEmail string, token string) error
	SendPromotionToAdminEmail(userEmail string) error
	SendDemotionFromAdminEmail(userEmail string) error
}

type EmailPyload struct {
	To      string
	Subject string
	Body    string
}
