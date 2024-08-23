package reactioncmd

import (
	"github.com/google/uuid"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// DeleteHandler is responsible for handling the delete reaction command.
type DeleteHandler struct {
	reactRepo irepo.Reaction
	blogRepo  irepo.Blog
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, bool] = &DeleteHandler{}

// New creates a new instance of Handler with the provided reaction repository.
func NewDeleteHandler(blogRepo irepo.Blog, reactRepo irepo.Reaction) *DeleteHandler {
	return &DeleteHandler{
		blogRepo:  blogRepo,
		reactRepo: reactRepo,
	}
}

// Handle processes the delete reaction and removes the reaction from the repository and reduce its count in blog.
func (h *DeleteHandler) Handle(id uuid.UUID) (bool, error) {
	r, err := h.reactRepo.FindReactionById(id)

	if err != nil {
		return false, err
	}

	blog, err := h.blogRepo.GetSingle(r.BlogID())

	if err != nil {
		return false, nil
	}

	err = h.reactRepo.Delete(id)
	if err != nil {
		return false, err
	}

	if r.IsLike() {
		if err = blog.UpdateLikeCount(false); err != nil {
			return false, err
		}
	} else {

		if err = blog.UpdateDislikeCount(false); err != nil {
			return false, err
		}
	}

	return true, nil
}
