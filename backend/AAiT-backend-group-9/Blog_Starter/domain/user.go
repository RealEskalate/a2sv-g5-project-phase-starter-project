package domain

import (
	"context"
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	UserID         primitive.ObjectID `json:"user_id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	Email          string             `json:"email" bson:"email"`
	Password       string             `json:"password" bson:"password"`
	Name           string             `json:"name" bson:"name"`
	Bio            string             `json:"bio" bson:"bio"`
	Role           string             `json:"role" bson:"role"`
	ContactInfo    ContactInfo        `json:"contact_info" bson:"contact_info"` //it requires contact_info in the user profile update requirement given
	IsActivated    bool               `json:"is_verified" bson:"is_verified"`   //useful for email verification
	ContactInfo    ContactInfo        `json:"contact_info" bson:"contact_info"` //it requires contact_info in the user profile update requirement given
	IsActivated    bool               `json:"is_verified" bson:"is_verified"`   //useful for email verification
	AccessToken    string             `json:"accessToken"`
	RefreshToken   string             `json:"refreshToken"`
	CreatedAt      time.Time          `json:"created_at" bson:"createtimestamp"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updatetimestamp"`
	ProfilePicture string             `json:"profile_picture" bson:"profile_picture"`
}

type ContactInfo struct {
	Phone     string `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`
}

type UserResponse struct {
	UserID         primitive.ObjectID `json:"user_id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	Email          string             `json:"email" bson:"email"`
	Name           string             `json:"name" bson:"name"`
	Bio            string             `json:"bio" bson:"bio"`
	ContactInfo     ContactInfo        `json:"contact_info" bson:"contact_info"`
	Role           string             `json:"role" bson:"role"`
	IsActivated    bool               `json:"is_verified" bson:"is_verified"` //useful for email verification
	IsActivated    bool               `json:"is_verified" bson:"is_verified"` //useful for email verification
	ProfilePicture string             `json:"profile_picture" bson:"profile_picture"`
}

type UserUpdate struct {
	Username    string      `json:"username" bson:"username"`
	Name        string      `json:"name" bson:"name"`
	Bio         string      `json:"bio" bson:"bio"`
	ContactInfo ContactInfo `json:"contact_info" bson:"contact_info"`
}

type AuthenticatedUser struct {
	UserID string
	Email  string
	Role   string
	UserID string
	Email  string
	Role   string
}

type UserRepository interface {
	CreateUser(c context.Context, user *User) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByID(c context.Context, userID string) (*User, error)
	GetAllUser(c context.Context) ([]*User, error)
	UpdateProfile(c context.Context, user *UserUpdate, userID string) (*User, error)
	UpdateProfilePicture(c context.Context, profilePicPath string, userID string) (*User, error)
	UpdatePassword(c context.Context, password, userID string) (*User, error)
	UpdateRole(c context.Context, role, userID string) (*User, error)
	UpdateSignup(c context.Context, user *User) error
	UpdateToken(c context.Context, accessToken, refreshToken, userID string) (*User, error)
	DeleteUser(c context.Context, userID string) error
}

type UserUsecase interface {
	GetUserByEmail(c context.Context, email string) (*UserResponse, error)
	GetUserByID(c context.Context, userID string) (*UserResponse, error)

	GetAllUser(c context.Context) ([]*UserResponse, error)              //superAdmin privilage
	DeleteUser(c context.Context, userID string, password string) error //superAdmin,Admin,User privilage
	PromoteUser(c context.Context, userID string) error                 //superAdmin privilage
	DemoteUser(c context.Context, userID string) error                  //superAdmin privilage

	UpdateUser(c context.Context, user *UserUpdate, userID string) (*UserResponse, error)
	UpdateProfilePicture(c context.Context, profilePicPath string, userID string) (*UserResponse, error)}

// ValidateEmail checks if the email format is valid
func ValidateEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword checks if the password meets the strength requirements
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	re := regexp.MustCompile(`[A-Z]`)
	if !re.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	re = regexp.MustCompile(`[a-z]`)
	if !re.MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	re = regexp.MustCompile(`[0-9]`)
	if !re.MatchString(password) {
		return errors.New("password must contain at least one digit")
	}
	re = regexp.MustCompile(`[!@#\$%\^&\*]`)
	if !re.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}
	return nil
}


// Validate validates the UserSignUp struct
func (u *UserSignUp) Validate() error {
    if err := ValidateEmail(u.Email); err != nil {
        return err
    }
    if err := ValidatePassword(u.Password); err != nil {
        return err
    }
    return nil
}

// Validate validates the UserLogin struct
func (u *UserLogin) Validate() error {
    if err := ValidateEmail(u.Email); err != nil {
        return err
    }
    if err := ValidatePassword(u.Password); err != nil {
        return err
    }
    return nil
}