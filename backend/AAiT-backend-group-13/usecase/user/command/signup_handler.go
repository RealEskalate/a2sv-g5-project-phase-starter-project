package usercmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	iemail "github.com/group13/blog/usecase/common/i_email"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	result "github.com/group13/blog/usecase/user/result"
)

// SignUpHandler handles user sign-up logic.
type SignUpHandler struct {
	repo         irepo.UserRepository
	jwtService   ijwt.Service
	hashService  ihash.Service
	emailService iemail.Service
}

// SignUpConfig holds the configuration for creating a SignUpHandler.
type SignUpConfig struct {
	UserRepo     irepo.UserRepository
	JwtService   ijwt.Service
	HashService  ihash.Service
	EmailService iemail.Service
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
	}
}

// Handle processes the sign-up command to register a new user.
// It creates a new user, checks for conflicts in username and email,
// generates a validation link, and sends a sign-up email.
func (h *SignUpHandler) Handle(command *SignUpCommand) (*result.SignUpResult, error) {
	cfg := models.UserConfig{
		Username:       command.username,
		Email:          command.email,
		PlainPassword:  command.password,
		FirstName:      command.firstName,
		LastName:       command.lastName,
		IsAdmin:        false,
		PasswordHasher: h.hashService,
	}

	user, err := models.NewUser(cfg)
	if err != nil {
		return nil, err
	}

	// Check if the username is already taken
	if res, err := h.repo.FindByUsername(user.Username()); res != nil || err != nil {
		if err == nil {
			err = er.NewConflict("username taken")
		}
		return nil, err
	}

	// Check if the email is already registered
	if res, err := h.repo.FindByEmail(user.Email()); res != nil || err != nil {
		if err == nil {
			err = er.NewConflict("email already exists")
		}
		return nil, err
	}

	// Generate a validation link
	validationLink, err := h.GenerateValidationLink(*user)
	if err != nil {
		return nil, err
	}

	// Send the sign-up email
	mails := []string{user.Email()}
	mail := iemail.NewSignUpEmail("", mails, validationLink)
	if err := h.emailService.Send(mail); err != nil {
		return nil, err
	}

	// Save the new user
	if err := h.repo.Save(user); err != nil {
		return nil, err
	}

	return &result.SignUpResult{
		ID:        user.ID(),
		Username:  user.Username(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		IsAdmin:   user.IsAdmin(),
	}, nil
}

// GenerateValidationLink creates a validation link for the user with encryption.
func (h *SignUpHandler) GenerateValidationLink(user models.User) (string, error) {
	userID := user.ID().String()
	expiryDay := time.Now().Add(time.Hour * 24).Format(time.RFC3339)
	username := user.Username()
	value := userID + "|" + expiryDay + "|" + username

	encryptedValue, err := encrypt(value)
	if err != nil {
		return "", er.NewUnexpected("failed to encrypt value")
	}

	validationLink := fmt.Sprintf("https://localhost:8080/validate?=%s", encryptedValue)
	return validationLink, nil
}

// encrypt encrypts a plain text value using AES-256 encryption.
func encrypt(plainText string) (string, error) {
	key := []byte("thisis32bitlongpassphraseimusing!") // 32 bytes key for AES-256

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plainText))

	return hex.EncodeToString(ciphertext), nil
}
