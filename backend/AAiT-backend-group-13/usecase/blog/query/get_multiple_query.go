package blogqry

import (
	"github.com/google/uuid"
)

// GetMultipleQuery represents a query for retrieving multiple blogs.
type GetMultipleQuery struct {
	UserID     uuid.UUID  // ID of the author whose blogs are to be retrieved; if nil, retrieve all blogs
	Limit      int        // Maximum number of blogs to retrieve
	LastSeenID *uuid.UUID // ID of the last seen blog (for pagination)
}
