package addcmd

import (
	blogmodel "github.com/group13/blog/domain/models/blog"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the logic for adding a new blog to the repository.
type Handler struct {
	repo irepo.Blog
}

// Ensure Handler implements icmd.IHandler
var _ icmd.IHandler[*Command, *blogmodel.Blog] = &Handler{}

// NewHandler creates a new instance of Handler with the given blog repository.
func NewHandler(repo irepo.Blog) *Handler {
	return &Handler{repo: repo}
}

// Handle processes the command to add a new blog to the repository.
func (h *Handler) Handle(cmd *Command) (*blogmodel.Blog, error) {
	blog, err := blogmodel.New(blogmodel.Config{
		Title:   cmd.title,
		Content: cmd.content,
		Tags:    cmd.tags,
		UserId:  cmd.userId,
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
