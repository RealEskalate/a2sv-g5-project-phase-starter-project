package blogqry

import (
	"github.com/google/uuid"
)

// GetMultipleQuery represents a query for retrieving multiple blogs.
type GetMultipleQuery struct {
	userID     uuid.UUID
	limit      int
	lastSeenID *uuid.UUID
}

// NewGetMultipleQuery creates a new instance of GetMultipleQuery with the specified parameters.
func NewGetMultipleQuery(userID uuid.UUID, limit int, lastSeenID *uuid.UUID) *GetMultipleQuery {
	return &GetMultipleQuery{
		userID:     userID,
		limit:      limit,
		lastSeenID: lastSeenID,
	}
}
