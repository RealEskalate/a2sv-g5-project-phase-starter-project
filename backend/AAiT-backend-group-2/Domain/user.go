package domain

import "time"

type User struct {
	ID        string    `bson:"_id" json:"id"`
	Username  string    `bson:"username" json:"username"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	Role      string    `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdateAt  time.Time     `bson:"update_at" json:"update_at"`
	Profile   UserProfile `bson:"profile" json:"profile"`
}

type UserProfile struct {
	Bio 	 string    `bson:"bio" json:"bio"`
	ProfilePic string `bson:"profile_pic" json:"profile_pic"`
	ContactInfo string `bson:"contact_info" json:"contact_info"`
}