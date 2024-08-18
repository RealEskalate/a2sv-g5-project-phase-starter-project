package domain

import(
	// "meleket/infrastructure"
	"time"
)

type OTP struct {
	Otp       string             `bson:"otp"`
	Email	  string 			 `bson:"email"`
	Username  string 			 `bson:"username"`
	ExpiresAt time.Time          `bson:"expires_at"`

	// for later registration
	Password string             `json:"-"`
	Role     string             `json:"role"`
}

type OTPRequest struct {
	Otp   string `json:"otp"`
	Email string `json:"email"`
}


type OTPUsecaseInterface interface {
    GenerateAndSendOTP(user *User) error
    VerifyOTP(email, otp string)(*OTP, error)
}


type OTPRepositoryInterface interface {
    StoreOTP(otp *OTP) error
    GetOTPByEmail(email string) (*OTP, error)
    DeleteOTPByEmail(email string) error
}