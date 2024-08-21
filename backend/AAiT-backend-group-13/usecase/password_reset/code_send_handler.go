package passwordreset

import (
	"crypto/rand"
	"log"
	"time"

	ihash "github.com/group13/blog/domain/i_hash"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	iemail "github.com/group13/blog/usecase/common/i_email"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

const otpChars = "1234567890"

// SendcodeHandler handles the logic for sending password reset codes.
type SendcodeHandler struct {
	userRepo     irepo.UserRepository
	emailService iemail.Service
	hashService  ihash.Service
}

// Ensure Handler implements the icmd.IHandler interface.
var _ icmd.IHandler[string, time.Time] = &SendcodeHandler{}

// New creates a new instance of Handler.
func NewSendcodeHandler(userrepo irepo.UserRepository, emailService iemail.Service, hashService ihash.Service) *SendcodeHandler {
	return &SendcodeHandler{
		userRepo:     userrepo,
		emailService: emailService,
		hashService:  hashService,
	}
}

// Handle generates and sets a password reset code for the user identified by the given email.
// It returns the expiration time of the reset code or an error if the process fails.
func (h *SendcodeHandler) Handle(email string) (time.Time, error) {
	log.Println("Finding user by username")
	user, err := h.userRepo.FindByEmail(email)
	if err != nil {
		return time.Now(), err
	}

	exprTime := time.Now().Add(time.Minute * 17).UTC()
	code, err := generateOTP(6)
	if err != nil {
		return exprTime, err
	}

	if err = user.UpdateResetCode(code, exprTime, h.hashService); err != nil {
		return exprTime, err
	}

	err = h.userRepo.Save(user)
	if err != nil {
		return exprTime, err
	}

	mailReciver := []string{email}
	mail := iemail.NewResetPasswordEmail(mailReciver, code)
	err = h.emailService.Send(mail)
	if err != nil {
		return exprTime, err
	}

	return exprTime, err
}

func generateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
