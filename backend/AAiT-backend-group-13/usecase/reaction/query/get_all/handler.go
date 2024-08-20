package getallquery

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/reaction"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the Get comments query by blog ID.
type Handler struct {
	repo irepo.Reaction // Repository for comment.
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *[]reaction.Reaction] = &Handler{}

// New creates a new instance of Handler with the provided reaction repository.
func New(reactionRepo irepo.Reaction) *Handler {
	return &Handler{repo: reactionRepo}
}

// Handle processes the Get query by blog ID and returns the corresponding reactions.
func (h *Handler) Handle(id uuid.UUID) (*[]reaction.Reaction, error) {
	return h.repo.FindReactionByBlogId(id)
}
