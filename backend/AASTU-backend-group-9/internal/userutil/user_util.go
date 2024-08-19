package userutil

import (
	"errors"
	"net/http"
	"regexp"
	"blog/domain"
	"github.com/xlzd/gotp"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares the password with the hashed password

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

// ValidateEmail validates the email

func ValidateEmail(email string) bool {
	// Define a regular expression for validating an email address
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Check if the email matches the regex pattern
	return emailRegex.MatchString(email)
}

// ValidatePassword validates the password

func ValidatePassword(password string) bool {
	return len(password) >= 8
}
func GenerateOTP() string {
	secretLength := 8
	return gotp.RandomSecret(secretLength)
}

// A function that checks if a the logged in user can manipulate the target user.
func CanManipulateUser(claims *domain.JwtCustomClaims, user *domain.User, manip string) *domain.Error {
	// If the user is a regular user, they can only manipulate their own account.
	if claims.Role == "user" {
		if user.ID != claims.UserID {
			var message string
			if manip == "add" {
				message = "A User cannot add a new user"
			} else {
				message = "A User cannot " + manip + " another user"
			}

			return &domain.Error{
				Err:        errors.New("unauthorized"),
				StatusCode: http.StatusForbidden,
				Message:    message,
			}
		}

		return nil
	}

	// If the user is an admin, they can manipulate all users except root user and other admin users.
	if claims.Role == "admin" {
		if user.Role == "root" {
			return &domain.Error{
				Err:        errors.New("forbidden"),
				StatusCode: http.StatusForbidden,
				Message:    "Cannot " + manip + " root user",
			}
		}

		if user.Role == "admin" && claims.UserID != user.ID {
			return &domain.Error{
				Err:        errors.New("unauthorized"),
				StatusCode: http.StatusForbidden,
				Message:    "Admin cannot " + manip + " another admin user",
			}
		}
	}

	// If the user is a root user, they can manipulate all users.
	return nil
}
