package config

import "fmt"

func ConfigBody(token string) (string, string) {
	subject := "Email Verification"
	body := fmt.Sprintf(
	`
	<h2>Verify Your Email</h2>
	<hr>
	<p>Click the link below to verify your email:</p>
	<a href="http://localhost:8080/api/verify-email/%s">Verify Email</a>
	`,token)

	return subject,body
}
