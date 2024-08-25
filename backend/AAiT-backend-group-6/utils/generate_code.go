package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomCode(length int) (string, error) {
    bytes := make([]byte, length)
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes)[:length], nil
}

