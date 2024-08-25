package entities

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type PasswordReset struct {
	Email    string `json:"email binding:"required email"`
	NewPassword string `json:"password" binding:"required,min=6"`
	Token    string `json:"token" binding:"required"`
}