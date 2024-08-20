package commentqry

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// GetHandler handles queries to retrieve a comment by its ID.
type GetHandler struct {
	repo irepo.Comment // Repository for comments.
}

// Ensure GetHandler implements icmd.IHandler
var _ icmd.IHandler[uuid.UUID, *models.Comment] = &GetHandler{}

// NewGetHandler creates a new instance of GetHandler with the provided comment repository.
func NewGetHandler(commentRepo irepo.Comment) *GetHandler {
	return &GetHandler{repo: commentRepo}
}

// Handle processes the query to retrieve a comment by its ID.
func (h *GetHandler) Handle(id uuid.UUID) (*models.Comment, error) {
	return h.repo.GetCommentById(id)
}
