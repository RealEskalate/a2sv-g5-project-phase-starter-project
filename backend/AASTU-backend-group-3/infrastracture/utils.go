package infrastracture

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"unicode"

	"gopkg.in/gomail.v2"
)

func IsValidEmail(email string) bool {

	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}

func IsValidPassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}


func SendActivationEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "bereket.meng@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Account Activation")


	m.SetBody("text/html", fmt.Sprintf("Click <a href=\"http://127.0.0.1:8080/auth/activate/?token=%s&Email=%s\">here</a> to activate your account.", token, email))

	d := gomail.NewDialer("smtp.gmail.com", 587, "bereket.meng@gmail.com", "xjbs vduu hkjd lqlf")

	
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}




func GenerateActivationToken() (string, error) {
	// Create a 32-byte random token
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	// Convert the token to a hex string
	return hex.EncodeToString(token), nil
}

func GenerateDeviceFingerprint(ip, userAgent string) string {
    data := ip + userAgent
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])
}

func SendResetLink(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "bereket.meng@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Link")


	m.SetBody("text/html", fmt.Sprintf("Click <a href=\"http://127.0.0.1:8080/auth/reset-password/%s\">here</a> to reset your password.", token))


	d := gomail.NewDialer("smtp.gmail.com", 587, "bereket.meng@gmail.com", "xjbs vduu hkjd lqlf")

	
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}