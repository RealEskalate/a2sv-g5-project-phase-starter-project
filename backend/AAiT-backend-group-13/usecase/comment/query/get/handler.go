// Package getqry provides the logic to retrieve a specific comment by its ID from the repository.
// It includes a handler that processes the Get query and returns the corresponding comment.
package getcomqry

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/comment"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the Get comment query by its ID.
type Handler struct {
	repo irepo.Comment // Repository for comment.
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *comment.Comment] = &Handler{}

// New creates a new instance of Handler with the provided comment repository.
func New(commmentRepo irepo.Comment) *Handler {
	return &Handler{repo: commmentRepo}
}

// Handle processes the Get query by its ID and returns the corresponding comment.
func (h *Handler) Handle(id uuid.UUID) (*comment.Comment, error) {
	return h.repo.GetCommentById(id)
}
