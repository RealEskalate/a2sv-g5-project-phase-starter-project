package utils

import (
	"math/rand"
	"fmt"
)

// GenerateRecoveryToken generates a random 64-character token.
func GenerateRecoveryToken() string {
	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	recoveryTokenRune := make([]rune, 64)
	for i := 0; i < 64; i++ {
		recoveryTokenRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes))]
	}
	return string(recoveryTokenRune)
}

// GenerateRecoveryLink creates a password recovery link using the base URL and the generated token.
func GenerateRecoveryLink(baseURL, username string, recoveryToken string) string {
	recoveryLink := fmt.Sprintf("http://%s/recover-password?token=%s&user=%s", baseURL, recoveryToken, username)
	return recoveryLink
}
