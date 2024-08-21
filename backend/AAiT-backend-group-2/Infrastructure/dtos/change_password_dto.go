package dtos


type ChangePasswordDto struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=8,max=32,eqfield=NewPassword"`
}