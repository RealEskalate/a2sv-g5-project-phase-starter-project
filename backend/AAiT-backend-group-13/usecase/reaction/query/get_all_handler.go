package reactionqry

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// GetAllHandler is responsible for handling the Get comments query by blog ID.
type GetAllHandler struct {
	repo irepo.Reaction // Repository for comment.
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *[]models.Reaction] = &GetAllHandler{}

// NewGetAllHandler creates a new instance of Handler with the provided reaction repository.
func NewGetAllHandler(reactionRepo irepo.Reaction) *GetAllHandler {
	return &GetAllHandler{repo: reactionRepo}
}

// Handle processes the Get query by blog ID and returns the corresponding reactions.
func (h *GetAllHandler) Handle(id uuid.UUID) (*[]models.Reaction, error) {
	return h.repo.FindReactionByBlogId(id)
}
