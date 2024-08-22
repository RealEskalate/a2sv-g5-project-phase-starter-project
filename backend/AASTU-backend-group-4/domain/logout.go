package domain

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}
