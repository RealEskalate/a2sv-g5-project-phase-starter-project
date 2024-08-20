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
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username          string             `json:"username"`
	Email             string             `json:"email"`
	Password          string             `json:"password"`
	Role              string             `json:"role"`
	ProfilePictureUrl string             `json:"profile_picture"`
	Bio               string             `json:"bio"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

func NewUser(username, email, password, profilePictureUrl, bio string) *User {
	return &User{
		Username:          username,
		Email:             email,
		Password:          password,
		Role:              "User",
		ProfilePictureUrl: profilePictureUrl,
		Bio:               bio,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func (u *User) Validate() error {
	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(u.Email) {
		return errors.New("Invalid email")
	}

	if !IsStrongPassword(u.Password) {
		return errors.New("Password is not strong enough")
	}

	return nil
}

func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		default:
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

