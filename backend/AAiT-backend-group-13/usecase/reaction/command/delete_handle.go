package reactioncmd

import (
	er "github.com/group13/blog/domain/errors"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// DeleteHandler is responsible for handling the delete reaction command.
type DeleteHandler struct {
	reactRepo irepo.Reaction
	blogRepo  irepo.Blog
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[*DeleteCommand, bool] = &DeleteHandler{}

// New creates a new instance of Handler with the provided reaction repository.
func New(blogRepo irepo.Blog, reactRepo irepo.Reaction) *DeleteHandler {
	return &DeleteHandler{
		blogRepo:  blogRepo,
		reactRepo: reactRepo,
	}
}

// Handle processes the delete reaction and removes the reaction from the repository and reduce its count in blog.
func (h *DeleteHandler) Handle(cmd *DeleteCommand) (bool, error) {
	r, err := h.reactRepo.FindReactionByUserIdAndBlogId(cmd.UserId, cmd.BlogId)

	if err != nil {
		return false, err
	}
	if r == nil {
		return false, er.NewNotFound("reaction not found")
	}

	blog, err := h.blogRepo.GetSingle(r.BlogID())

	if err != nil {
		return false, nil
	}

	err = h.reactRepo.Delete(r.ID())
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

	h.blogRepo.Save(blog)

	return true, nil
}
