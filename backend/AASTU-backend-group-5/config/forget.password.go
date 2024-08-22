package config

import "fmt"

func ConfigFogetBody(token string, id string) (string, string) {
	subject := "Reset Password"
	body := fmt.Sprintf(
	`
	<h1>Reset password</h1>
	<hr>
	<p>Click the link below to reset password:</p>
	<a href="http://localhost:8080/api/forget-password/?id=%s&token=%s">Reset-Password</a>
	`,id,token)

	return subject,body
}