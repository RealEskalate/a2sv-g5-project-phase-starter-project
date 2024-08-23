package otputil

import (
	"crypto/rand"
	"math/big"
)


func GenerateOTP() (string, error) {
	otpLength := 8
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	otp := make([]byte, otpLength)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range otp {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		otp[i] = charset[randomIndex.Int64()]
	}
	return string(otp), nil
}
