package commentcmd

import (
	"github.com/google/uuid"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the delete comment command.
type DeleteHandler struct {
	commentRepo irepo.Comment
	blogRepo    irepo.Blog
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, bool] = &DeleteHandler{}

// New creates a new instance of Handler with the provided blog repository.
func New(blogRepo irepo.Blog, commentRepo irepo.Comment) *DeleteHandler {
	return &DeleteHandler{
		blogRepo:    blogRepo,
		commentRepo: commentRepo,
	}
}

// Handle processes the delete command and removes the comment from the repository and reduce its count in blog.
func (h *DeleteHandler) Handle(id uuid.UUID) (bool, error) {
	c, err := h.commentRepo.GetCommentById(id)

	if err != nil {
		return false, err
	}
	if c == nil {
		return false, err
	}

	blog, err := h.blogRepo.GetSingle(c.BlogID())

	if err != nil {
		return false, nil
	}

	err = h.commentRepo.Delete(id)
	if err != nil {
		return false, err
	}

	err = blog.UpdateCommentCount(false)
	if err != nil {
		return false, err
	}

	return true, nil
}
