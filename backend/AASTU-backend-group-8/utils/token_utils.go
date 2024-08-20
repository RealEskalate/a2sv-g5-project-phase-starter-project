package utils

import(
	"math/rand"
	"regexp"
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

func ValidateEmail(email string) bool {
	// regular expression for validating an email
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}

//just simple validation
func ValidatePassword(password string) bool {
	return len(password) >= 8
}