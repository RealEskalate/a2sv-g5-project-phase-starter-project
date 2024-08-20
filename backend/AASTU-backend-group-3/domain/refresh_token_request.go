package domain

type RefreshTokenRequest struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}