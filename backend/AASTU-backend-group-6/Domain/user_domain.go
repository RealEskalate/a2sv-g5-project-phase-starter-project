package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {

	ID                primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Email             string             `json:"email"`
	Username          string             `json:"username"`
	Password          string             `json:"password"`
	Profile_image_url string             `json:"profile_image"`
	GoogleID          string             `json:"googleId"`
	Posts             []Post             `json:"posts"`
	RefreshToken      string             `json:"refreshToken"`
	AccessToken       string             `json:"accessToken"`
	Contact           string             `json:"contact"`
	Bio               string             `json:"bio"`
	Role              string             `json:"roles"`
	Comments          []Comment          `json:"comments"`
}



type UserRepository interface {
	CreateUser(user User) (User, error)
	FindUserByEmail(email string) (User, error)
	FindUserByUsername(username string) (User, error)
	FindUserByID(id string) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id string) error
	ForgotPassword(email string, token string) error
}

type UserUseCase interface { 
	CreateUser(user User) interface{}
	FindUserByEmail(email string) interface{}
	FindUserByUsername(username string) interface{}
	FindUserByID(id string) interface{}
	UpdateUser(user User) interface{}
	DeleteUser(id string) interface{}
	ForgotPassword(email string, token string) interface{}
	
}
