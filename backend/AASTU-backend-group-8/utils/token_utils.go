package utils

import(
	"math/rand"
	"time"
)

func GenerateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < length; i++ {
		otp += string(rune(48 + rand.Intn(10)))
	}
	return otp
}

