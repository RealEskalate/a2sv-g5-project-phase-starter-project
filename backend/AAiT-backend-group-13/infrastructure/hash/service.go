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
	log.Printf("Starting Decode with encryptedText: %s", encryptedText)
	
	key := []byte("thisis32bitlongpassphraseimusing!") // Same key used for encryption

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
	log.Printf("Starting Hash with word: %s", word)
	
	hashed, err := bcrypt.GenerateFromPassword([]byte(word), 13)
	if err != nil {
		log.Printf("Error generating hash: %v", err)
		return "", err
	}

	hashedWord := string(hashed)
	log.Printf("Successfully hashed word: %s", hashedWord)
	return hashedWord, nil
}

// Match implements ihash.Service.
func (s *Service) Match(hashedWord string, plainWord string) (bool, error) {
	log.Printf("Starting Match with hashedWord: %s and plainWord: %s", hashedWord, plainWord)
	
	err := bcrypt.CompareHashAndPassword([]byte(hashedWord), []byte(plainWord))
	if err != nil {
		log.Printf("Password mismatch or error: %v", err)
		return false, nil
	}

	log.Printf("Password matched successfully")
	return true, nil
}

var _ ihash.Service = &Service{}
