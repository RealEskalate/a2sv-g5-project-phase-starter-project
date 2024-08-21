package blogcmd

import (
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// UpdateHandler handles the logic for updating an existing blog in the repository.
type UpdateHandler struct {
	repo irepo.Blog
}

// Ensure UpdateHandler implements icmd.Handler.
var _ icmd.IHandler[*UpdateCommand, *models.Blog] = &UpdateHandler{}

// NewUpdateHandler creates a new instance of UpdateHandler with the provided blog repository.
func NewUpdateHandler(blogRepo irepo.Blog) *UpdateHandler {
	return &UpdateHandler{repo: blogRepo}
}

// Handle processes the command to update an existing blog.
func (h *UpdateHandler) Handle(cmd *UpdateCommand) (*models.Blog, error) {
	blog, err := h.repo.GetSingle(cmd.id)
	if err != nil {
		return nil, err
	}

	if cmd.title != "" {
		err = blog.UpdateTitle(cmd.title)
		if err != nil {
			return nil, err
		}
	}

	if cmd.content != "" {
		err = blog.UpdateContent(cmd.content)
		if err != nil {
			return nil, err
		}
	}

	if cmd.tags != nil {
		err = blog.UpdateTags(cmd.tags)
		if err != nil {
			return nil, err
		}
	}

	err = h.repo.Save(blog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}
