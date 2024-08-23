package blogcmd

import (
	"log"

	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// AddHandler handles the logic for adding a new blog to the repository.
type AddHandler struct {
	repo irepo.Blog
}

// Ensure Handler implements icmd.Handler.
var _ icmd.IHandler[*AddCommand, *models.Blog] = &AddHandler{}

// NewAddHandler creates a new instance of Handler with the given blog repository.
func NewAddHandler(repo irepo.Blog) *AddHandler {
	return &AddHandler{repo: repo}
}

// Handle processes the command to add a new blog to the repository.
func (h *AddHandler) Handle(cmd *AddCommand) (*models.Blog, error) {
	blog, err := models.NewBlog(models.BlogConfig{
		Title:   cmd.title,
		Content: cmd.content,
		Tags:    cmd.tags,
		UserID:  cmd.userID,
	})
	log.Println("blog id is this", blog.UserID())
	if err != nil {
		return nil, err
	}

	err = h.repo.Save(blog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}
