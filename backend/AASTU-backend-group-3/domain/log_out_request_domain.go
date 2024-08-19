package domain

type LogoutRequest struct {
	UserID string `json:"userID"`
	Token  string `json:"token"`
}