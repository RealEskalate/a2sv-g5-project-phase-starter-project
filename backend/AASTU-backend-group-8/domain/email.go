package domain

type EmailRequest struct {
    Email string `json:"email" binding:"required"`
}