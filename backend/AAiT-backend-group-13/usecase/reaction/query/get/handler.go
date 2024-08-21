package getquery

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/reaction"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the Get reaction query by its ID.
type Handler struct {
	repo irepo.Reaction // Repository for comment.
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *reaction.Reaction] = &Handler{}

// New creates a new instance of Handler with the provided comment repository.
func New(reationRepo irepo.Reaction) *Handler {
	return &Handler{repo: reationRepo}
}

// Handle processes the Get query by its ID and returns the corresponding reaction.
func (h *Handler) Handle(id uuid.UUID) (*reaction.Reaction, error) {
	return h.repo.FindReactionById(id)
}
