package interfaces

type PasswordService interface {
	EncryptPassword(password string) (string, error)
	ValidatePassword(password string, hashedPassword string) bool
}