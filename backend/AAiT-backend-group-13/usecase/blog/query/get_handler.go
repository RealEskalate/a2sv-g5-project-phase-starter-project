package blogqry

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	iqry "github.com/group13/blog/usecase/common/cqrs/query"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

type GetHandler struct {
	blogRepo irepo.Blog
}

var _ iqry.IHandler[uuid.UUID, *models.Blog] = &GetHandler{}

func NewGetHandler(blogRepo irepo.Blog) *GetHandler {
	return &GetHandler{blogRepo: blogRepo}
}

func (g *GetHandler) Handle(id uuid.UUID) (*models.Blog, error) {
	return g.blogRepo.GetSingle(id)
}
