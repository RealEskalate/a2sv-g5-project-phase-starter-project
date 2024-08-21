package reactionqry

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// GetHandler is responsible for handling the Get reaction query by its ID.
type GetHandler struct {
	repo irepo.Reaction
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *models.Reaction] = &GetHandler{}

// NewGetHandler creates a new instance of Handler with the provided comment repository.
func NewGetHandler(reationRepo irepo.Reaction) *GetHandler {
	return &GetHandler{repo: reationRepo}
}

// Handle processes the Get query by its ID and returns the corresponding reaction.
func (h *GetHandler) Handle(id uuid.UUID) (*models.Reaction, error) {
	return h.repo.FindReactionById(id)
}
