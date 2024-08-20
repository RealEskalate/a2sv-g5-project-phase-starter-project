// Package deletereact provides the logic to delete a deletereact.
// It includes the handler to process the delete command.

package deletereact

import (
	"github.com/google/uuid"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the delete reaction command.
type Handler struct {
	reactRepo irepo.Reaction
	blogRepo  irepo.Blog
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, bool] = &Handler{}

// New creates a new instance of Handler with the provided reaction repository.
func New(blogRepo irepo.Blog, reactRepo irepo.Reaction) *Handler {
	return &Handler{
		blogRepo:  blogRepo,
		reactRepo: reactRepo,
	}
}

// Handle processes the delete reaction and removes the reaction from the repository and reduce its count in blog.
func (h *Handler) Handle(id uuid.UUID) (bool, error) {
	r, err := h.reactRepo.FindReactionById(id)

	if err != nil {
		return false, err
	}

	blog, err := h.blogRepo.GetSingle(r.BlogId())

	if err != nil {
		return false, nil
	}

	err = h.reactRepo.Delete(id)
	if err != nil {
		return false, err
	}

	if r.IsLike() {
		blog.UpdateLikeCount(false)
	} else {
		blog.UpdateDisLikeCount(false)
	}

	return true, nil
}
