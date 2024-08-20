package services

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateOTP() string {
	bytes := make([]byte, 3)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
