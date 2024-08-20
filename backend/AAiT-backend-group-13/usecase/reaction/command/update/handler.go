package updatereaction

import (
	"github.com/group13/blog/domain/models/reaction"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the logic for adding a new blog to the repository.
type Handler struct {
	reactionRepo irepo.Reaction
	bloRepo      irepo.Blog
}

// Ensure Handler implements icmd.IHandler
var _ icmd.IHandler[*Command, bool] = &Handler{}

// NewHandler creates a new instance of Handler with the given blog repository.
func NewHandler(reationRepo irepo.Reaction, blogRepo irepo.Blog) *Handler {
	return &Handler{reactionRepo: reationRepo}
}

// Handle processes the command to add a new blog to the repository.
func (h *Handler) Handle(cmd *Command) (bool, error) {

	c, err := h.reactionRepo.FindReactionByUserIdAndBlogId(cmd.UserId, cmd.BlogId)

	if err == nil {
		err := h.reactionRepo.Delete(c.ID())

		if err != nil {
			return false, err
		}

		blog, err := h.bloRepo.GetSingle(cmd.BlogId)

		if err != nil {
			return false, err
		}

		if c.IsLike() {
			blog.UpdateLikeCount(false)
		} else {
			blog.UpdateDisLikeCount(false)
		}

	}

	r := reaction.New(reaction.Config{
		IsLike: cmd.IsLike,
		UserId: cmd.UserId,
		BlogId: cmd.BlogId,
	})

	err = h.reactionRepo.Save(r)

	if err != nil {
		return false, err
	}

	blog, err := h.bloRepo.GetSingle(cmd.BlogId)

	if err != nil {
		return false, err
	}

	if cmd.IsLike {
		blog.UpdateLikeCount(true)
	} else {
		blog.UpdateDisLikeCount(true)
	}

	err = h.bloRepo.Save(blog)

	if err != nil {
		return false, err
	}

	return true, nil

}
