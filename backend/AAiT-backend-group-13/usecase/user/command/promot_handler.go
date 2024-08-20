package usercmd

import (
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// PromoteHandler handles the logic for promoting or demoting a user.
type PromoteHandler struct {
	userRepo irepo.UserRepository
}

// Ensure PromoteHandler implements icmd.IHandler
var _ icmd.IHandler[*PromoteCommand, bool] = &PromoteHandler{}

// NewPromoteHandler creates a new instance of PromoteHandler with the provided user repository.
func NewPromoteHandler(userRepo irepo.UserRepository) *PromoteHandler {
	return &PromoteHandler{userRepo: userRepo}
}

// Handle processes the PromoteCommand to update a user's admin status.
func (h *PromoteHandler) Handle(cmd *PromoteCommand) (bool, error) {
	user, err := h.userRepo.FindByUsername(cmd.username)
	if err != nil {
		return false, err
	}

	user.UpdateAdminStatus(cmd.toAdmin)
	if err := h.userRepo.Save(user); err != nil {
		return false, err
	}

	return true, nil
}
