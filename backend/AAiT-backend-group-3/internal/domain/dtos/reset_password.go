package dtos

type ForgotPassword struct {
	Email string `json:"email" binding:"required, email"`
}

type ResetPassword struct {
	OTP      string `json:"otp" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}