package Interfaces

import (
	domain "AAiT-backend-group-8/Domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type IUserRepository interface {
// 	CreateUser(user *domain.User) error
// 	GetUserByEmail(email string) (*domain.User, error)
// 	VerifyUser(user *domain.User) error
// 	GetUserByVerificationToken(token string) (*domain.User, error)
// 	GetUserCount() (int64, error)
// 	UpdatePasswordByEmail(email string, newPassword string) error
// 	StoreResetToken(email string, resetToken string) error
// 	InvalidateResetToken(email string) error
// 	GetResetTokenByEmail(email string) (string, error)
// }

// type IUserUseCase interface {
// 	Login(email, password string) (string, string, error)
// 	GetSingleUser(email string) (*domain.User, error)
// 	RegisterUser(user *domain.User) error
// 	VerifyEmail(token string) error
// 	RefreshToken(email, refresher string) (string, error)
// 	GenerateResetPasswordToken(email string) error
// 	ResetPassword(token string, newPassword string) error
// 	StoreToken(token string) error
// }

type IUserUseCase interface {
	Login(email, password string) (string, string, error)
	GetSingleUser(email string) (*domain.User, error)
	RegisterUser(user *domain.User) error
	VerifyEmail(token string) error
	DeleteUser(email string) error
	PromoteUser(email string) error
	DemoteUser(email string) error
	RefreshToken(email, refresher string) (string, error)
	GenerateResetPasswordToken(email string) error
	ResetPassword(token string, newPassword string) error
	StoreToken(token string) error
	Logout(email, refresher string) error
	// DeleteRefresher(email, refresher string) error
}

// type IPasswordService interface {
// 	HashPassword(password string) (string, error)
// 	VerifyPassword(hashedPassword, plainPassword string) error
// }
// type ITokenService interface {
// 	GenerateToken(email string, id primitive.ObjectID, role, name string, expiryDuration int64) (string, error)
// 	ValidateToken(string) (map[string]interface{}, error)
// 	GetClaimsOfToken(token string) (map[string]interface{}, error)
// }

// type ITokenRepository interface {
// 	InsertRefresher(credential domain.Credential) error
// 	GetRefresher(email string) (string, error)
// 	InvalidateResetToken(email string) error
// 	StoreResetToken(email string, resetToken string) error
// }

// type IMailService interface {
// 	SendVerificationEmail(to, token string) error
// 	SendPasswordResetEmail(email string, resetToken string) error
// }

type IUserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	VerifyUser(user *domain.User) error
	GetUserByVerificationToken(token string) (*domain.User, error)
	GetUserCount() (int64, error)
	UpdatePasswordByEmail(email string, newPassword string) error
	StoreResetToken(email string, resetToken string) error
	InvalidateResetToken(email string) error
	GetResetTokenByEmail(email string) (string, error)
	DeleteUser(email string) error
	PromoteUser(email string) error
	DemoteUser(email string) error
}



type IPasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, plainPassword string) error
}
type ITokenService interface {
	GenerateToken(email string, id primitive.ObjectID, role, name string, expiryDuration int64) (string, error)
	ValidateToken(string) (map[string]interface{}, error)
	GetClaimsOfToken(token string) (map[string]interface{}, error)
}

type ITokenRepository interface {
	InsertRefresher(credential domain.Credential) error
	CheckRefresher(email, refresher string) error
	InvalidateResetToken(email string) error
	StoreResetToken(email string, resetToken string) error
	UpdateRefresher(email, refresher string) error
	DeleteRefresher(email, refresher string) error
	DeleteAllRefreshers(email string) error
}

type IMailService interface {
	SendVerificationEmail(to, token string) error
	SendPasswordResetEmail(email string, resetToken string) error
}
