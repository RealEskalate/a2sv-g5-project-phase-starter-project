package reactioncmd

import (
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// UpdateHandler handles the logic for updating reactions on a blog.
type UpdateHandler struct {
	reactionRepo irepo.Reaction
	blogRepo     irepo.Blog
}

// Ensure UpdateHandler implements icmd.IHandler
var _ icmd.IHandler[*UpdateCommand, bool] = &UpdateHandler{}

// NewUpdateHandler creates a new instance of UpdateHandler.
func NewUpdateHandler(reactionRepo irepo.Reaction, blogRepo irepo.Blog) *UpdateHandler {
	return &UpdateHandler{
		reactionRepo: reactionRepo,
		blogRepo:     blogRepo,
	}
}

// Handle processes the command to update a reaction and adjusts the blog's reaction counts accordingly.
func (h *UpdateHandler) Handle(cmd *UpdateCommand) (bool, error) {
	// Find and delete existing reaction by user and blog ID
	existingReaction, err := h.reactionRepo.FindReactionByUserIdAndBlogId(cmd.UserId, cmd.BlogId)
	if err == nil && existingReaction != nil {
		if err = h.reactionRepo.Delete(existingReaction.ID()); err != nil {
			return false, err
		}

		// Adjust blog reaction count based on the old reaction
		blog, err := h.blogRepo.GetSingle(cmd.BlogId)
		if err != nil {
			return false, err
		}

		if existingReaction.IsLike() {
			if err = blog.UpdateLikeCount(false); err != nil {
				return false, err
			}
		} else {
			if err = blog.UpdateDislikeCount(false); err != nil {
				return false, err
			}
		}
	}

	// Create and save the new reaction
	newReaction := models.NewReaction(models.ReactionConfig{
		IsLike: cmd.IsLike,
		UserID: cmd.UserId,
		BlogID: cmd.BlogId,
	})
	if err = h.reactionRepo.Save(newReaction); err != nil {
		return false, err
	}

	// Update blog reaction count based on the new reaction
	blog, err := h.blogRepo.GetSingle(cmd.BlogId)
	if err != nil {
		return false, err
	}

	if cmd.IsLike {
		if err = blog.UpdateLikeCount(true); err != nil {
			return false, err
		}
	} else {
		if err = blog.UpdateDislikeCount(true); err != nil {
			return false, err
		}
	}

	if err = h.blogRepo.Save(blog); err != nil {
		return false, err
	}

	return true, nil
}
