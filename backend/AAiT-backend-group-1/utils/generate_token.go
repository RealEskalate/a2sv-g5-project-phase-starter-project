package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateTokenWithLength(legth int) (int, error) {
	code, errGen := rand.Int(rand.Reader, big.NewInt(int64(legth)))
	if errGen != nil {
		return 0, errGen
	}
	return int(code.Int64()), nil
}
