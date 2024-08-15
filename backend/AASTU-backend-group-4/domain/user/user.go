package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	Role      string             `json:"role" bson:"role"` // e.g., "Admin" or "User"
}

type UserUsecase interface {
	RegisterUser(user User) error
	LoginUser(email, password string) (*User, error)
	LogOutUser(userID primitive.ObjectID) error
	ForgetPassword(email string) error
	UpdateUser(userID primitive.ObjectID, updatedUser *User) error
}

type UserRepository interface {
}

type UserDatabase interface {
}
