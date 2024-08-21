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

type EmailVerificationRequest struct {
    Email string `json:"email" binding:"required"`
    Code  string `json:"code" binding:"required"`
}
