package commentcmd

import (
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the logic for adding a new blog to the repository.
type AddHandler struct {
	blogRepo    irepo.Blog
	commentRepo irepo.Comment
	userRepo    irepo.UserRepository
}

// Ensure Handler implements icmd.IHandler
var _ icmd.IHandler[*AddCommand, *models.Comment] = &AddHandler{}

// NewHandler creates a new instance of Handler with the given blog repository.
func NewHandler(blogRepo irepo.Blog, userRepo irepo.UserRepository, commentRepo irepo.Comment) *AddHandler {
	return &AddHandler{blogRepo: blogRepo, commentRepo: commentRepo, userRepo: userRepo}
}

// Handle processes the command to add a new blog to the repository.
func (h *AddHandler) Handle(cmd *AddCommand) (*models.Comment, error) {
	comment, err := models.NewComment(models.CommentConfig{
		Content: cmd.content,
		UserID:  cmd.userID,
		BlogID:  cmd.blogID,
	})

	if err != nil {
		return nil, err
	}

	blog, err := h.blogRepo.GetSingle(cmd.blogID)

	if err != nil {
		return nil, err
	}

	_, err = h.userRepo.FindById(cmd.userID)

	if err != nil {
		return nil, err
	}

	err = h.commentRepo.Save(*comment)

	if err != nil {
		return nil, err
	}

	err = blog.UpdateCommentCount(true)
	if err != nil {
		return nil, err
	}

	err = h.blogRepo.Save(blog)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
