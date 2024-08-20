package usecase

import (
	"Blog_Starter/domain"
	"context"
	"time"
)

type otpUsecase struct {
	otpRepository  domain.OtpRepository
	contextTimeout time.Duration
}


func NewOtpUsecase(otpRepository domain.OtpRepository, timeout time.Duration) domain.OtpUsecase {
	return &otpUsecase{
		otpRepository:  otpRepository,
		contextTimeout: timeout,
	}
}

func (su *otpUsecase) SaveOtp(c context.Context, otp *domain.Otp) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.otpRepository.SaveOtp(ctx, otp)
}

func (su *otpUsecase) GetOtpByEmail(c context.Context, email string) (domain.Otp, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.otpRepository.GetOtpByEmail(ctx, email)
}

func (su *otpUsecase) InvalidateOtp(c context.Context, otp *domain.Otp) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.otpRepository.InvalidateOtp(ctx, otp)
}

// GetByID implements domain.OtpUsecase.
func (su *otpUsecase) GetByID(c context.Context, id string) (domain.Otp, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.otpRepository.GetByID(ctx, id)
}
