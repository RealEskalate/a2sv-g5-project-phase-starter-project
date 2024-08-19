// Package promotcmd provides the logic for promoting a user to admin status.
// It includes the necessary command structure and a handler to execute the promotion.
package promotcmd

import (
	"github.com/google/uuid"
	usermodel "github.com/group13/blog/domain/models/user"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
)

type UserRepository interface {
	Save(*usermodel.User) error
	ByUsername(string) (*usermodel.User, error)
	ById(uuid.UUID) (bool, error)
}

// Handler handles the promote command logic.
type Handler struct {
	userRepo UserRepository
}

// Ensure Handler implement icmd.Handler
var _ icmd.IHandler[*Command, bool] = &Handler{}

// New creates a new instance of the Handler with the provided user repository.
func New(userRepo UserRepository) *Handler {
	return &Handler{userRepo: userRepo}
}

// Promote promotes a user to admin status or demote user based on the provided command.
func (h *Handler) Handle(cmd *Command) (bool, error) {
	user, err := h.userRepo.ByUsername(cmd.Username)
	if err != nil {
		return false, err
	}

	_, err = h.userRepo.ById(cmd.PromoterID)

	if err != nil {
		return false, err
	}

	user.UpdateAdminStatus(cmd.ToAdmin)
	if err := h.userRepo.Save(user); err != nil {
		return false, err
	}

	return true, nil
}
