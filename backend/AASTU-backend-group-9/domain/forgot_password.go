package domain


type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

// ResetPasswordRequest represents a request to reset the password using an OTP
type ResetPasswordRequest struct {
	OTPValue    string `json:"otp"`
	NewPassword string `json:"new_password"`
}

