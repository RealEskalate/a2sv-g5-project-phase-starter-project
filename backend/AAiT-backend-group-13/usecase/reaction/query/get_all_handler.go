package reactionqry

import (
	"log"

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	cache "github.com/group13/blog/usecase/common/i_cache"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// GetAllHandler is responsible for handling the Get comments query by blog ID.
type GetAllHandler struct {
	repo irepo.Reaction // Repository for comment.
	reactCache cache.ICache
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *[]models.Reaction] = &GetAllHandler{}

// NewGetAllHandler creates a new instance of Handler with the provided reaction repository.
func NewGetAllHandler(reactionRepo irepo.Reaction, cacheService *cache.ICache) *GetAllHandler {
	return &GetAllHandler{
		repo: reactionRepo,
		reactCache: *cacheService,

	}
}

// Handle processes the Get query by blog ID and returns the corresponding reactions.
func (h *GetAllHandler) Handle(id uuid.UUID) (*[]models.Reaction, error) {
	cacheKey := id.String()
	
	if cachedReactions, err := h.reactCache.Get(cacheKey); err == nil && cachedReactions != nil {
		if reactions, ok := cachedReactions.(*[]models.Reaction); ok {
			return reactions, nil
		}
	}


	var err error
	var reactions *[]models.Reaction

	reactions, err = h.repo.FindReactionByBlogId(id)
	
	if err != nil {
		return nil, err
	}

	// Store the retrieved blogs in the cache
	if err := h.reactCache.Set(cacheKey, reactions); err != nil {
		// Log the caching error but don't fail the request
		log.Println("Failed to cache reactions:", err)
	}
	return reactions, nil

}


