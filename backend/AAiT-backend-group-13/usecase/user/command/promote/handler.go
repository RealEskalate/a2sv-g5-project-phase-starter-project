// Package promotcmd provides the logic for promoting a user to admin status.
// It includes the necessary command structure and a handler to execute the promotion.
package promotcmd

import (
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the promote command logic.
type Handler struct {
	userRepo irepo.UserRepository
}

// Ensure Handler implement icmd.Handler
var _ icmd.IHandler[*Command, bool] = &Handler{}

// New creates a new instance of the Handler with the provided user repository.
func New(userRepo irepo.UserRepository) *Handler {
	return &Handler{userRepo: userRepo}
}

// Promote promotes a user to admin status or demote user based on the provided command.
func (h *Handler) Handle(cmd *Command) (bool, error) {
	user, err := h.userRepo.ByUsername(cmd.Username)
	if err != nil {
		return false, err
	}

	user.UpdateAdminStatus(cmd.ToAdmin)
	if err := h.userRepo.Save(user); err != nil {
		return false, err
	}

	return true, nil
}
