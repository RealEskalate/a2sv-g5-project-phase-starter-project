package blogcmd

import (
	"github.com/google/uuid"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the delete blog command.
type DeleteHandler struct {
	repo irepo.Blog
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, bool] = &DeleteHandler{}

// New creates a new instance of Handler with the provided blog repository.
func NewDeleteHandler(blogRepo irepo.Blog) *DeleteHandler {
	return &DeleteHandler{repo: blogRepo}
}

// Handle processes the delete command and removes the blog from the repository.
func (h *DeleteHandler) Handle(id uuid.UUID) (bool, error) {
	err := h.repo.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
