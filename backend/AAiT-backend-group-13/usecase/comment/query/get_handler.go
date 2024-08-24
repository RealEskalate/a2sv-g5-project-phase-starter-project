package commentqry

import (
	"log"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	cache "github.com/group13/blog/usecase/common/i_cache"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// GetHandler handles queries to retrieve a comment by its ID.
type GetHandler struct {
	repo         irepo.Comment // Repository for comments.
	commentCache cache.ICache
}

// Ensure GetHandler implements icmd.IHandler
var _ icmd.IHandler[uuid.UUID, *models.Comment] = &GetHandler{}

// NewGetHandler creates a new instance of GetHandler with the provided comment repository.
func NewGetHandler(commentRepo irepo.Comment, commentCache cache.ICache) *GetHandler {
	return &GetHandler{
		repo:         commentRepo,
		commentCache: commentCache,
	}
}

// Handle processes the query to retrieve a comment by its ID.
func (h *GetHandler) Handle(id uuid.UUID) (*models.Comment, error) {
	cacheKey := id.String()

	if cachedComment, err := h.commentCache.Get(cacheKey); err == nil && cachedComment != nil {
		if comment, ok := cachedComment.(*models.Comment); ok {
			return comment, nil
		}
	}

	var err error
	var comment *models.Comment

	comment, err = h.repo.GetCommentById(id)

	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, er.NewNotFound("comment not found")
	}

	// Store the retrieved blogs in the cache
	if err := h.commentCache.Set(cacheKey, comment); err != nil {
		// Log the caching error but don't fail the request
		log.Println("Failed to cache comment:", err)
	}

	return comment, nil

}
