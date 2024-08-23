package entities

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type PasswordReset struct {
	Email    string `json:"email"`
	NewPassword string `json:"password" binding:"required"`
	Token    string `json:"token" binding:"required"`
}