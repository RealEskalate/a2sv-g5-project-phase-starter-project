package usecases

import (
    "context"
    "errors"
    "time"
    "aait-backend-group4/Domain"
)

type forgotPasswordUsecase struct {
    UserRepository          domain.UserRepository
    OTPRepository           domain.OTPRepository
    OTPService              domain.OtpInfrastructure
    PasswordInfrastructure domain.PasswordInfrastructure
    OTPUsecase              domain.OTPUsecase
}

func NewForgotPasswordUsecase(
    userRepo domain.UserRepository,
    otpRepo domain.OTPRepository,
    otpService domain.OtpInfrastructure,
    passwordInfrastructure domain.PasswordInfrastructure,
    otpUsecase domain.OTPUsecase,
) domain.ForgotPasswordUsecase {
    return &forgotPasswordUsecase{
        UserRepository:          userRepo,
        OTPRepository:           otpRepo,
        OTPService:              otpService,
        PasswordInfrastructure: passwordInfrastructure,
        OTPUsecase:             otpUsecase,

    }
}

// ForgotPassword handles the process of initiating a password reset.
// It generates an OTP and updates the user's record with the OTP and its expiry.
func (u *forgotPasswordUsecase) ForgotPassword(ctx context.Context, request domain.ForgotPasswordRequest) (domain.ForgotPasswordResponse, error) {
    // Fetch the user by email
    user, err := u.UserRepository.GetByEmail(ctx, request.Email)
    if err != nil {
        return domain.ForgotPasswordResponse{}, err
    }

    // Use the OTPService to generate an OTP and send it via email
    otpRequest := &domain.UserOTPRequest{
        UserID: user.ID.Hex(),
        Email:  user.Email,
    }
    otpResponse, err := u.OTPUsecase.GenerateOTP(ctx, otpRequest)
    if err != nil {
        return domain.ForgotPasswordResponse{}, err
    }

    // The OTPService.GenerateOTP already handles OTP creation and sending

    return domain.ForgotPasswordResponse{Message: otpResponse.Message}, nil
}

func (u *forgotPasswordUsecase) ResetPassword(ctx context.Context, request domain.ResetPasswordRequest) (domain.ResetPasswordResponse, error) {
    // Fetch the OTP record
    otpRecord, err := u.OTPRepository.GetOTPByEmail(ctx, request.Email)
    if err != nil || time.Now().After(otpRecord.Expires_At) {
        return domain.ResetPasswordResponse{}, errors.New("invalid or expired OTP")
    }

    // Compare the provided OTP with the stored OTP
    err = u.PasswordInfrastructure.ComparePasswords(request.OTP, otpRecord.OTP)
    if err != nil {
        return domain.ResetPasswordResponse{}, errors.New("incorrect OTP")
    }

    // Hash the new password
    hashedPassword, err := u.PasswordInfrastructure.HashPassword(request.NewPassword)
    if err != nil {
        return domain.ResetPasswordResponse{}, err
    }

    // Fetch the user to update their password
    user, err := u.UserRepository.GetByEmail(ctx, request.Email)
    if err != nil {
        return domain.ResetPasswordResponse{}, err
    }

    // Update the user's password and clear OTP related fields
    updates := map[string]interface{}{
        "password":             hashedPassword,
        "password_reset_token": "",
        "token_expiry":         time.Time{},
    }

    err = u.UserRepository.UpdatePasswordTokens(ctx, user.ID.Hex(), updates)
    if err != nil {
        return domain.ResetPasswordResponse{}, err
    }

    // Delete the OTP record after successful password reset
    err = u.OTPRepository.DeleteOTPByEmail(ctx, request.Email)
    if err != nil {
        return domain.ResetPasswordResponse{}, err
    }

    return domain.ResetPasswordResponse{Message: "Password has been reset successfully"}, nil
}
