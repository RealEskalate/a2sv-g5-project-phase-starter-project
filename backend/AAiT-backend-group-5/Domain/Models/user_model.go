package models

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID          string `bson:"_id,omitempty" json:"id"`
	Username    string `bson:"username" json:"username" validate:"required,min=3,max=30"`
	Name        string `bson:"name" json:"name" validate:"required"`
	Email       string `bson:"email" json:"email" validate:"required,email"`
	Password    string `bson:"password" json:"password"`
	Role        Role   `bson:"role" json:"role"`
	Bio         string `bson:"bio" json:"bio"`
	ImageKey    string `bson:"image_key" json:"image_key"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
}
