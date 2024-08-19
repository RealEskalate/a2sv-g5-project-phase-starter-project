package resetpassword

import (
	er "github.com/group13/blog/domain/errors"
	ihash "github.com/group13/blog/domain/i_hash"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the logic for resetting a user's password.
type Handler struct {
	userrepo    irepo.User
	hashService ihash.Service
}

// New creates a new instance of Handler.
func New(userrepo irepo.User) *Handler {
	return &Handler{userrepo: userrepo}
}

// Handle processes the reset password command by validating the user, updating their password,
// and saving the updated user information to the repository. It returns a boolean indicating success,
// or an error if the operation fails.
func (h *Handler) Handle(cmd Command) (bool, error) {
	user, err := h.userrepo.FindById(cmd.Id)
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
