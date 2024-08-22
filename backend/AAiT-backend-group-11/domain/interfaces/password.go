package interfaces

import "backend-starter-project/domain/entities"



type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}


type PasswordResetService interface {
    RequestPasswordReset(email string) error
    ResetPassword(token, newPassword string) error
	GeneratePasswordResetToken(user *entities.User) (string, error)
}

type PasswordTokenRepository interface {
	CreatePasswordResetToken(token *entities.PasswordResetToken) (*entities.PasswordResetToken, error)
	FindPasswordReset(tok string) (*entities.PasswordResetToken, error)
	DeletePasswordResetTokenByUserId(userId string) error

}
