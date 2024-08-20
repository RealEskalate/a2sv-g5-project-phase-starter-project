package Domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name"`
	Email              string             `bson:"email"`
	Password           string             `bson:"password"`
	Role               string             `bson:"role"`
	ImageUrl           string             `bson:"image_url"`
	CreatedAt          time.Time          `bson:"created_at"`
	Verified           bool               `bson:"verified"`
	VerificationToken  string             `bson:"verification_token"`
	PasswordResetToken string             `bson:"password_reset_token"`
}

type Credential struct {
	Email     string `json:"email" bson:"email"`
	Refresher string `json:"refresher" bson:"refresher"`
}

type IUserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	VerifyUser(user *User) error
	GetUserByVerificationToken(token string) (*User, error)
	GetUserCount() (int64, error)
	UpdatePasswordByEmail(email string, newPassword string) error
	StoreResetToken(email string, resetToken string) error
	InvalidateResetToken(email string) error
	GetResetTokenByEmail(email string ) (string, error)
}

type IUserUseCase interface {
	Login(email, password string) (string, string, error)
	GetSingleUser(email string) (*User, error)
	RegisterUser(user *User) error
	VerifyEmail(token string) error
	RefreshToken(email, refresher string) (string, error)
	GenerateResetPasswordToken(email string) error
	ResetPassword(token string, newPassword string) error
	StoreToken(token string ) error
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
	InsertRefresher(credential Credential) error
	GetRefresher(email string) (string, error)
	InvalidateResetToken(email string) error
	StoreResetToken(email string, resetToken string) error
}

type IMailService interface {
	SendVerificationEmail(to, token string) error
	SendPasswordResetEmail(email string, resetToken string) error
}
