package dto

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserCreateRequestDTO struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
    Role     string `json:"role"`
}

type UserResponse struct {
    Username string             `json:"username"`
    Email    string             `json:"email"`
    Role     string             `json:"role"`
}
