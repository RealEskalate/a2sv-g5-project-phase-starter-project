package blogqry

import (
	"github.com/google/uuid"
	"github.com/group13/blog/domain/models/blog"
	irepo "github.com/group13/blog/usecase/common/i_repo"
)

const (
	defaultLimit = 10  // Default limit for the number of blogs to retrieve
	minLimit     = 5   // Minimum limit for the number of blogs
	maxLimit     = 100 // Maximum limit for the number of blogs
)

// GetMultipleHandler handles queries for retrieving multiple blogs.
type GetMultipleHandler struct {
	blogRepo irepo.Blog // Repository for accessing blog data
}

// NewGetMultipleHandler creates a new instance of GetMultipleHandler with the given repository.
func NewGetMultipleHandler(blogRepo irepo.Blog) *GetMultipleHandler {
	return &GetMultipleHandler{blogRepo: blogRepo}
}

// Handle processes a GetMultipleQuery to retrieve multiple blogs based on the provided query parameters.
//
// Returns:
// - []*blogmodel.Blog: A slice of pointers to Blog models that match the query.
// - error: An error if the retrieval fails, such as issues with accessing the repository.
func (h *GetMultipleHandler) Handle(query *GetMultipleQuery) ([]*blogmodel.Blog, error) {
	// Set the limit with boundaries
	limit := h.determineLimit(query.Limit)

	// Determine if we are retrieving blogs for a specific user or all blogs
	if query.UserID == uuid.Nil {
		// Retrieve all blogs sorted by interaction or date
		return h.blogRepo.ListByTotalInteraction(query.LastSeenID, limit)
	}

	// Retrieve blogs for a specific author (user)
	return h.blogRepo.ListByAuthor(query.UserID, query.LastSeenID, limit)
}

// determineLimit returns a valid limit value based on the query.
func (h *GetMultipleHandler) determineLimit(requestedLimit int) int {
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
