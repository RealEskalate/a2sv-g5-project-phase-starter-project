package domain

type SignUpRequest struct {
	Email             string `json:"email" validate:"required"`
	Username          string `json:"username" validate:"required"`
	Password          string `json:"password" validate:"required"`
	Profile_image_url string `json:"profile_image"`
	GoogleID          string `json:"googleId"`
	Contact           string `json:"contact"`
	Bio               string `json:"bio"`
}

type SignUpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}




type ForgotPasswordRequest struct { 
	Email string `json:"email" validate:"required"`
}

type ResetPasswordRequest struct {
	Password string `json:"password" validate:"required"`
}