package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

func GenerateToken(length int) (string, error) {
	token := make([]byte, length)

	_, err := rand.Read(token)
	if err != nil {
		return "", fmt.Errorf("error: %v", err.Error())
	}

	tokenString := base64.URLEncoding.EncodeToString(token)

	return strings.Trim(tokenString, "="), nil
}
