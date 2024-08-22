package interfaces

import models "github.com/aait.backend.g5.main/backend/Domain/Models"

type EmailService interface {
	IsValidEmail(email string) bool
	SendEmail(email string, subject string, body string) *models.ErrorResponse
}
