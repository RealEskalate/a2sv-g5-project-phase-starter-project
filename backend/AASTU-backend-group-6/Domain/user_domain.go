package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Profile  string             `json:"profile"`
	GoogleID string             `json:"googleId"`
	Posts	[]Post           `json:"posts"`
	ResetToken string           `json:"resetToken"`
	Contact  string             `json:"contact"`	
	Bio	  string             `json:"bio"`		
	Role   string           `json:"roles"`
	
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

