package interfaces


import "backend-starter-project/domain/entities"


type OTPService interface {
	SaveOtp(otp *entities.OTP) error
	InvalidateOtp(otp *entities.OTP) error
	GetOtpByEmail(email string) (entities.OTP, error)
}

type OTPRepository interface {
	SaveOtp(otp *entities.OTP) error
	InvalidateOtp(otp *entities.OTP) error
	GetOtpByEmail(email string) (entities.OTP, error)
	GetByID(id string) (entities.OTP, error)
}
