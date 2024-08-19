package dtos

type UpdateUser struct {
	PhoneNumber string `json:"phone_number"`
	Bio         string `json:"bio"`
}

type ResetPassword struct {
	Username    string `json:"username" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}
