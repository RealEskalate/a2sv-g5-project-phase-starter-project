package reactionqry

import (
	"log"

	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	cache "github.com/group13/blog/usecase/common/i_cache"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

// GetHandler is responsible for handling the Get reaction query by its ID.
type GetHandler struct {
	repo irepo.Reaction
	reactionCache cache.ICache
}

// Ensure Handler implements the IHandler interface
var _ icmd.IHandler[uuid.UUID, *models.Reaction] = &GetHandler{}

// NewGetHandler creates a new instance of Handler with the provided comment repository.
func NewGetHandler(reationRepo irepo.Reaction, cacheService *cache.ICache) *GetHandler {
	return &GetHandler{
		repo: reationRepo,
		reactionCache: *cacheService,
	}
}

// Handle processes the Get query by its ID and returns the corresponding reaction.
func (h *GetHandler) Handle(id uuid.UUID) (*models.Reaction, error) {
	cacheKey := id.String()
	
	if cachedReaction, err := h.reactionCache.Get(cacheKey); err == nil && cachedReaction != nil {
		if reactions, ok := cachedReaction.(*models.Reaction); ok {
			return reactions, nil
		}
	}


	var err error
	var reaction *models.Reaction

	reaction, err = h.repo.FindReactionById(id)
	
	if err != nil {
		return nil, err
	}

	// Store the retrieved blogs in the cache
	if err := h.reactionCache.Set(cacheKey, reaction); err != nil {
		// Log the caching error but don't fail the request
		log.Println("Failed to cache reactions:", err)
	}
	return reaction, nil
}
