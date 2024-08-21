package passwordreset

import (
	"strconv"
	"time"

	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// ValidateCodeHandler handles the validation of password reset codes.
type ValidateCodeHandler struct {
	userRepo    irepo.UserRepository
	jwtService  ijwt.Service
	hashService ihash.Service
}

// Ensure ValidateCodeHandler implements the icmd.IHandler interface.
var _ icmd.IHandler[*ValidateCodeCommand, string] = &ValidateCodeHandler{}

// NewValidateCodeHandler creates a new instance of ValidateCodeHandler.
func NewValidateCodeHandler(userRepo irepo.UserRepository, jwtService ijwt.Service, hashService ihash.Service) *ValidateCodeHandler {
	return &ValidateCodeHandler{
		userRepo:    userRepo,
		jwtService:  jwtService,
		hashService: hashService,
	}
}

// Handle validates the reset code and returns a JWT token if successful.
// It checks if the code is valid, unexpired, and associated with the correct user.
func (h *ValidateCodeHandler) Handle(cmd *ValidateCodeCommand) (string, error) {
	user, err := h.userRepo.FindByEmail(cmd.email)
	if err != nil {
		return "", er.NewUnauthorized(err.Error())
	}

	code := strconv.Itoa(cmd.code)

	resetCode := user.ResetCode()
	if resetCode == nil || resetCode.Expr.Before(time.Now()) {
		return "", er.NewUnauthorized("invalid or expired code")
	}

	isMatch, err := h.hashService.Match(resetCode.CodeHash, code)
	if err != nil {
		return "", err
	} else if !isMatch {
		return "", er.NewUnauthorized("invalid or expired code")
	}

	user.RemoveResetCode()
	if err := h.userRepo.Save(user); err != nil {
		return "", err
	}

	token, err := h.jwtService.Generate(user, ijwt.Reset)
	if err != nil {
		return "", err
	}

	return token, nil
}
