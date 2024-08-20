package usecases

import (
	"errors"
	"fmt"
	"meleket/domain"
	"meleket/infrastructure"
	"meleket/utils"
	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type OTPUsecase struct {
	otpRepository domain.OTPRepositoryInterface
	// infrastruct   infrastructure.EmailService
}

func NewOTPUsecase(or domain.OTPRepositoryInterface) *OTPUsecase {
	return &OTPUsecase{
		otpRepository: or,
	}
}

func (ou *OTPUsecase) GenerateAndSendOTP(user *domain.User) error {
	if !utils.ValidateEmail(user.Email) {
		return errors.New("invalid email")
	}
	if !utils.ValidatePassword(user.Password) {
		return errors.New("password must be at least 8 characters long")
	}

	// Generate OTP
	otp := utils.GenerateOTP(6)
	existingOtp := domain.OTP{
		Otp:       otp,
		Email:     user.Email,
		Username:  user.Name,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),

		Password: user.Password,
		Role:     user.Role,
	}

	// Store OTP in the database
	err := ou.otpRepository.StoreOTP(&existingOtp)
	if err != nil {
		return err
	}

	// Send OTP via email
	err = infrastructure.SendOTPEmail(user.Email, otp)
	if err != nil {
		return err
	}

	return nil
}

// verification endpoint
func (ou *OTPUsecase) VerifyOTP(email, otp string) (*domain.OTP, error) {
	existingOtp, err := ou.otpRepository.GetOTPByEmail(email)
	fmt.Println("verify", existingOtp.Otp)
	fmt.Println(otp)
	if err != nil {
		return nil, err
	}

	if time.Now().After(existingOtp.ExpiresAt) {
		return nil, errors.New("otp expired")
	}
	fmt.Println("verify", existingOtp.Otp, otp)
	if existingOtp.Otp != otp {
		return nil, errors.New("invalid OTP")
	}

	// Clean up the used OTP
	if err = ou.otpRepository.DeleteOTPByEmail(email); err != nil {
		return nil, errors.New("couldn't delete OTP")
	}

	return existingOtp, nil
}
