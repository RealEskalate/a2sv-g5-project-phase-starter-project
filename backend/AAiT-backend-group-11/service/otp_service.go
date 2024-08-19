package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"time"
)

type OtpService struct {
	otpRepo interfaces.OTPRepository
}

// GetOtpByEmail implements interfaces.OTPService.
func (o *OtpService) GetOtpByEmail(email string) (entities.OTP, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return o.otpRepo.GetOtpByEmail(email)
}

// InvalidateOtp implements interfaces.OTPService.
func (o *OtpService) InvalidateOtp(otp *entities.OTP) error {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return o.otpRepo.InvalidateOtp(otp)
}

// SaveOtp implements interfaces.OTPService.
func (o *OtpService) SaveOtp(otp *entities.OTP) error {
	_, cancel := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()

	return o.otpRepo.SaveOtp(otp)
}

func NewOtpService(otpRepo interfaces.OTPRepository) interfaces.OTPService {
	return &OtpService{otpRepo: otpRepo}

}
