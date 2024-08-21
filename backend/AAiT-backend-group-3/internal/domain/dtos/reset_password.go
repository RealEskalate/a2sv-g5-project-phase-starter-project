package dtos

type ForgotPassword struct {
	Email string `bson:"email" json:"email"`
}

type ResetPassword struct {
	NewPassword string `bson:"new_password" json:"new_password"`
}