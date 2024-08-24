package interfaces

import models "github.com/aait.backend.g5.main/backend/Domain/Models"

type OTPService interface {
	GenerateOTP(token string) (string, *models.ErrorResponse)
	RemoveOTP(otp_code string) *models.ErrorResponse
	GetOTP(otp_code string) (*models.URL, *models.ErrorResponse)
}
