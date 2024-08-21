package dtos


type RefreshTokenDto struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}