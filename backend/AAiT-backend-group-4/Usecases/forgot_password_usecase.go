package usecases

import (
	domain "aait-backend-group4/Domain"
	"context"
	"fmt"
	"time"
)

type ForgotPasswordUsecase struct {
	userRepository  domain.UserRepository
	passwordService domain.PasswordInfrastructure
	otpService      domain.OtpInfrastructure
	contextTimeout  time.Duration
}

func NewPasswordUsecase(userRepository domain.UserRepository, otpService domain.OtpInfrastructure,
	timeout time.Duration, passwordService domain.PasswordInfrastructure) domain.ForgotPasswordUsecase {
	return &ForgotPasswordUsecase{
		userRepository:  userRepository,
		otpService:      otpService,
		passwordService: passwordService,
		contextTimeout:  timeout,
	}
}

func (fu *ForgotPasswordUsecase) VerifyChangePassword(c context.Context, email string, request domain.ForgotPasswordRequest) (resp domain.ForgotPasswordResponse, err error) {
	userFound, err := fu.GetByEmail(c, email)
	if err != nil {
		return resp, err
	}

	if request.New_Password != request.Confirmation {
		return resp, fmt.Errorf("passwords do not match")
	}

	newHashedPassword, err := fu.passwordService.HashPassword(request.New_Password)
	if err != nil {
		return resp, err
	}

	updatePassword := domain.UserUpdate{
		Password: &newHashedPassword,
	}

	_, err = fu.userRepository.UpdateUser(c, userFound.ID.Hex(), updatePassword)
	if err != nil {
		return resp, err
	}

	resp.Message = "Password updated successfully"

	return resp, err
}

func (fu *ForgotPasswordUsecase) ForgotPassword(c context.Context, email string, key string) (resp domain.OTPVerificationResponse, err error) {
	err = fu.otpService.SendPasswordResetEmail(email, "Reset Password", key)
	if err != nil {
		return resp, err
	}
	resp.Status = "Not verified"
	resp.Message = "Password reset email sent successfully"

	return resp, err
}

func (fu *ForgotPasswordUsecase) GetByEmail(c context.Context, email string) (user domain.User, err error) {
	user, err = fu.userRepository.GetByEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}

	return user, err
}

func (fu *ForgotPasswordUsecase) GetByUsername(c context.Context, userName string) (user domain.User, err error) {
	user, err = fu.userRepository.GetByUsername(c, userName)
	if err != nil {
		return domain.User{}, err
	}

	return user, err
}
