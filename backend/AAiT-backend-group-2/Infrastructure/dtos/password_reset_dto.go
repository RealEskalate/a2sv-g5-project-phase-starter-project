package dtos


type PasswordResetDto struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=8,max=32,eqfield=Password"`
	Token string `json:"token" binding:"required"`
}