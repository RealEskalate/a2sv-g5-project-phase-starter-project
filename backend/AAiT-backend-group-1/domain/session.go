package domain

import (
	"time"
	"github.com/google/uuid"
)

type Session struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	AccessToken string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}