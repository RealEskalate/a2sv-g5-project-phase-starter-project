package domain


type PasswordResetModel struct {
	UserName string `json:"user_name" binding:"required"`
	ResetLink string `json:"reset_link" binding:"required"`
}