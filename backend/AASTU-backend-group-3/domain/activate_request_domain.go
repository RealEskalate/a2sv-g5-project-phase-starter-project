package domain

type ActivateRequest struct {
	Email string `json:"email"`
	ActivationToken string `json:"otp"`
}
