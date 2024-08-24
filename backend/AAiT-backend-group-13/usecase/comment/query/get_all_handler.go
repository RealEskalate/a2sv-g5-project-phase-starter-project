package commentqry

import (
	"log"

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	cache "github.com/group13/blog/usecase/common/i_cache"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// Handler is responsible for handling the Get comments query by blog ID.
type GetAllHandler struct {
	repo         irepo.Comment // Repository for comment.
	commentCache cache.ICache
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *[]models.Comment] = &GetAllHandler{}

// NewGetAllHandler creates a new instance of Handler with the provided comment repository.
func NewGetAllHandler(commmentRepo irepo.Comment, cacheService cache.ICache) *GetAllHandler {
	return &GetAllHandler{
		repo:         commmentRepo,
		commentCache: cacheService,
	}
}


// Handle processes the Get query by blog ID and returns the corresponding comments.
func (h *GetAllHandler) Handle(id uuid.UUID) (*[]models.Comment, error) {
	cacheKey := id.String()

	if cachedComments, err := h.commentCache.Get(cacheKey); err == nil && cachedComments != nil {
		if comments, ok := cachedComments.(*[]models.Comment); ok {
			return comments, nil
		}
	}

	var err error
	var comments *[]models.Comment

	comments, err = h.repo.GetCommentsByBlogId(id)

	if err != nil {
		return nil, err
	}

	// Store the retrieved blogs in the cache
	if err := h.commentCache.Set(cacheKey, comments); err != nil {
		// Log the caching error but don't fail the request
		log.Println("Failed to cache comments:", err)
	}
	return comments, nil
}
