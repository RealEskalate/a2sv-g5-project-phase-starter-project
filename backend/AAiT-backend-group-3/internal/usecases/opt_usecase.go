package usecases

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/infrastructures/services"
	"AAIT-backend-group-3/internal/repositories/interfaces"
	"context"
	"errors"
	"fmt"
	"time"

)

type IOtpUsecase interface {
	GenerateAndSendOtp(ctx context.Context, email string) error
	ValidateOtp(ctx context.Context, otp string) (*models.OtpEntry, error)
	ResetPassword(ctx context.Context, userID string, newPassword string) error
}

type OtpUsecase struct {
	otpRepo  repository_interface.IOtpRepository
	userRepo repository_interface.UserRepositoryInterface
	emailSvc services.IEmailService
	passSvc services.IHashService
	redirectURL  string 
}

func NewOtpUseCase(otpRepo repository_interface.IOtpRepository,userRepo repository_interface.UserRepositoryInterface,emailSvc services.IEmailService,passSvc services.IHashService,redirectURL string,) IOtpUsecase {
	return &OtpUsecase{
		otpRepo:  otpRepo,
		userRepo: userRepo,
		emailSvc: emailSvc,
		redirectURL:  redirectURL,
		passSvc: passSvc,
	}
}

func (u *OtpUsecase) GenerateAndSendOtp(ctx context.Context, email string) error {
	var otpEntry models.OtpEntry
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	otp := services.GenerateOTP()
	expiresAt := time.Now().Add(10 * time.Minute)

	otpEntry.ExpiresAt = expiresAt
	otpEntry.UserID = user.ID.Hex()
	otpEntry.OTP = otp;
	
	err = u.otpRepo.SaveOtp(ctx, otpEntry)
	if err != nil {
		return err
	}

	resetLink := fmt.Sprintf("%s/reset-password?otp=%s", u.redirectURL, otp)

	subject := "Password Reset Request"
	body := fmt.Sprintf(`
		<h1>Password Reset</h1>
		<p>To reset your password, use the following OTP:</p>
		<p><strong>%s</strong></p>
		<p>Or click on the link below:</p>
		<a href="%s">Reset Password</a>
	`, otp, resetLink)

	err = u.emailSvc.SendResetEmail(email, subject, body)
	if err != nil {
		return err
	}

	return nil
}

func (u *OtpUsecase) ValidateOtp(ctx context.Context, otp string) (*models.OtpEntry, error) {
	otpEntry, err := u.otpRepo.FindByOtp(ctx, otp)
	if err != nil {
		return nil, errors.New("invalid OTP")
	}

	if time.Now().After(otpEntry.ExpiresAt) {
		return nil, errors.New("OTP has expired")
	}

	return otpEntry, nil
}

func (u *OtpUsecase) ResetPassword(ctx context.Context, userID string, newPassword string) error {
	hashedPassword, err := u.passSvc.HashPassword(newPassword)
	if err != nil {
		return err
	}

	err = u.userRepo.UpdatePassword(userID, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}