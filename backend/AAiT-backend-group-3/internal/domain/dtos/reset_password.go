package dtos

type ForgotPassword struct {
	Email string `bson:"email" json:"email"`
}

type ResetPassword struct {
	NewPassword string `bson:"new_password" json:"new_password"`
}

type Profile struct {
	UserName   string             `bson:"username,omitempty" json:"username"`
	PhoneNum   string             `bson:"phone_num" json:"phone_num"`
	Bio        string             `bson:"bio" json:"bio"`
	ProfilePic string             `bson:"profile_pic" json:"profile_pic"`
}