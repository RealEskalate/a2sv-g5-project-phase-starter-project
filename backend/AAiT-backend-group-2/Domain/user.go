package domain

import (
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
}

type UserProfile struct {
	Bio 	 	string    	`bson:"bio" json:"bio"`
	ProfilePic 	string 		`bson:"profile_pic" json:"profile_pic"`
	ContactInfo string 		`bson:"contact_info" json:"contact_info"`
}

type UserUsecase interface {
	GetAllUsers(c context.Context) ([]User, error)
	GetUserByID(c context.Context, id string) (*User, error)
	CreateUser(c context.Context, user User) error
	UpdateUser(c context.Context, id string, user *User) error
	DeleteUser(c context.Context, id string) error
	Login(c context.Context, user *User) (map[string]string, error)
	PromoteUser(c context.Context, id string) error
	DemoteAdmin(c context.Context, id string) error
	RefreshToken(c context.Context, token string) (map[string]string, error)
}

type UserRepository interface {
	FindAll(c context.Context) ([]User, error)
	FindByID(c context.Context, id string) (*User, error)
	Save(c context.Context, user User) error
	Update(c context.Context, id string, user *User) error
	Delete(c context.Context, id string) error
	Login(c context.Context, user *User) (map[string]string, error)
	PromoteUser(c context.Context, id string) error
	DemoteAdmin(c context.Context, id string) error
	RefreshToken(c context.Context, token string) (map[string]string, error)
}