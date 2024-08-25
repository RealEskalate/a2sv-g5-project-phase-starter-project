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
    Email string `json:"email" binding:"required, email"`
    Code  string `json:"code" binding:"required, min=5"`
}
