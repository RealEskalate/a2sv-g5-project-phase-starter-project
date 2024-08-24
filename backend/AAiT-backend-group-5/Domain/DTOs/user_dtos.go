package dtos

type CreateAccountRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type CreateAccountResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type PasswordResetRequest struct {
	Email string `bson:"email" json:"email" validate:"required,email"`
}

type SetUpPasswordRequest struct {
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LogoutRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type Response struct {
	Message string `json:"message"`
}

type ProfileUpdateRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Password    string `bson:"password" json:"password"`
	Bio         string `bson:"bio" json:"bio"`
	ImageKey    string `bson:"image_url" json:"image_url"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
}

type Profile struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Bio         string `bson:"bio" json:"bio"`
	ImageURL    string `bson:"image_url" json:"image_url"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
}

type OAuthRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	PhoneNumber  string `json:"phone_number"`
	Bio          string `json:"bio"`
	Password     string `bson:"password" json:"password"`
}
