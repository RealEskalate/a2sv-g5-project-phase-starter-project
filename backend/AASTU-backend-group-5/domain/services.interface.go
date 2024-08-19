package domain

type PasswordService interface {
	HashPasword(string) (string, error)
	ComparePassword(string, string) (bool, error)
}

type TokenService interface {
	GenerateAccessToken(user User) (string, error)
	GenerateRefreshToken(user User) (string, error)
	ValidateAccessToken(token string) (*User, error)
	ValidateRefreshToken(token string) (*User, error)
	GeneratePasswordResetToken(user User) (string, error)   
	ValidatePasswordResetToken(token string) (*User, error) 
	GenerateVerificationToken(user User) (string, error)  
	ValidateVerificationToken(token string) (*User, error) 
}

type OAuthService interface {
	VerifyOAuthToken(provider, token string) (string, error)
	GetUserDataFromProvider(provider, token string) (User, error)
}