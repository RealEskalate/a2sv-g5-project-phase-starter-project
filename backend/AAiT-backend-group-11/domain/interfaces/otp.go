package interfaces


import "backend-starter-project/domain/entities"

type OTPService interface {
	GenerateOTP(email string) (*entities.OTP, error)
	VerifyOTP(email, otp string) error
}


