package resetpassword

import (
	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	ijwt "github.com/group13/blog/usecase/common/i_jwt"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the logic for resetting a user's password.
type Handler struct {
	userrepo    irepo.User
	hashService ihash.Service
	jwtService  ijwt.Service
}

// New creates a new instance of Handler.
func New(userrepo irepo.User) *Handler {
	return &Handler{userrepo: userrepo}
}

// Handle processes the reset password command by validating the provided JWT token,
// verifying the user's identity, updating their password, and saving the changes to the repository.
// It returns a boolean indicating success, or an error if any step in the process fails.
func (h *Handler) Handle(cmd Command) (bool, error) {
	// Decode the provided JWT token.
	claims, err := h.jwtService.Decode(cmd.token)
	if err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	// Validate that the token is meant for a password reset.
	isForReset, ok := claims["is_for_reset"].(bool)
	if !ok || !isForReset {
		return false, er.NewUnauthorized("invalid token")
	}

	id, ok := claims["id"].(uuid.UUID)
	if !ok {
		return false, er.NewUnauthorized("invalid token")
	}

	user, err := h.userrepo.FindById(id)
	if err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	err = user.UpdatePassword(cmd.NewPassword, h.hashService)
	if err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	err = h.userrepo.Save(user)
	if err != nil {
		return false, er.NewUnauthorized(err.Error())
	}

	return true, nil
}
