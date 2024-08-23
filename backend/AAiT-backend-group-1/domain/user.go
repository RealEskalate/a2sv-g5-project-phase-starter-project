package domain

import (
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username          string             `json:"username"`
	Email             string             `json:"email"`
	Password          string             `json:"password"`
	Role              string             `json:"role"`
	ProfilePictureUrl Photo              `json:"profile_picture"`
	Bio               string             `json:"bio"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

type Photo struct {
	Filename   string    `bson:"filename"`
	FilePath   string    `bson:"file_path"`
	Public_id  string    `bson:"public_id"`
	UploadedAt time.Time `bson:"uploaded_at"`
}

func NewUser(username, email, password, bio string, profilePictureUrl Photo) *User {
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
		return errors.New("invalid email format")
	}

	// if IsValidDomain(u.Email) {
	// 	return errors.New("invalid email Domain")
	// }

	// if HasMXRecord(u.Email) {
	// 	return errors.New("inactive email record")
	// }

	// if IsValidDomain(u.Email) {
	// 	return errors.New("invalid email domain")
	// }
	//
	// if HasMXRecord(u.Email) {
	// 	return errors.New("inactive email")
	// }
	//

	if !IsStrongPassword(u.Password) {
		return errors.New("password is not strong enough")
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
