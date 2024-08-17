package domain

import ( 
	"time"
)

type OtpToken struct { 
	Email string `json:"email"`
	OTP string `json:"otp"`
	ExpiresAt time.Time `json:"expires_at"` 
}

