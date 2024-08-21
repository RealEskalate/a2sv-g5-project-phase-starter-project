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
}

func NewForgotPasswordUsecase(
    userRepo domain.UserRepository,
    otpRepo domain.OTPRepository,
    otpService domain.OtpInfrastructure,
    passwordInfrastructure domain.PasswordInfrastructure,
) domain.ForgotPasswordUsecase {
    return &forgotPasswordUsecase{
        UserRepository:          userRepo,
        OTPRepository:           otpRepo,
        OTPService:              otpService,
        PasswordInfrastructure: passwordInfrastructure,
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

    otpCode, err := u.OTPService.CreateOTP(&domain.UserOTPRequest{
        UserID: user.ID.Hex(),
        Email:  user.Email,
    })
    if err != nil {
        return domain.ForgotPasswordResponse{}, err
    }

    hashedOTPCode, err := u.PasswordInfrastructure.HashPassword(otpCode)
    if err != nil {
        return domain.ForgotPasswordResponse{}, err
    }

    otpRecord := domain.UserOTPVerification{
        ID:         user.ID,
        User_ID:    user.ID.Hex(),
        Email:      user.Email,
        OTP:        hashedOTPCode,
        Created_At: time.Now(),
        Expires_At: time.Now().Add(15 * time.Minute),
    }


    err = u.OTPRepository.CreateOTP(ctx, &otpRecord)
    if err != nil {
        return domain.ForgotPasswordResponse{}, err
    }

    return domain.ForgotPasswordResponse{Message: "Password reset OTP generated successfully"}, nil
}


func (u *forgotPasswordUsecase) ResetPassword(ctx context.Context, request domain.ResetPasswordRequest) (domain.ResetPasswordResponse, error) {

    otpRecord, err := u.OTPRepository.GetOTPByEmail(ctx, request.Email)
    if err != nil || time.Now().After(otpRecord.Expires_At) {
        return domain.ResetPasswordResponse{}, errors.New("invalid or expired OTP")
    }

    err = u.PasswordInfrastructure.ComparePasswords(request.OTP, otpRecord.OTP)
    if err != nil {
        return domain.ResetPasswordResponse{}, errors.New("incorrect OTP")
    }

    hashedPassword, err := u.PasswordInfrastructure.HashPassword(request.NewPassword)
    if err != nil {
        return domain.ResetPasswordResponse{}, err
    }

    user, err := u.UserRepository.GetByEmail(ctx, request.Email)
    if err != nil {
        return domain.ResetPasswordResponse{}, err
    }

    updates := map[string]interface{}{
        "password":            hashedPassword,
        "password_reset_token": "",
        "token_expiry":         time.Time{}, 
    }

    err = u.UserRepository.UpdatePasswordTokens(ctx, user.ID.Hex(), updates)
    if err != nil {
        return domain.ResetPasswordResponse{}, err
    }

    err = u.OTPRepository.DeleteOTPByEmail(ctx, request.Email)
    if err != nil {
        return domain.ResetPasswordResponse{}, err
    }

    return domain.ResetPasswordResponse{Message: "Password has been reset successfully"}, nil
}
