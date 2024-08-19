// Package updatecmd provides the logic to update an existing blog.
// It includes a command structure and a handler to process the update command.
package updatecmd

import (
	blogmodel "github.com/group13/blog/domain/models/blog"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

type Handler struct {
	repo irepo.Blog
}

// Ensure Handler implements icmd.IHandler
var _ icmd.IHandler[*Command, *blogmodel.Blog] = &Handler{}

// NewHandler creates a new instance of Handler with the provided blog repository.
func NewHandler(blogRepo irepo.Blog) *Handler {
	return &Handler{repo: blogRepo}
}

// HandleUpdate handles updating an existing blog.
func (h *Handler) Handle(cmd *Command) (*blogmodel.Blog, error) {
	blog, err := h.repo.GetSingle(cmd.id)
	if err != nil {
		return nil, err
	}

	err = blog.Update(blogmodel.Config{
		Title:   cmd.title,
		Content: cmd.content,
		Tags:    cmd.tags,
	})
	if err != nil {
		return nil, err
	}

	err = h.repo.Save(blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}
