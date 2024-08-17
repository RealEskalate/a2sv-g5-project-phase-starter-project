package domain

type SignupRequest struct {
	First_Name string  `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name  string  `json:"last_name" validate:"required,min=2,max=100"`
	User_Name  string  `json:"user_name" validate:"required,min=5"`
	Email      string  `json:"email" validate:"required,email"`
	Password   string  `json:"password" validate:"required,min=6"`
	Phone      *string `json:"phone"`
	Bio        *string `json:"bio"`
}
