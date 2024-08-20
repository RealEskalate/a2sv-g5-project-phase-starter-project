package addcom

import (
	"github.com/group13/blog/domain/models/comment"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler handles the logic for adding a new blog to the repository.
type Handler struct {
	blogRepo    irepo.Blog
	commentRepo irepo.Comment
}

// Ensure Handler implements icmd.IHandler
var _ icmd.IHandler[*Command, *comment.Comment] = &Handler{}

// NewHandler creates a new instance of Handler with the given blog repository.
func NewHandler(blogRepo irepo.Blog, commentRepo irepo.Comment) *Handler {
	return &Handler{blogRepo: blogRepo, commentRepo: commentRepo}
}

// Handle processes the command to add a new blog to the repository.
func (h *Handler) Handle(cmd *Command) (*comment.Comment, error) {
	comment, err := comment.New(comment.Config{
		Content: cmd.content,
		UserId:  cmd.userId,
		BlogId:  cmd.blogId,
	})

	if err != nil {
		return nil, err
	}

	blog, err := h.blogRepo.GetSingle(cmd.blogId)

	if err != nil {
		return nil, err
	}

	err = h.commentRepo.Save(comment)

	if err != nil {
		return nil, err
	}
	
	blog.UpdateCommentCount(true)

	err = h.blogRepo.Save(blog)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
