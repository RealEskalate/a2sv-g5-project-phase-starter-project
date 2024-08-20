package hash

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/group13/blog/domain/i_hash"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func (h *Service) Decode(encryptedText string) (string, error) {
	key := []byte("thisis32bitlongpassphraseimusing!") // Same key used for encryption

	ciphertext, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

// Hash implements ihash.Service.
func (s *Service) Hash(word string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(word), 13)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Match implements ihash.Service.
func (s *Service) Match(hashedWord string, plainWord string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedWord), []byte(plainWord))
	if err != nil {
		return false, nil
	}
	return true, nil
}

var _ ihash.Service = &Service{}
