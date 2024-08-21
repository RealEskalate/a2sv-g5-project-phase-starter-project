package domain

import (
	"AAiT-backend-group-2/Infrastructure/dtos"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID    `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string    			`bson:"username" json:"username"`
	Email     string    			`bson:"email" json:"email" validate:"required,email"`
	Password  string    			`bson:"password" json:"password" validate:"required,min=8"`
	Role      string    			`bson:"role" json:"role"`
	CreatedAt time.Time 			`bson:"created_at" json:"created_at"`
	UpdateAt  time.Time     		`bson:"update_at" json:"update_at"`
	Profile   UserProfile 			`bson:"profile" json:"profile"`
	RefreshToken string 			`bson:"refresh_token" json:"refresh_token"`

}

type UserProfile struct {
	Bio 	 	string    	`bson:"bio" json:"bio"`
	ProfilePic 	string 		`bson:"profile_pic" json:"profile_pic"`
	ContactInfo string 		`bson:"contact_info" json:"contact_info"`
}

type UpdateData map[string]interface{}

type UserUsecase interface {
	GetAllUsers(c context.Context) ([]User, error)
	GetUserByID(c context.Context, id string) (*User, error)
	CreateUser(c context.Context, user User) error
	UpdateUser(c context.Context, id string, user *User) error
	DeleteUser(c context.Context, id string) error
	Login(c context.Context, loginDto *dtos.LoginDTO) (map[string]string, error)
	PromoteUser(c context.Context, id string) error
	DemoteAdmin(c context.Context, id string) error
	RefreshToken(c context.Context, token string) (string, error)
	Logout(c context.Context, id string) error
	ForgotPassword(c context.Context, userId, email string) error
	ResetPassword(c context.Context, userId string, passwordResetDto *dtos.PasswordResetDto) error
	ChangePassword(c context.Context, userId string, changePasswordDto *dtos.ChangePasswordDto) error
}

type UserRepository interface {
	FindAll(c context.Context) ([]User, error)
	FindByID(c context.Context, id string) (*User, error)
	FindByEmailOrUsername(c context.Context, emailOrUsername string) (*User, error)
	CountDocuments(c context.Context) (int64, error)
	Save(c context.Context, user User) error
	Update(c context.Context, id string, updateData UpdateData) error
	Delete(c context.Context, id string) error
	PromoteUser(c context.Context, id string, updateData UpdateData) error
	DemoteAdmin(c context.Context, id string, updateData UpdateData) error
	ForgotPassword(c context.Context, email ,token string) (string, error)
	ValidateResetToken(c context.Context, userID, token string) error
	InvalidateResetToken(c context.Context, userID string) error
}