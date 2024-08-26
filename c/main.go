package main

import (
	"crypto/rand"
	"fmt"
)

func main(){
	// Generate a 6-digit number OTP
	otp, err := GenerateOTP()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("OTP:", otp)

}
func GenerateOTP() (string, error) {
	// Create a 6-digit random OTP
	digits := "0123456789"
	var otp string
	for i := 0; i < 6; i++ {
		randomByte := make([]byte, 1)
		_, err := rand.Read(randomByte)
		if err != nil {
			return "", err
		}
		randomByte[0] = randomByte[0] % byte(len(digits))
		otp += string(digits[randomByte[0]])
	}
	return otp, nil
}