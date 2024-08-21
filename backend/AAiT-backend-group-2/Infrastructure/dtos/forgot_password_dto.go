package dtos


type ForgotPasswordDto struct {
	Email string `json:"email" binding:"required,email"`
}