package dtos

type UpdateUser struct {
	PhoneNumber string `json:"phone_number"`
	Bio         string `json:"bio"`
}

type ResetPassword struct {
	Username    string `json:"username" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}

type GoogleResponse struct {
	RawData struct {
		Email         string `json:"email" binding:"required"`
		ID            string `json:"id" binding:"required"`
		Picture       string `json:"picture"`
		VerifiedEmail bool   `json:"verified_email" binding:"required"`
	} `json:"RawData"`
	Provider     string `json:"Provider" binding:"required"`
	Email        string `json:"Email" binding:"required"`
	UserID       string `json:"UserID" binding:"required"`
	AccessToken  string `json:"AccessToken" binding:"required"`
	RefreshToken string `json:"RefreshToken" binding:"required"`
	ExpiresAt    string `json:"ExpiresAt" binding:"required"`
	IDToken      string `json:"IDToken" binding:"required"`
}
