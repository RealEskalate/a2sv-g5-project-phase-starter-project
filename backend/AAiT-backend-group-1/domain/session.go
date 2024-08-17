// id (UUID): Unique identifier for the session.
// user_id (UUID): Foreign key linking to the User model.
// access_token (String): JWT access token.
// refresh_token (String): JWT refresh token.
// expires_at (Timestamp): Expiration time for the access token.
// created_at (Timestamp): Date and time when the session was created.
// updated_at (Timestamp): Date and time when the session was last updated.

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