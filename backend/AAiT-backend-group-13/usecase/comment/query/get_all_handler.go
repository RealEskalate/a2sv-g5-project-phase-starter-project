package commentqry

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the Get comments query by blog ID.
type GetAllHandler struct {
	repo irepo.Comment // Repository for comment.
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *[]models.Comment] = &GetAllHandler{}

// NewGetAllHandler creates a new instance of Handler with the provided comment repository.
func NewGetAllHandler(commmentRepo irepo.Comment) *GetAllHandler {
	return &GetAllHandler{repo: commmentRepo}
}

// Handle processes the Get query by blog ID and returns the corresponding comments.
func (h *GetAllHandler) Handle(id uuid.UUID) (*[]models.Comment, error) {
	return h.repo.GetCommentsByBlogId(id)
}
