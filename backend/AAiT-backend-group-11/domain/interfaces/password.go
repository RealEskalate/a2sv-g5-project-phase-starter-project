package interfaces



type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}


type PasswordResetService interface {
    RequestPasswordReset(email string) error
    ResetPassword(token, newPassword string) error
}


