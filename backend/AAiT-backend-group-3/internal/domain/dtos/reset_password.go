package dtos

type ForgotPassword struct {
	Email string `json:"email" binding:"required, email"`
}

type ResetPassword struct {
	Otp     string `json:"otp" binding:"required"`
	NewPassword string `json:"password" binding:"required,min=8"`
}