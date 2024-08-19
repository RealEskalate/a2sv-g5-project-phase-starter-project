package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role"`
	ImageUrl  string             `bson:"image_url"`
	CreatedAt time.Time          `bson:"created_at"`
	Verified  bool               `bson:"verified"`
	VerificationToken string     `bson:"verification_token"` 
}

type IUserRepository interface {
	CreateUser(user *User) error
    GetUserByEmail(email string) (*User, error)
    VerifyUser(user *User) error
	GetUserByVerificationToken(token string) (*User, error)
	GetUserCount() (int64, error) 
	
}

type IUserUseCase interface {
	RegisterUser(user *User) error
    VerifyEmail(token string) error
}


type IPasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, plainPassword string) error
}
type ITokenService interface {
	GenerateToken(claims []string) (string, error)
}

type ITokenRepo interface {
	InsertRefresher (id string) error 
}
