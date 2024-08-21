package blogqry

import (
	"fmt"
	"log"
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
	cache "github.com/group13/blog/usecase/common/i_cache"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

const (
	defaultLimit = 10  // Default limit for the number of blogs to retrieve
	minLimit     = 5   // Minimum limit for the number of blogs to retrieve
	maxLimit     = 100 // Maximum limit for the number of blogs to retrieve
)

// GetMultipleHandler handles queries for retrieving multiple blogs.
type GetMultipleHandler struct {
	blogRepo irepo.Blog
	blogCache cache.ICache
}

// NewGetMultipleHandler creates a new instance of GetMultipleHandler with the provided blog repository.
func NewGetMultipleHandler(blogRepo irepo.Blog, cacheService *cache.ICache) *GetMultipleHandler {
	return &GetMultipleHandler{
		blogRepo: blogRepo,
		blogCache: *cacheService,
	}
}

// Handle processes a GetMultipleQuery to retrieve multiple blogs based on the provided query parameters.
func (h *GetMultipleHandler) Handle(query *GetMultipleQuery) ([]*models.Blog, error) {
	limit := h.validateLimit(query.limit)
	cacheKey := h.generateCacheKey(query)

	if cachedBlogs, err := h.blogCache.Get(cacheKey); err == nil && cachedBlogs != nil {
		if blogs, ok := cachedBlogs.([]*models.Blog); ok {
			return blogs, nil
		}
	}
	var blogs []*models.Blog
	var err error
	if query.userID == uuid.Nil {
		blogs, err = h.blogRepo.ListByTotalInteraction(query.lastSeenID, limit)
	} else {
		blogs, err = h.blogRepo.ListByAuthor(query.userID, query.lastSeenID, limit)
	}
	if err != nil {
		return nil, err
	}

	// Store the retrieved blogs in the cache
	if err := h.blogCache.Set(cacheKey, blogs); err != nil {
		// Log the caching error but don't fail the request
		log.Println("Failed to cache blogs:", err)
	}
	return blogs, nil
}

// validateLimit ensures the limit value is within the acceptable range.
func (h *GetMultipleHandler) validateLimit(requestedLimit int) int {
	if requestedLimit <= 0 {
		return defaultLimit
	}
	if requestedLimit < minLimit {
		return minLimit
	}
	if requestedLimit > maxLimit {
		return maxLimit
	}
	return requestedLimit
}


func (h *GetMultipleHandler) generateCacheKey(query *GetMultipleQuery) string {
	return fmt.Sprintf("blogs:user:%s:lastSeenID:%s:limit:%d", query.userID.String(), query.lastSeenID.String(), query.limit)
}