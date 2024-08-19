package usecase

import (
    "blog/domain"
    "blog/internal/userutil"
    "context"
    "errors"
    "net/smtp"
    "time"
)

type forgotPasswordUsecase struct {
    userRepository domain.UserRepository
    otpRepository  domain.OTPRepository
    contextTimeout time.Duration
}

func NewForgotPasswordUsecase(userRepo domain.UserRepository, otpRepo domain.OTPRepository, timeout time.Duration) domain.ForgotPasswordUsecase {
    return &forgotPasswordUsecase{
        userRepository: userRepo,
        otpRepository:  otpRepo,
        contextTimeout: timeout,
    }
}

func (fpu *forgotPasswordUsecase) SendResetOTP(c context.Context, email string, smtpUsername, smtpPassword string) error {
    ctx, cancel := context.WithTimeout(c, fpu.contextTimeout)
    defer cancel()

    // Validate email
    if !userutil.ValidateEmail(email) {
        return errors.New("invalid email")
    }

    // Check if the email exists
    user, err := fpu.userRepository.GetUserByEmail(ctx, email)
    if err != nil {
        return err
    }
    if user == nil {
        return errors.New("email not found")
    }

    // Generate OTP
    otp := domain.OTP{
        Value:     userutil.GenerateOTP(),
        Email:     email,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(time.Minute * 15),
    }

    // Save OTP to repository
    err = fpu.otpRepository.SaveOTP(ctx, &otp)
    if err != nil {
        return err
    }

    // Send OTP via email
    return fpu.SendEmail(email, otp.Value, smtpUsername, smtpPassword)
}

func (fpu *forgotPasswordUsecase) ResetPassword(c context.Context, email, otpValue, newPassword string) error {
    ctx, cancel := context.WithTimeout(c, fpu.contextTimeout)
    defer cancel()

    // Validate password
    if !userutil.ValidatePassword(newPassword) {
        return errors.New("password must be at least 8 characters long")
    }

    // Retrieve OTP from repository
    otp, err := fpu.otpRepository.GetOTPByEmail(ctx, email)
    if err != nil {
        return err
    }
    if otp == nil || otp.Value != otpValue || time.Now().After(otp.ExpiresAt) {
        return errors.New("invalid or expired OTP")
    }

    // Update user's password
    user, err := fpu.userRepository.GetUserByEmail(ctx, email)
    if err != nil {
        return err
    }
    hashedPassword, err := userutil.HashPassword(newPassword)
    if err != nil {
        return err
    }
    user.Password = hashedPassword
    err = fpu.userRepository.UpdatePassword(ctx, user)
    if err != nil {
        return err
    }

    // Delete OTP after successful password reset
    return fpu.otpRepository.DeleteOTP(ctx, email)
}

func (fpu *forgotPasswordUsecase) SendEmail(email, otpValue, smtpUsername, smtpPassword string) error {
    from := smtpUsername
    password := smtpPassword
    to := []string{email}
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    message := []byte("Your OTP is " + otpValue)
    auth := smtp.PlainAuth("", from, password, smtpHost)
    return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
}
