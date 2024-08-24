// Package models defines the domain models for the blogging application,
// including structures for managing blog posts, users, comments, and reactions.
package models

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	"github.com/nbutton23/zxcvbn-go"
)

const (
	minPasswordStrengthScore = 3
	usernamePattern          = `^[a-zA-Z0-9_]+$`
	emailPattern             = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	minUsernameLength        = 3
	maxUsernameLength        = 20
	minFirstNameLength       = 3
	maxFirstNameLength       = 250
)

var (
	usernameRegex = regexp.MustCompile(usernamePattern)
	emailRegex    = regexp.MustCompile(emailPattern)
)

// ResetCode represents a code used for password resets.
type ResetCode struct {
	CodeHash string    // The reset code
	Expr     time.Time // Expiration time of the reset code
}

// User represents a system user with private fields.
type User struct {
	id           uuid.UUID
	firstName    string
	lastName     string
	username     string
	email        string
	passwordHash string
	isAdmin      bool
	resetCode    *ResetCode
	createdAt    time.Time
	updatedAt    time.Time
	isActive     bool
}

// UserConfig holds parameters for creating a new User.
type UserConfig struct {
	FirstName      string        // User's first name
	LastName       string        // User's last name
	Username       string        // User's username
	Email          string        // User's email address
	PlainPassword  string        // Plain text password
	IsAdmin        bool          // True if the user should be an admin
	PasswordHasher ihash.Service // Service for hashing passwords
}

// MapUserConfig holds parameters for mapping an existing User.
type MapUserConfig struct {
	Id             uuid.UUID
	FirstName      string
	LastName       string
	Username       string
	Email          string
	HashedPassword string
	IsAdmin        bool
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ResetCode      *ResetCode
}

// NewUser creates a new User with the provided configuration.
func NewUser(config UserConfig) (*User, error) {
	if err := validateUsername(config.Username); err != nil {
		return nil, err
	}
	if err := validatePassword(config.PlainPassword); err != nil {
		return nil, err
	}
	if err := validateEmail(config.Email); err != nil {
		return nil, err
	}
	if err := validateFirstName(config.FirstName); err != nil {
		return nil, err
	}

	passwordHash, err := config.PasswordHasher.Hash(config.PlainPassword)
	if err != nil {
		return nil, err
	}

	return &User{
		id:           uuid.New(),
		firstName:    config.FirstName,
		lastName:     config.LastName,
		username:     config.Username,
		email:        config.Email,
		passwordHash: passwordHash,
		isAdmin:      config.IsAdmin,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}, nil
}

// NewUser creates a new User with the provided configuration.
func NewFederatedUser(config UserConfig) *User {

	return &User{
		id:        uuid.New(),
		firstName: config.FirstName,
		lastName:  config.LastName,
		email:     config.Email,
		isAdmin:   config.IsAdmin,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

// MapUser maps an existing User from the database configuration.
func MapUser(config MapUserConfig) *User {
	return &User{
		id:           config.Id,
		firstName:    config.FirstName,
		lastName:     config.LastName,
		username:     config.Username,
		email:        config.Email,
		passwordHash: config.HashedPassword,
		isAdmin:      config.IsAdmin,
		isActive:     config.IsActive,
		createdAt:    config.CreatedAt,
		updatedAt:    config.UpdatedAt,
		resetCode:    config.ResetCode,
	}
}

// validateUsername checks if the username meets length and format requirements.
func validateUsername(username string) error {
	if len(username) < minUsernameLength {
		return er.UsernameTooShort
	} else if len(username) > maxUsernameLength {
		return er.UsernameTooLong
	}
	if !usernameRegex.MatchString(username) {
		return er.UsernameInvalidFormat
	}
	return nil
}

// validatePassword checks if the password meets strength requirements.
func validatePassword(password string) error {
	if zxcvbn.PasswordStrength(password, nil).Score < minPasswordStrengthScore {
		return er.WeakPassword
	}
	return nil
}

// validateEmail checks if the email format is valid.
func validateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return er.EmailInvalidFormat
	}
	return nil
}

// validateFirstName checks if the first name meets length requirements.
func validateFirstName(firstName string) error {
	if len(firstName) < minFirstNameLength {
		return er.FirstNameTooShort
	} else if len(firstName) > maxFirstNameLength {
		return er.FirstNameTooLong
	}
	return nil
}

// ID returns the unique identifier of the User.
func (u *User) ID() uuid.UUID { return u.id }

// Username returns the User's username.
func (u *User) Username() string { return u.username }

// FirstName returns the User's first name.
func (u *User) FirstName() string { return u.firstName }

// LastName returns the User's last name.
func (u *User) LastName() string { return u.lastName }

// Email returns the User's email address.
func (u *User) Email() string { return u.email }

// PasswordHash returns the hashed password of the User.
func (u *User) PasswordHash() string { return u.passwordHash }

// CreatedAt returns the date and time when the User account was created.
func (u *User) CreatedAt() time.Time { return u.createdAt }

// UpdatedAt returns the date and time when the User account was last updated.
func (u *User) UpdatedAt() time.Time { return u.updatedAt }

// IsAdmin returns true if the User is an admin.
func (u *User) IsAdmin() bool { return u.isAdmin }

// IsActive returns true if the User is active.
func (u *User) IsActive() bool { return u.isActive }

// ResetCode returns the User's password reset code.
func (u *User) ResetCode() *ResetCode { return u.resetCode }

// UpdateUsername updates the User's username after validation.
func (u *User) UpdateUsername(username string) error {
	if err := validateUsername(username); err != nil {
		return err
	}
	u.username = username
	return nil
}

// MakeActive sets the User's status to active.
func (u *User) MakeActive()   { u.isActive = true }
func (u *User) MakeInactive() { u.isActive = false }

// UpdateFirstName updates the User's first name after validation.
func (u *User) UpdateFirstName(firstName string) error {
	if err := validateFirstName(firstName); err != nil {
		return err
	}
	u.firstName = firstName
	return nil
}

// UpdateLastName updates the User's last name.
func (u *User) UpdateLastName(lastName string) error {
	u.lastName = lastName
	return nil
}

func (u *User) UpdateEmail(email string) error {
	u.email = email
	u.MakeInactive()
	return nil
}

// UpdateResetCode sets a new password reset code for the User.
func (u *User) UpdateResetCode(code string, expr time.Time, hashService ihash.Service) error {
	hashedCode, err := hashService.Hash(code)
	if err != nil {
		return err
	}
	u.resetCode = &ResetCode{
		CodeHash: hashedCode,
		Expr:     expr,
	}
	return nil
}

func (u *User) RemoveResetCode() {
	u.resetCode = nil
}

// UpdatePassword updates the User's password after validation.
func (u *User) UpdatePassword(newPassword string, hasher ihash.Service) error {
	if err := validatePassword(newPassword); err != nil {
		return err
	}
	hashedPassword, err := hasher.Hash(newPassword)
	if err != nil {
		return err
	}
	u.passwordHash = hashedPassword
	return nil
}

// UpdateAdminStatus updates the User's admin status.
func (u *User) UpdateAdminStatus(isAdmin bool) { u.isAdmin = isAdmin }


