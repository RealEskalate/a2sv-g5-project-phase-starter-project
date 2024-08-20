package resetcodevalidate

import (
	"time"

	er "github.com/group13/blog/domain/errors"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the validation of password reset codes.
type Handler struct {
	userrepo   irepo.UserRepository
	jwtService ijwt.Service
}

// Ensure Handler implements the icmd.IHandler interface.
var _ icmd.IHandler[Command, string] = &Handler{}

// New creates a new instance of Handler.
func New(userrepo irepo.UserRepository, jwtService ijwt.Service) *Handler {
	return &Handler{
		userrepo:   userrepo,
		jwtService: jwtService,
	}
}

// Handle validates the reset code from the Command and returns a JWT token if successful.
// It checks if the code is valid, unexpired, and associated with the correct user.
func (h *Handler) Handle(cmd Command) (string, error) {
	user, err := h.userrepo.FindById(cmd.Id)
	if err != nil {
		return "", er.NewUnauthorized(err.Error())
	}

	resetCode := user.ResetCode()
	if resetCode == nil || resetCode.Code != cmd.Code {
		return "", er.NewUnauthorized("invalid code")
	}

	exprTime := user.ResetCode().Expr
	if exprTime.After(time.Now()) {
		return "", er.NewUnauthorized("time expired")
	}

	if err = user.UpdateResetCode(nil); err != nil {
		return "", err
	}

	if err = h.userrepo.Save(user); err != nil {
		return "", err
	}

	token, err := h.jwtService.Generate(user, ijwt.Reset)
	if err != nil {
		return "", err
	}

	return token, nil
}
