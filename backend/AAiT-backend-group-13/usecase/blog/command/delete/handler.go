// Package deletecmd provides the logic to delete a blog.
// It includes the handler to process the delete command.
package deletecmd

import (
	"github.com/google/uuid"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the delete blog command.
type Handler struct {
	repo irepo.Blog // Repository for task-related operations.
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, bool] = &Handler{}

// New creates a new instance of Handler with the provided blog repository.
func New(blogRepo irepo.Blog) *Handler {
	return &Handler{repo: blogRepo}
}

// Handle processes the delete command and removes the blog from the repository.
func (h *Handler) Handle(id uuid.UUID) (bool, error) {
	err := h.repo.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}