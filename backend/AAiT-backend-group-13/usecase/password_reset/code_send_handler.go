package passwordreset

import (
	"log"
	"math/rand"
	"time"

	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// SendcodeHandler handles the logic for sending password reset codes.
type SendcodeHandler struct {
	userrepo irepo.UserRepository
}

// Ensure Handler implements the icmd.IHandler interface.
var _ icmd.IHandler[string, time.Time] = &SendcodeHandler{}

// New creates a new instance of Handler.
func NewSendcodeHandler(userrepo irepo.UserRepository) *SendcodeHandler {
	return &SendcodeHandler{userrepo: userrepo}
}

// Handle generates and sets a password reset code for the user identified by the given email.
// It returns the expiration time of the reset code or an error if the process fails.
func (h *SendcodeHandler) Handle(email string) (time.Time, error) {
	log.Println("Finding user by username")
	user, err := h.userrepo.FindByEmail(email)
	if err != nil {
		return time.Now(), er.NewUnauthorized(err.Error())
	}

	exprTime := time.Now().Add(time.Minute * 17).UTC()
	if err = user.UpdateResetCode(&models.ResetCode{
		Code: rand.Int63(),
		Expr: exprTime,
	}); err != nil {
		return exprTime, err
	}

	err = h.userrepo.Save(user)
	return exprTime, err
}
