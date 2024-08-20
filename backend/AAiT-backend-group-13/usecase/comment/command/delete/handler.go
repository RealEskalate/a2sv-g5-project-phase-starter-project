package deletecom

// Package deletecmd provides the logic to delete a comment.
// It includes the handler to process the delete command.

import (
	"github.com/google/uuid"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the delete comment command.
type Handler struct {
	commentRepo irepo.Comment
	blogRepo    irepo.Blog
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, bool] = &Handler{}

// New creates a new instance of Handler with the provided blog repository.
func New(blogRepo irepo.Blog, commentRepo irepo.Comment) *Handler {
	return &Handler{
		blogRepo:    blogRepo,
		commentRepo: commentRepo,
	}
}

// Handle processes the delete command and removes the comment from the repository and reduce its count in blog.
func (h *Handler) Handle(id uuid.UUID) (bool, error) {
	c, err := h.commentRepo.GetCommentById(id)

	if err != nil {
		return false, err
	}

	blog, err := h.blogRepo.GetSingle(c.BlogId())

	if err != nil {
		return false, nil
	}

	err = h.commentRepo.Delete(id)
	if err != nil {
		return false, err
	}

	blog.UpdateCommentCount(false)

	return true, nil
}
