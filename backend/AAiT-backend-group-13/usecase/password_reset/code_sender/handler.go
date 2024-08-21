package resetcodesend

import (
	"math/rand"
	"time"

	er "github.com/group13/blog/domain/errors"
	usermodel "github.com/group13/blog/domain/models/user"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the logic for sending password reset codes.
type Handler struct {
	userrepo irepo.UserRepository
}

// Ensure Handler implements the icmd.IHandler interface.
var _ icmd.IHandler[string, time.Time] = &Handler{}

// New creates a new instance of Handler.
func New(userrepo irepo.UserRepository) *Handler {
	return &Handler{userrepo: userrepo}
}

// Handle generates and sets a password reset code for the user identified by the given email.
// It returns the expiration time of the reset code or an error if the process fails.
func (h *Handler) Handle(email string) (time.Time, error) {
	user, err := h.userrepo.FindByUsername(email)
	if err != nil {
		return time.Now(), er.NewUnauthorized(err.Error())
	}

	exprTime := time.Now().Add(time.Minute * 17).UTC()
	if err = user.UpdateResetCode(&usermodel.ResetCode{
		Code: rand.Int63(),
		Expr: exprTime,
	}); err != nil {
		return exprTime, err
	}

	err = h.userrepo.Save(user)
	return exprTime, err
}
