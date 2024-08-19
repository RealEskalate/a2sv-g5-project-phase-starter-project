package domain

import "time"

type User struct {
	FirstName  string    `json:"firstname" bson:"firstname"`
	LastName   string    `json:"lastname" bson:"lastname"`
	Bio        string    `json:"bio" bson:"bio"`
	Avatar     string    `json:"avatar" bson:"avatar"`
	Username   string    `json:"username" bson:"username" binding:"required"`
	Password   string    `json:"password" bson:"password" binding:"required"`
	Email      string    `json:"email" bson:"email" binding:"required"`
	Role       string    `json:"role" bson:"role"`
	Address    string    `json:"address" bson:"address"`
	JoinedDate time.Time `json:"joined_date" bson:"joined_date"`
	IsVerified bool      `json:"is_verified" bson:"is_verified"`
}

type Token struct {
	Username  string `json:"username" bson:"username"`
	ExpiresAt int64  `json:"expires_at" bson:"expires_at"`
}

type UserRepository interface {
	CheckUsernameAndEmail(username, email string) error
	RegisterUser(user *User) error
	GetUserByUsernameorEmail(usernameoremail string) (*User, error)
	UpdateProfile(usernameoremail string, user *User) error
	Resetpassword(usernameoremail string, password string) error
	InsertToken(token *Token) error
	GetTokenByUsername(username string) (*Token, error)
	DeleteToken(username string) error
	CheckRoot() error
}

type UserUsecase interface {
	RegisterUser(user *User) error
	LoginUser(usernameoremail, password string) (string, string, error)
	UpdateProfile(user *User, claims *LoginClaims) error
	ResetPassword(tokenString string) error
	LogoutUser(username string) error
	ForgotPassword(email, newPassword string) error
	PromoteUser(username string, promoted bool, claims *LoginClaims) error
	VerifyUser(token string) error
	AddRoot() error
	RefreshToken(claims *LoginClaims) (string, error)
}
