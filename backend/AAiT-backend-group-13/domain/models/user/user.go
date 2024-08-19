package usermodel

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

type ResetCode struct {
	Code int64
	Expr time.Time
}

// User represents the aggregate user with private fields.
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

// Config holds parameters for creating a new User.
type Config struct {
	FirstName      string
	LastName       string
	Username       string
	Email          string
	PlainPassword  string
	IsAdmin        bool
	PasswordHasher ihash.Service
}

// Config holds parameters for creating a new User.
type MapConfig struct {
	Id             uuid.UUID
	FirstName      string
	LastName       string
	Username       string
	Email          string
	HashedPassword string
	IsAdmin        bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ResetCode      *ResetCode
}

// New creates a new User with the provided configuration.
func New(config Config) (*User, error) {
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

	//returns user with specified fields
	return &User{
		id:           uuid.New(),
		username:     config.Username,
		passwordHash: passwordHash,
		isAdmin:      config.IsAdmin,
		email:        config.Email,
		firstName:    config.FirstName,
		lastName:     config.LastName,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}, nil
}

// Map maps a User from database and returns user pointer.
func Map(config MapConfig) (*User, error) {

	//returns user with specified fields
	return &User{
		id:           config.Id,
		username:     config.Username,
		passwordHash: config.HashedPassword,
		isAdmin:      config.IsAdmin,
		email:        config.Email,
		firstName:    config.FirstName,
		lastName:     config.LastName,
		resetCode:    config.ResetCode,
		createdAt:    config.CreatedAt,
		updatedAt:    config.UpdatedAt,
	}, nil
}

// validateUsername validates the username.
func validateUsername(username string) error {
	if len(username) < minUsernameLength {
		return er.UsernameTooShort
	}
	if len(username) > maxUsernameLength {
		return er.UsernameTooLong
	}
	if !usernameRegex.MatchString(username) {
		return er.UsernameInvalidFormat
	}
	return nil
}

// validatePassword checks the strength of the password.
func validatePassword(password string) error {
	result := zxcvbn.PasswordStrength(password, nil)
	if result.Score < minPasswordStrengthScore {
		return er.WeakPassword
	}
	return nil
}

// validateEmail checks the email validity.
func validateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return er.EmailInvalidFormat
	}
	return nil
}

// validateFirstName validates FirstName.
func validateFirstName(firstName string) error {
	if len(firstName) < minFirstNameLength {
		return er.UsernameTooShort
	}
	if len(firstName) > maxFirstNameLength {
		return er.UsernameTooLong
	}
	return nil
}

// ID returns the user's ID.
func (u *User) ID() uuid.UUID {
	return u.id
}

// Username returns the user's username.
func (u *User) Username() string {
	return u.username
}

// FirstName returns the user's firstname.
func (u *User) FirstName() string {
	return u.firstName
}

// LastName returns user's lastname.
func (u *User) LastName() string {
	return u.lastName
}

// Email returns user's email.
func (u *User) Email() string {
	return u.email
}

// PasswordHash returns the user's password hash.
func (u *User) PasswordHash() string {
	return u.passwordHash
}

// CreatedAt returns the user's account Created date.
func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAT returns the user's account last Updated date.
func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

// IsAdmin returns whether the user is an admin.
func (u *User) IsAdmin() bool {
	return u.isAdmin
}

// IsActive returns whether the user is an active.
func (u *User) IsActive() bool {
	return u.isActive

func (u *User) ResetCode() *ResetCode {
	return u.resetCode
}

// UpdateUsername updates the user's username after validation.
func (u *User) UpdateUsername(newUsername string) error {
	if err := validateUsername(newUsername); err != nil {
		return err
	}
	u.username = newUsername
	return nil
}

func (u *User) MakeActive() {
	u.isActive = true
}

// UpdateFirstName updates the user's firstname after validation.
func (u *User) UpdateFirstName(newFirstName string) error {
	if err := validateFirstName(newFirstName); err != nil {
		return err
	}
	u.firstName = newFirstName
	return nil
}

// UpdateLastNname updates the user's lastname.
func (u *User) UpdateLastName(newLastName string) error {
	u.firstName = newLastName
	return nil
}

func (u *User) UpdateResetCode(code *ResetCode) error {
	u.resetCode = code
	return nil
}

// UpdatePassword updates the user's password after validation.
func (u *User) UpdatePassword(newPassword string, passwordHasher ihash.Service) error {
	if err := validatePassword(newPassword); err != nil {
		return err
	}

	hashedPassword, err := passwordHasher.Hash(newPassword)
	if err != nil {
		return err
	}

	u.passwordHash = hashedPassword
	return nil
}

// UpdateAdminStatus updates the user's admin status.
func (u *User) UpdateAdminStatus(isAdmin bool) {
	u.isAdmin = isAdmin
}
