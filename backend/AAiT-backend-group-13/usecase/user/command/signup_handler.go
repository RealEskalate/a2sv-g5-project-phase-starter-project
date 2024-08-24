package usercmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"time"

	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	iemail "github.com/group13/blog/usecase/common/i_email"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	icache "github.com/group13/blog/usecase/common/i_cache"
	result "github.com/group13/blog/usecase/user/result"
)

// SignUpHandler handles user sign-up logic.
type SignUpHandler struct {
	repo         irepo.UserRepository
	jwtService   ijwt.Service
	hashService  ihash.Service
	emailService iemail.Service
	cacheService icache.ICache
}

// SignUpConfig holds the configuration for creating a SignUpHandler.
type SignUpConfig struct {
	UserRepo     irepo.UserRepository
	JwtService   ijwt.Service
	HashService  ihash.Service
	EmailService iemail.Service
	CacheService icache.ICache
}

// Ensure SignUpHandler implements icmd.IHandler.
var _ icmd.IHandler[*SignUpCommand, *result.SignUpResult] = &SignUpHandler{}

// NewSignUpHandler creates a new instance of SignUpHandler with the provided configuration.
func NewSignUpHandler(config SignUpConfig) *SignUpHandler {
	return &SignUpHandler{
		repo:         config.UserRepo,
		jwtService:   config.JwtService,
		hashService:  config.HashService,
		emailService: config.EmailService,
		cacheService: config.CacheService,
	}
}

// Handle processes the sign-up command to register a new user.
// It creates a new user, checks for conflicts in email,
// generates a validation link, and sends a sign-up email.
func (h *SignUpHandler) Handle(command *SignUpCommand) (*result.SignUpResult, error) {
	log.Printf("Starting sign-up process for user %s -- SignUpHandler", command.username)

	user, err := models.NewUser(models.UserConfig{
		Username:       command.username,
		Email:          command.email,
		PlainPassword:  command.password,
		FirstName:      command.firstName,
		LastName:       command.lastName,
		IsAdmin:        false,
		PasswordHasher: h.hashService,
	})
	if err != nil {
		log.Printf("Error creating new user: %v", err)
		return nil, err
	}
	log.Println("New user created")

	// Check if the username is already taken
	res, err := h.repo.FindByUsername(user.Username())
	if err != nil {
		if err != er.UserNotFound {
			log.Printf("Error finding user by username: %v -- SignUpHandler", err.Error())
			return nil, err
		}
	} else if res != nil {
		log.Printf("Username %s is already taken -- SignUpHandler", user.Username())
		return nil, er.NewConflict("username taken")
	}

	log.Printf("Username %s is available", user.Username())

	// Check if the email is already registered
	res, err = h.repo.FindByEmail(user.Email())
	if err != nil {
		if err != er.UserNotFound {
			log.Printf("Error finding user by email: %v -- SignUpHandler", err)
			return nil, err
		}
	} else if res != nil {
		log.Printf("Email %s is already registered -- SignUpHandler", user.Email())
		return nil, er.NewConflict("email already exists")
	}

	log.Printf("Email %s is available", user.Email())

	// Generate a validation link
	validationLink, err := h.generateValidationLink(*user)
	if err != nil {
		log.Printf("Error generating validation link: %v", err)
		return nil, err
	}
	log.Println("Validation link generated")
	log.Printf("user id is this: %s", user.ID())

	// Send the sign-up email
	mails := []string{user.Email()}
	mail := iemail.NewSignUpEmail("", mails, validationLink)
	if err := h.emailService.Send(mail); err != nil {
		log.Printf("Error sending sign-up email: %v", err)
		return nil, err
	}
	log.Println("Sign-up email sent")

	// Save the new user
	if err := h.repo.Save(user); err != nil {
		log.Printf("Error saving new user: %v", err)
		return nil, err
	}
	log.Printf("New user %s saved successfully -- SignUpHandler", user.Username())

	return &result.SignUpResult{
		ID:        user.ID(),
		Username:  user.Username(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		IsAdmin:   user.IsAdmin(),
	}, nil
}

// generateValidationLink creates a validation link for the user with encryption.
func (h *SignUpHandler) generateValidationLink(user models.User) (string, error) {
	log.Printf("Generating validation link for user %s", user.Username())

	userID := user.ID().String()
	expiryDay := time.Now().Add(time.Hour * 24).Format(time.RFC3339)
	username := user.Username()
	value := userID + "|" + expiryDay + "|" + username

	encryptedValue, err := encrypt(value)
	if err != nil {
		log.Printf("Error encrypting validation link: %v", err)
		return "", er.NewUnexpected("failed to encrypt value")
	}

	validationLink := fmt.Sprintf("http://localhost:8080/api/v1/auth/validateEmail?secret?=%s", encryptedValue)
	log.Printf("Validation link generated: %s", validationLink)
	return validationLink, nil
}

// encrypt encrypts a plain text value using AES-256 encryption.
func encrypt(plainText string) (string, error) {
	log.Println("Encrypting value")

	key := []byte("thisis32bitlongpassphraseimusig!") // 32 bytes key for AES-256

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Error creating AES cipher: %v", err)
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Printf("Error reading random bytes: %v", err)
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plainText))

	encryptedValue := hex.EncodeToString(ciphertext)
	log.Println("Value encrypted successfully")
	return encryptedValue, nil
}

