package config

import "fmt"

func ConficBody(to string, token string) (string, string) {
	subject := "Email Verification"
	body := fmt.Sprintf(
	`
	<h2>Verify Your Email</h2>
	<p>Click the link below to verify your email:</p>
	<a href="http://localhost:8080/verify-email?token=%s">Verify Email</a>
	`,token)

	return subject,body
}
