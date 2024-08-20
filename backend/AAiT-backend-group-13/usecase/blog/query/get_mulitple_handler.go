package blogqry

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models"
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
}

// NewGetMultipleHandler creates a new instance of GetMultipleHandler with the provided blog repository.
func NewGetMultipleHandler(blogRepo irepo.Blog) *GetMultipleHandler {
	return &GetMultipleHandler{blogRepo: blogRepo}
}

// Handle processes a GetMultipleQuery to retrieve multiple blogs based on the provided query parameters.
func (h *GetMultipleHandler) Handle(query *GetMultipleQuery) ([]*models.Blog, error) {
	limit := h.validateLimit(query.limit)

	if query.userID == uuid.Nil {
		return h.blogRepo.ListByTotalInteraction(query.lastSeenID, limit)
	}

	return h.blogRepo.ListByAuthor(query.userID, query.lastSeenID, limit)
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
