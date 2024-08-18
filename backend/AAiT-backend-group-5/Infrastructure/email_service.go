package infrastructure

import (
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/asaskevich/govalidator"
)

type emailService struct {
}

func NewEmailService() interfaces.EmailService {
	return &emailService{}
}

func (es *emailService) IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func (es *emailService) SendEmail(email string) *models.ErrorResponse {
	return nil
}
