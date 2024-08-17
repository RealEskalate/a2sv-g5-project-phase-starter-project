// . User
// Attributes:
// id (UUID): Unique identifier for the user.
// username (String): User's chosen username.
// email (String): User's email address.
// password (String): Hashed password for authentication.
// role (Enum): User role (e.g., Admin, User).
// profile_picture (String): URL or path to the profile picture.
// bio (String): User bio or description.
// created_at (Timestamp): Date and time when the user registered.
// updated_at (Timestamp): Date and time when the user's profile was last updated.
// is_active (Boolean): Indicates if the user account is active.

package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Role           string    `json:"role"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
