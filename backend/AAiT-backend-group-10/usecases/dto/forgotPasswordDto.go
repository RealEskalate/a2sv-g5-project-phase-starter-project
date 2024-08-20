package dto

type ForgotPasswordRequestDTO struct {
	Email string `json:"email" binding:"required,email"`
}

