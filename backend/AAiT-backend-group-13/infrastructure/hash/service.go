package hash

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"

	"github.com/group13/blog/domain/i_hash"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func (h *Service) Decode(encryptedText string) (string, error) {
	key := []byte("thisis32bitlongpassphraseimusig!") // Same key used for encryption

	ciphertext, err := hex.DecodeString(encryptedText)
	if err != nil {
		log.Printf("Error decoding hex string: %v", err)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Error creating AES cipher: %v", err)
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	decodedText := string(ciphertext)
	log.Printf("Successfully decoded text: %s", decodedText)
	return decodedText, nil
}

// Hash implements ihash.Service.
func (s *Service) Hash(word string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(word), 13)
	if err != nil {
		log.Printf("Error generating hash: %v", err)
		return "", err
	}

	hashedWord := string(hashed)
	return hashedWord, nil
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
