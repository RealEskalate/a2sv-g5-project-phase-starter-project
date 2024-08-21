package dtos

type ForgotPassword struct {
	Email string `bson:"email" json:"email"`
}

type ResetPassword struct {
	Otp     string `bson:"otp" json:"otp"`
	NewPassword string `bson:"new_password" json:"password"`
}