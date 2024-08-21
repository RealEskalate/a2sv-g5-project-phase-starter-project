package config

import "fmt"

func ConfigFogetBody(token string, id string) (string, string) {
	subject := "Email Verification"
	body := fmt.Sprintf(
	`
	<h2>Verify Your Email</h2>
	<hr>
	<p>Click the link below to reset password:</p>
	<a href="http://localhost:8080/api/forget-password/?id=%s&token=%s">Forget-Password</a>
	`,id,token)

	return subject,body
}