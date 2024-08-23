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
	GetAllUsers(c context.Context) ([]User, CodedError)
	GetUserByID(c context.Context, id string) (*User, CodedError)
	CreateUser(c context.Context, user User) CodedError
	UpdateUser(c context.Context, id string, user *User) CodedError
	DeleteUser(c context.Context, id string) CodedError
	Login(c context.Context, loginDto *dtos.LoginDTO) (map[string]string, CodedError)
	PromoteUser(c context.Context, id string) CodedError
	DemoteAdmin(c context.Context, id string) CodedError
	RefreshToken(c context.Context, token string) (string, CodedError)
	Logout(c context.Context, id string) CodedError
	ForgotPassword(c context.Context, userId, email string) CodedError
	ResetPassword(c context.Context, userId string, passwordResetDto *dtos.PasswordResetDto) CodedError
	ChangePassword(c context.Context, userId string, changePasswordDto *dtos.ChangePasswordDto) CodedError
	UpdateProfile(c context.Context, userId string, updateProfileDto *dtos.UpdateProfileDto) CodedError
}

type UserRepository interface {
	FindAll(c context.Context) ([]User, CodedError)
	FindByID(c context.Context, id string) (*User, CodedError)
	FindByEmailOrUsername(c context.Context, emailOrUsername string) (*User, CodedError)
	CountDocuments(c context.Context) (int64, CodedError)
	Save(c context.Context, user User) CodedError
	Update(c context.Context, id string, updateData UpdateData) CodedError
	Delete(c context.Context, id string) CodedError
	PromoteUser(c context.Context, id string, updateData UpdateData) CodedError
	DemoteAdmin(c context.Context, id string, updateData UpdateData) CodedError
	ForgotPassword(c context.Context, email ,token string) (string, CodedError)
	ValidateResetToken(c context.Context, userID, token string) CodedError
	InvalidateResetToken(c context.Context, userID string) CodedError
}