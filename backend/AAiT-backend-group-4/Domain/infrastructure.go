package domain

type TokenInfrastructure interface {
	CreateAllTokens(user *User, accessSecret string, refreshSecret string,
		accessExpiry int, refreshExpiry int) (accessToken string, refreshToken string, err error)
	ValidateToken(tokenString string, secret string) (claims *JwtCustomClaims, err error)
	ExtractClaims(tokenString string, secret string) (map[string]interface{}, error)
	ExtractRoleFromToken(tokenString string, secret string) (string, error)
	CheckTokenExpiry(tokenString string, secret string) (bool, error)
	UpdateTokens(id string) (accessToken string, refreshToken string, err error)
	RemoveTokens(id string) error
}

type PasswordInfrastructure interface {
	HashPassword(password string) (string, error)
	ComparePasswords(password string, hashedPassword string) error
}

type OtpInfrastructure interface {
	CreateOTP(otp *UserOTPRequest) (otpcode string, err error)
	SendEmail(email string, subject string, key string, otp string) error
	SendPasswordResetEmail(email string, subject string, key string) error
}
