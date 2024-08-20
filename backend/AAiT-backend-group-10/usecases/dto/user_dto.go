package dto

type RegisterUserDTO struct {
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio" bson:"bio"`
}

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TokenResponseDTO struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
