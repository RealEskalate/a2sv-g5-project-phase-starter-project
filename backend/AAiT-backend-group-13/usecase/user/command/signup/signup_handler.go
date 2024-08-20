package signcommand

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	usermodel "github.com/group13/blog/domain/models/user"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	iemail "github.com/group13/blog/usecase/common/i_email"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
	result "github.com/group13/blog/usecase/user/result"
	"io"
	"time"
)

type SignUpHandler struct {
	repo         irepo.UserRepository
	jwtService   ijwt.Service
	hashService  ihash.Service
	emailService iemail.Service
}

type SignUpConfig struct {
	UserRepo     irepo.UserRepository
	JwtService   ijwt.Service
	HashService  ihash.Service
	EmailService iemail.Service
}

// Ensure Handler implements icmd.IHandler
var _ icmd.IHandler[*SignUpCommand, *result.SignUpResult] = &SignUpHandler{}

func NewSignUpHandler(config SignUpConfig) *SignUpHandler {
	return &SignUpHandler{
		repo:         config.UserRepo,
		jwtService:   config.JwtService,
		hashService:  config.HashService,
		emailService: config.EmailService,
	}
}

func (h *SignUpHandler) Handle(command *SignUpCommand) (*result.SignUpResult, error) {
	cfg := &usermodel.Config{

		Username:       command.username,
		Email:          command.email,
		PlainPassword:  command.password,
		FirstName:      command.firstName,
		LastName:       command.lastName,
		IsAdmin:        false,
		PasswordHasher: h.hashService,
	}

	user, err := usermodel.New(*cfg)
	if err != nil {
		return nil, err
	}

	res, err := h.repo.FindByUsername(user.Username())
	if res != nil {
		return nil, er.NewConflict("username taken")
	}

	if err != nil {
		return nil, err
	}

	res, err = h.repo.FindByEmail(user.Email())
	if res != nil {
		return nil, er.NewConflict("email already exists.")
	}

	if err != nil {
		return nil, err
	}

	password := user.PasswordHash()
	user.UpdatePassword(password, h.hashService)
	validationLink, err := h.GenerateValidationLink(*user)
	if err != nil {
		return nil, err
	}

	mails := []string{user.Email()}

	mail := iemail.NewSignUpEmail("", mails, validationLink)

	if err := h.emailService.Send(mail); err != nil {
		return nil, err
	}

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

func (h *SignUpHandler) GenerateValidationLink(user usermodel.User) (string, error) {
	userID := user.ID().String()
	expiryDay := time.Now().Add(time.Hour * 24).Format(time.RFC3339)
	username := user.Username()
	// Concatenate the values
	value := userID + "|" + expiryDay + "|" + username

	// Encrypt the concatenated value
	encryptedValue, err := h.encrypt(value)
	if err != nil {
		return "", er.NewUnexpected("failed to encrypt value")
	}

	validationLink := fmt.Sprintf("https://localhost:8080/validate?=%s", encryptedValue)
	return validationLink, nil
}

func (h *SignUpHandler) encrypt(plainText string) (string, error) {
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
