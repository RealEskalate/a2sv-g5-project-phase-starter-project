package domain

type PasswordService interface {
	ValidatePassword(password string) error
	HashPassword(password string) (string, error)
	ValidateEmail(email string) error
	ComparePassword(password, hashedPassword string) bool
}
